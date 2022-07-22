package controllers

import (
	"context"
	"fmt"
	"net/url"

	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/epam/edp-argocd-operator/pkg/argoclient"
	"github.com/epam/edp-argocd-operator/pkg/argoclient/client/repository_service"
	"github.com/epam/edp-argocd-operator/pkg/argoclient/models"
	"github.com/epam/edp-argocd-operator/pkg/mapper"
	"github.com/epam/edp-argocd-operator/pkg/objectutil"
)

// RepositoryReconciler reconciles a Secret object and creates repositories
type RepositoryReconciler struct {
	client.Client
	Scheme               *runtime.Scheme
	RepositoryHTTPClient *argoclient.RepositoryHTTPClient
}

//+kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch;update
//+kubebuilder:rbac:groups="",resources=secrets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *RepositoryReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	reqLogger := log.FromContext(ctx)
	reqLogger.Info("Processing Repository")

	secret := &corev1.Secret{}
	if err := r.Get(ctx, req.NamespacedName, secret); err != nil {
		if k8serrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		return reconcile.Result{}, err
	}

	if secret.GetDeletionTimestamp() != nil {
		if controllerutil.ContainsFinalizer(secret, argoFinalizer) {
			if err := r.finalizeSecret(ctx, secret, r.RepositoryHTTPClient); err != nil {
				return ctrl.Result{}, err
			}

			reqLogger.Info(fmt.Sprintf("Repository %s deleted", secret.Name))

			controllerutil.RemoveFinalizer(secret, argoFinalizer)
			err := r.Update(ctx, secret)
			if err != nil {
				return ctrl.Result{}, err
			}
		}

		return ctrl.Result{}, nil
	}

	if !controllerutil.ContainsFinalizer(secret, argoFinalizer) {
		controllerutil.AddFinalizer(secret, argoFinalizer)
		err := r.Update(ctx, secret)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	repository, err := mapper.SecretToRestRepositoryModel(secret)
	if err != nil {
		return ctrl.Result{}, err
	}

	getParams := repository_service.NewRepositoryServiceGetParams().
		WithContext(ctx).
		WithRepo(url.QueryEscape(repository.Repo))

	_, err = r.RepositoryHTTPClient.Client.RepositoryServiceGet(getParams, r.RepositoryHTTPClient.AuthOptions)
	if err != nil {
		if argoclient.IsNotFoundError(err) {
			if err := r.createRepository(ctx, repository, r.RepositoryHTTPClient); err != nil {
				return reconcile.Result{}, fmt.Errorf("failed to create repository, errror: %w", err)
			}

			reqLogger.Info(fmt.Sprintf("Repository %s created", secret.Name))

			return ctrl.Result{}, nil
		}

		return reconcile.Result{}, fmt.Errorf("failed to get repository, errror: %w", err)
	}

	if err := r.updateRepository(ctx, repository, r.RepositoryHTTPClient); err != nil {
		return reconcile.Result{}, fmt.Errorf("failed to update repository, errror: %w", err)
	}

	reqLogger.Info(fmt.Sprintf("Repository %s updated", secret.Name))

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RepositoryReconciler) SetupWithManager(mgr ctrl.Manager) error {
	const argoEDPSecretLabel = "argocd.edp.epam.com/secret-type"
	const argoEDPSecretLabelRepository = "repository"

	p := predicate.Funcs{
		CreateFunc: func(e event.CreateEvent) bool {
			return objectutil.ContainsLabel(e.Object, argoEDPSecretLabel, argoEDPSecretLabelRepository)
		},
		DeleteFunc: func(e event.DeleteEvent) bool {
			return objectutil.ContainsLabel(e.Object, argoEDPSecretLabel, argoEDPSecretLabelRepository)
		},
		UpdateFunc: func(e event.UpdateEvent) bool {
			return objectutil.ContainsLabel(e.ObjectNew, argoEDPSecretLabel, argoEDPSecretLabelRepository)
		},
		GenericFunc: func(e event.GenericEvent) bool {
			return objectutil.ContainsLabel(e.Object, argoEDPSecretLabel, argoEDPSecretLabelRepository)
		},
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Secret{}).
		WithEventFilter(p).
		Complete(r)
}

func (r *RepositoryReconciler) createRepository(
	ctx context.Context,
	repository *models.V1alpha1Repository,
	httpClient *argoclient.RepositoryHTTPClient,
) error {
	params := repository_service.NewRepositoryServiceCreateRepositoryParams().
		WithContext(ctx).
		WithBody(repository)

	_, err := httpClient.Client.RepositoryServiceCreateRepository(params, httpClient.AuthOptions)

	return err
}

func (r *RepositoryReconciler) updateRepository(
	ctx context.Context,
	repository *models.V1alpha1Repository,
	httpClient *argoclient.RepositoryHTTPClient,
) error {
	params := repository_service.NewRepositoryServiceUpdateRepositoryParams().
		WithContext(ctx).
		WithBody(repository).
		WithRepoRepo(url.QueryEscape(repository.Repo))

	_, err := httpClient.Client.RepositoryServiceUpdateRepository(params, httpClient.AuthOptions)

	return err
}

func (r *RepositoryReconciler) deleteRepository(
	ctx context.Context,
	repository *models.V1alpha1Repository,
	httpClient *argoclient.RepositoryHTTPClient,
) error {
	params := repository_service.NewRepositoryServiceDeleteRepositoryParams().
		WithContext(ctx).
		WithRepo(url.QueryEscape(repository.Repo))

	_, err := httpClient.Client.RepositoryServiceDeleteRepository(params, httpClient.AuthOptions)

	return err
}

func (r *RepositoryReconciler) finalizeSecret(
	ctx context.Context,
	secret *corev1.Secret,
	httpClient *argoclient.RepositoryHTTPClient,
) error {
	repository, err := mapper.SecretToRestRepositoryModel(secret)
	if err != nil {
		return err
	}

	err = r.deleteRepository(ctx, repository, httpClient)
	if err != nil && !argoclient.IsNotFoundError(err) {
		return fmt.Errorf("failed to delete repository, errror: %w", err)
	}

	return nil
}
