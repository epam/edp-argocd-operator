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

	argoApi "github.com/epam/edp-argocd-operator/api/v1alpha1"
	"github.com/epam/edp-argocd-operator/pkg/argoclient"
	"github.com/epam/edp-argocd-operator/pkg/argoclient/client/application_service"
	"github.com/epam/edp-argocd-operator/pkg/mapper"
	"github.com/epam/edp-argocd-operator/pkg/objectutil"
)

// ArgoApplicationReconciler reconciles a ArgoApplication object.
type ArgoApplicationReconciler struct {
	client.Client
	Scheme                *runtime.Scheme
	ApplicationHTTPClient *argoclient.ApplicationHTTPClient
}

//+kubebuilder:rbac:groups=v1.edp.epam.com,namespace=placeholder,resources=argoapplications,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=v1.edp.epam.com,namespace=placeholder,resources=argoapplications/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=v1.edp.epam.com,namespace=placeholder,resources=argoapplications/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *ArgoApplicationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	reqLogger := log.FromContext(ctx)
	reqLogger.Info("Processing ArgoApplication")

	argoApplication := &argoApi.ArgoApplication{}
	if err := r.Get(ctx, req.NamespacedName, argoApplication); err != nil {
		if k8serrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		return reconcile.Result{}, fmt.Errorf("failed to get application: %w", err)
	}

	reconcileErr := r.syncApplication(ctx, reqLogger, argoApplication)

	if !objectutil.ObjectMarkDeleted(argoApplication) {
		argoApplication.Status.Status = argoApi.ApplicationReconciliationStatusSuccess
		argoApplication.Status.Error = ""

		if reconcileErr != nil {
			argoApplication.Status.Status = argoApi.ApplicationReconciliationStatusError
			argoApplication.Status.Error = reconcileErr.Error()
		}

		if err := r.Status().Update(ctx, argoApplication); err != nil {
			return reconcile.Result{}, fmt.Errorf("failed to update application status: %w", err)
		}
	}

	if reconcileErr != nil {
		return ctrl.Result{}, fmt.Errorf("application reconciliation failed: %w", reconcileErr)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ArgoApplicationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	err := ctrl.NewControllerManagedBy(mgr).
		For(&argoApi.ArgoApplication{}).
		Complete(r)

	if err != nil {
		return fmt.Errorf("failed to setup application controller: %w", err)
	}

	return nil
}

func (r *ArgoApplicationReconciler) syncApplication(
	ctx context.Context,
	reqLogger logr.Logger,
	argoApplication *argoApi.ArgoApplication,
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
				return fmt.Errorf("failed to remove finalizer from application: %w", err)
			}
		}

		return nil
	}

	if err := r.addFinalizer(ctx, argoApplication); err != nil {
		return err
	}

	getParams := application_service.NewApplicationServiceGetParams().
		WithContext(ctx).
		WithName(argoApplication.Name)

	_, err := r.ApplicationHTTPClient.Client.ApplicationServiceGet(getParams, r.ApplicationHTTPClient.AuthOptions)
	if err != nil {
		if argoclient.IsNotFoundError(err) {
			if err = r.createArgoApplication(ctx, argoApplication, r.ApplicationHTTPClient); err != nil {
				return fmt.Errorf("failed to create argoapplication: %w", err)
			}

			reqLogger.Info(fmt.Sprintf("ArgoApplication %s created", argoApplication.Name))

			return nil
		}

		return fmt.Errorf("failed to get argoapplication: %w", err)
	}

	if err := r.updateArgoApplication(ctx, argoApplication, r.ApplicationHTTPClient); err != nil {
		return fmt.Errorf("failed to update argoapplication: %w", err)
	}

	reqLogger.Info(fmt.Sprintf("ArgoApplication %s updated", argoApplication.Name))

	return nil
}

func (*ArgoApplicationReconciler) createArgoApplication(
	ctx context.Context,
	argoApplication *argoApi.ArgoApplication,
	httpClient *argoclient.ApplicationHTTPClient,
) error {
	body, err := mapper.ApplicationToRestModel(argoApplication)
	if err != nil {
		return fmt.Errorf("failed to map application to rest model: %w", err)
	}

	params := application_service.NewApplicationServiceCreateParams().
		WithContext(ctx).
		WithBody(body)

	_, err = httpClient.Client.ApplicationServiceCreate(params, httpClient.AuthOptions)
	if err != nil {
		return fmt.Errorf("failed to create application: %w", err)
	}

	return nil
}

func (*ArgoApplicationReconciler) updateArgoApplication(
	ctx context.Context,
	argoApplication *argoApi.ArgoApplication,
	httpClient *argoclient.ApplicationHTTPClient,
) error {
	body, err := mapper.ApplicationSpecToRestModel(&argoApplication.Spec)
	if err != nil {
		return fmt.Errorf("failed to map application spec to rest model: %w", err)
	}

	params := application_service.NewApplicationServiceUpdateSpecParams().
		WithContext(ctx).
		WithBody(body).
		WithName(argoApplication.Name)

	_, err = httpClient.Client.ApplicationServiceUpdateSpec(params, httpClient.AuthOptions)
	if err != nil {
		return fmt.Errorf("failed to update applcation spec: %w", err)
	}

	return nil
}

func (*ArgoApplicationReconciler) deleteArgoApplication(
	ctx context.Context,
	argoApplication *argoApi.ArgoApplication,
	httpClient *argoclient.ApplicationHTTPClient,
) error {
	params := application_service.NewApplicationServiceDeleteParams().
		WithContext(ctx).
		WithName(argoApplication.Name)

	_, err := httpClient.Client.ApplicationServiceDelete(params, httpClient.AuthOptions)
	if err != nil {
		return fmt.Errorf("failed to delete application: %w", err)
	}

	return nil
}

func (r *ArgoApplicationReconciler) finalizeArgoApplication(
	ctx context.Context,
	argoApplication *argoApi.ArgoApplication,
	applicationHTTPClient *argoclient.ApplicationHTTPClient,
) error {
	err := r.deleteArgoApplication(ctx, argoApplication, applicationHTTPClient)
	if err != nil && !argoclient.IsNotFoundError(err) {
		return fmt.Errorf("failed to delete argoapplication: %w", err)
	}

	return nil
}

func (r *ArgoApplicationReconciler) addFinalizer(
	ctx context.Context,
	argoApplication *argoApi.ArgoApplication,
) error {
	if controllerutil.ContainsFinalizer(argoApplication, argoFinalizer) {
		return nil
	}

	controllerutil.AddFinalizer(argoApplication, argoFinalizer)

	err := r.Update(ctx, argoApplication)
	if err != nil {
		return fmt.Errorf("failed to add finalaizer to argo application: %w", err)
	}

	return nil
}
