package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	applicationApi "github.com/epam/edp-argocd-operator/api/v1alpha1"
	"github.com/epam/edp-argocd-operator/pkg/argoclient"
	"github.com/epam/edp-argocd-operator/pkg/argoclient/client/application_service"
	"github.com/epam/edp-argocd-operator/pkg/mapper"
)

// ArgoApplicationReconciler reconciles a ArgoApplication object
type ArgoApplicationReconciler struct {
	client.Client
	Scheme                *runtime.Scheme
	ApplicationHTTPClient *argoclient.ApplicationHTTPClient
}

//+kubebuilder:rbac:groups=v1.edp.epam.com,resources=argoapplications,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=v1.edp.epam.com,resources=argoapplications/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=v1.edp.epam.com,resources=argoapplications/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *ArgoApplicationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	reqLogger := log.FromContext(ctx)
	reqLogger.Info("Processing ArgoApplication")

	argoApplication := &applicationApi.ArgoApplication{}
	if err := r.Get(ctx, req.NamespacedName, argoApplication); err != nil {
		if k8serrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		return reconcile.Result{}, err
	}

	argoApplication.Status.Status = applicationApi.ApplicationReconciliationStatusSuccess
	argoApplication.Status.Error = ""

	err := r.syncApplication(ctx, reqLogger, argoApplication)
	if err != nil {
		argoApplication.Status.Status = applicationApi.ApplicationReconciliationStatusError
		argoApplication.Status.Error = err.Error()
	}

	if err := r.Status().Update(ctx, argoApplication); err != nil {
		return reconcile.Result{}, err
	}

	return ctrl.Result{}, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *ArgoApplicationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&applicationApi.ArgoApplication{}).
		Complete(r)
}

func (r *ArgoApplicationReconciler) syncApplication(
	ctx context.Context,
	reqLogger logr.Logger,
	argoApplication *applicationApi.ArgoApplication,
) error {
	if argoApplication.GetDeletionTimestamp() != nil {
		if controllerutil.ContainsFinalizer(argoApplication, argoFinalizer) {
			if err := r.finalizeArgoApplication(ctx, argoApplication, r.ApplicationHTTPClient); err != nil {
				return err
			}

			reqLogger.Info(fmt.Sprintf("ArgoApplication %s deleted", argoApplication.Name))

			controllerutil.RemoveFinalizer(argoApplication, argoFinalizer)
			err := r.Update(ctx, argoApplication)
			if err != nil {
				return err
			}
		}

		return nil
	}

	if !controllerutil.ContainsFinalizer(argoApplication, argoFinalizer) {
		controllerutil.AddFinalizer(argoApplication, argoFinalizer)
		err := r.Update(ctx, argoApplication)
		if err != nil {
			return err
		}
	}

	getParams := application_service.NewApplicationServiceGetParams().
		WithContext(ctx).
		WithName(argoApplication.Name)

	_, err := r.ApplicationHTTPClient.Client.ApplicationServiceGet(getParams, r.ApplicationHTTPClient.AuthOptions)
	if err != nil {
		if argoclient.IsNotFoundError(err) {
			if err := r.createArgoApplication(ctx, argoApplication, r.ApplicationHTTPClient); err != nil {
				return fmt.Errorf("failed to create argoapplication, error: %w", err)
			}

			reqLogger.Info(fmt.Sprintf("ArgoApplication %s created", argoApplication.Name))

			return nil
		}

		return fmt.Errorf("failed to get argoapplication, error: %w", err)
	}

	if err := r.updateArgoApplication(ctx, argoApplication, r.ApplicationHTTPClient); err != nil {
		return fmt.Errorf("failed to update argoapplication, error: %w", err)
	}

	reqLogger.Info(fmt.Sprintf("ArgoApplication %s updated", argoApplication.Name))

	return nil
}

func (r *ArgoApplicationReconciler) createArgoApplication(
	ctx context.Context,
	argoApplication *applicationApi.ArgoApplication,
	httpClient *argoclient.ApplicationHTTPClient,
) error {
	body, err := mapper.ApplicationToRestModel(argoApplication)
	if err != nil {
		return err
	}

	params := application_service.NewApplicationServiceCreateParams().
		WithContext(ctx).
		WithBody(body)

	_, err = httpClient.Client.ApplicationServiceCreate(params, httpClient.AuthOptions)

	return err
}

func (r *ArgoApplicationReconciler) updateArgoApplication(
	ctx context.Context,
	argoApplication *applicationApi.ArgoApplication,
	httpClient *argoclient.ApplicationHTTPClient,
) error {
	body, err := mapper.ApplicationSpecToRestModel(&argoApplication.Spec)
	if err != nil {
		return err
	}

	params := application_service.NewApplicationServiceUpdateSpecParams().
		WithContext(ctx).
		WithBody(body).
		WithName(argoApplication.Name)

	_, err = httpClient.Client.ApplicationServiceUpdateSpec(params, httpClient.AuthOptions)

	return err
}

func (r *ArgoApplicationReconciler) deleteArgoApplication(
	ctx context.Context,
	argoApplication *applicationApi.ArgoApplication,
	httpClient *argoclient.ApplicationHTTPClient,
) error {
	params := application_service.NewApplicationServiceDeleteParams().
		WithContext(ctx).
		WithName(argoApplication.Name)

	_, err := httpClient.Client.ApplicationServiceDelete(params, httpClient.AuthOptions)

	return err
}

func (r *ArgoApplicationReconciler) finalizeArgoApplication(
	ctx context.Context,
	argoApplication *applicationApi.ArgoApplication,
	applicationHTTPClient *argoclient.ApplicationHTTPClient,
) error {
	err := r.deleteArgoApplication(ctx, argoApplication, applicationHTTPClient)
	if err != nil && !argoclient.IsNotFoundError(err) {
		return fmt.Errorf("failed to delete argoapplication, error: %w", err)
	}

	return nil
}
