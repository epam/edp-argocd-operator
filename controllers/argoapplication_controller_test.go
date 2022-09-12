package controllers

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	openApiRuntime "github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	argoApi "github.com/epam/edp-argocd-operator/api/v1alpha1"
	"github.com/epam/edp-argocd-operator/pkg/argoclient"
	"github.com/epam/edp-argocd-operator/pkg/argoclient/client/application_service"
)

func TestArgoApplicationReconciler_Reconcile(t *testing.T) {
	scheme := runtime.NewScheme()
	utilruntime.Must(argoApi.AddToScheme(scheme))
	utilruntime.Must(v1.AddToScheme(scheme))

	timeNow := metav1.NewTime(time.Now())

	type fields struct {
		TransportMock *transportMock
		Object        runtime.Object
	}

	type args struct {
		req controllerruntime.Request
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    controllerruntime.Result
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "create application success",
			fields: fields{
				Object: &argoApi.ArgoApplication{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "test-ns",
						Name:      "test-app",
					},
				},
				TransportMock: &transportMock{
					SubmitMock: func(operation *openApiRuntime.ClientOperation) (interface{}, error) {
						if operation.ID == "ApplicationService_Get" {
							return nil, application_service.NewApplicationServiceGetDefault(http.StatusNotFound)
						}

						if operation.ID == "ApplicationService_Create" {
							return application_service.NewApplicationServiceCreateOK(), nil
						}

						panic(fmt.Sprintf("need to mock method %s", operation.ID))
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-app",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.NoError,
		},
		{
			name: "update application success",
			fields: fields{
				Object: &argoApi.ArgoApplication{
					ObjectMeta: metav1.ObjectMeta{
						Namespace:  "test-ns",
						Name:       "test-app",
						Finalizers: []string{argoFinalizer},
					},
				},
				TransportMock: &transportMock{
					SubmitMock: func(operation *openApiRuntime.ClientOperation) (interface{}, error) {
						if operation.ID == "ApplicationService_Get" {
							return application_service.NewApplicationServiceGetOK(), nil
						}

						if operation.ID == "ApplicationService_UpdateSpec" {
							return application_service.NewApplicationServiceUpdateSpecOK(), nil
						}

						panic(fmt.Sprintf("need to mock method %s", operation.ID))
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-app",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.NoError,
		},
		{
			name: "delete application success",
			fields: fields{
				Object: &argoApi.ArgoApplication{
					ObjectMeta: metav1.ObjectMeta{
						Namespace:         "test-ns",
						Name:              "test-app",
						DeletionTimestamp: &timeNow,
						Finalizers:        []string{argoFinalizer},
					},
				},
				TransportMock: &transportMock{
					SubmitMock: func(operation *openApiRuntime.ClientOperation) (interface{}, error) {
						if operation.ID == "ApplicationService_Delete" {
							return application_service.NewApplicationServiceDeleteOK(), nil
						}

						panic(fmt.Sprintf("need to mock method %s", operation.ID))
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-app",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.NoError,
		},
		{
			name: "object with delete timestamp",
			fields: fields{
				Object: &argoApi.ArgoApplication{
					ObjectMeta: metav1.ObjectMeta{
						Namespace:         "test-ns",
						Name:              "test-app",
						DeletionTimestamp: &timeNow,
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-app",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.NoError,
		},
		{
			name: "object test-app does not exist",
			fields: fields{
				Object: &argoApi.ArgoApplication{
					ObjectMeta: metav1.ObjectMeta{
						Namespace:         "test-ns",
						Name:              "test-app2",
						DeletionTimestamp: &timeNow,
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-app",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.NoError,
		},
		{
			name: "failed to get application",
			fields: fields{
				Object: &argoApi.ArgoApplication{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "test-ns",
						Name:      "test-app",
					},
				},
				TransportMock: &transportMock{
					SubmitMock: func(operation *openApiRuntime.ClientOperation) (interface{}, error) {
						if operation.ID == "ApplicationService_Get" {
							return nil, application_service.NewApplicationServiceGetDefault(http.StatusBadRequest)
						}

						panic(fmt.Sprintf("need to mock method %s", operation.ID))
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-app",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.Error,
		},
		{
			name: "failed to update application",
			fields: fields{
				Object: &argoApi.ArgoApplication{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "test-ns",
						Name:      "test-app",
					},
				},
				TransportMock: &transportMock{
					SubmitMock: func(operation *openApiRuntime.ClientOperation) (interface{}, error) {
						if operation.ID == "ApplicationService_Get" {
							return application_service.NewApplicationServiceGetOK(), nil
						}

						if operation.ID == "ApplicationService_UpdateSpec" {
							return nil, application_service.NewApplicationServiceUpdateSpecDefault(http.StatusBadRequest)
						}

						panic(fmt.Sprintf("need to mock method %s", operation.ID))
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-app",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.Error,
		},
		{
			name: "failed to delete application",
			fields: fields{
				Object: &argoApi.ArgoApplication{
					ObjectMeta: metav1.ObjectMeta{
						Namespace:         "test-ns",
						Name:              "test-app",
						DeletionTimestamp: &timeNow,
						Finalizers:        []string{argoFinalizer},
					},
				},
				TransportMock: &transportMock{
					SubmitMock: func(operation *openApiRuntime.ClientOperation) (interface{}, error) {
						if operation.ID == "ApplicationService_Delete" {
							return nil, application_service.NewApplicationServiceDeleteDefault(http.StatusBadRequest)
						}

						panic(fmt.Sprintf("need to mock method %s", operation.ID))
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-app",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.Error,
		},
		{
			name: "failed to create application",
			fields: fields{
				Object: &argoApi.ArgoApplication{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "test-ns",
						Name:      "test-app",
					},
				},
				TransportMock: &transportMock{
					SubmitMock: func(operation *openApiRuntime.ClientOperation) (interface{}, error) {
						if operation.ID == "ApplicationService_Get" {
							return nil, application_service.NewApplicationServiceGetDefault(http.StatusNotFound)
						}

						if operation.ID == "ApplicationService_Create" {
							return nil, application_service.NewApplicationServiceCreateDefault(http.StatusBadRequest)
						}

						panic(fmt.Sprintf("need to mock method %s", operation.ID))
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-app",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(tt.fields.Object).Build()

			r := &ArgoApplicationReconciler{
				Client: fakeClient,
				Scheme: scheme,
				ApplicationHTTPClient: &argoclient.ApplicationHTTPClient{
					Client:      application_service.New(tt.fields.TransportMock, strfmt.Default),
					AuthOptions: application_service.ClientOption(func(op *openApiRuntime.ClientOperation) {}),
				},
			}

			got, err := r.Reconcile(context.Background(), tt.args.req)
			if !tt.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
