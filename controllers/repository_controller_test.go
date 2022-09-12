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
	"github.com/epam/edp-argocd-operator/pkg/argoclient/client/repository_service"
)

type transportMock struct {
	SubmitMock func(operation *openApiRuntime.ClientOperation) (interface{}, error)
}

func (t *transportMock) Submit(operation *openApiRuntime.ClientOperation) (interface{}, error) {
	return t.SubmitMock(operation)
}

func TestRepositoryReconciler_Reconcile(t *testing.T) {
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
			name: "create repository success",
			fields: fields{
				Object: &v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "test-ns",
						Name:      "test-secret",
					},
					Data: map[string][]byte{
						"project":       []byte("project1"),
						"url":           []byte("https://argo.com"),
						"sshPrivateKey": []byte("sshPrivateSecretKey"),
						"type":          []byte("git"),
						"insecure":      []byte("true"),
					},
				},
				TransportMock: &transportMock{
					SubmitMock: func(operation *openApiRuntime.ClientOperation) (interface{}, error) {
						if operation.ID == "RepositoryService_Get" {
							return nil, repository_service.NewRepositoryServiceGetDefault(http.StatusNotFound)
						}

						if operation.ID == "RepositoryService_CreateRepository" {
							return repository_service.NewRepositoryServiceCreateRepositoryOK(), nil
						}

						panic(fmt.Sprintf("need to mock method %s", operation.ID))
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-secret",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.NoError,
		},
		{
			name: "update repository success",
			fields: fields{
				Object: &v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Namespace:  "test-ns",
						Name:       "test-secret",
						Finalizers: []string{argoFinalizer},
					},
					Data: map[string][]byte{
						"project":       []byte("project1"),
						"url":           []byte("https://argo.com"),
						"sshPrivateKey": []byte("sshPrivateSecretKey"),
						"type":          []byte("git"),
						"insecure":      []byte("true"),
					},
				},
				TransportMock: &transportMock{
					SubmitMock: func(operation *openApiRuntime.ClientOperation) (interface{}, error) {
						if operation.ID == "RepositoryService_Get" {
							return repository_service.NewRepositoryServiceGetOK(), nil
						}

						if operation.ID == "RepositoryService_UpdateRepository" {
							return repository_service.NewRepositoryServiceUpdateRepositoryOK(), nil
						}

						panic(fmt.Sprintf("need to mock method %s", operation.ID))
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-secret",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.NoError,
		},
		{
			name: "delete repository success",
			fields: fields{
				Object: &v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Namespace:         "test-ns",
						Name:              "test-secret",
						Finalizers:        []string{argoFinalizer},
						DeletionTimestamp: &timeNow,
					},
					Data: map[string][]byte{
						"project":       []byte("project1"),
						"url":           []byte("https://argo.com"),
						"sshPrivateKey": []byte("sshPrivateSecretKey"),
						"type":          []byte("git"),
						"insecure":      []byte("true"),
					},
				},
				TransportMock: &transportMock{
					SubmitMock: func(operation *openApiRuntime.ClientOperation) (interface{}, error) {
						if operation.ID == "RepositoryService_DeleteRepository" {
							return repository_service.NewRepositoryServiceDeleteRepositoryOK(), nil
						}

						panic(fmt.Sprintf("need to mock method %s", operation.ID))
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-secret",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.NoError,
		},
		{
			name: "object with delete timestamp",
			fields: fields{
				Object: &v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Namespace:         "test-ns",
						Name:              "test-secret",
						DeletionTimestamp: &timeNow,
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-secret",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.NoError,
		},
		{
			name: "object test-secret does not exist",
			fields: fields{
				Object: &v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "test-ns",
						Name:      "another-secret",
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-secret",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.NoError,
		},
		{
			name: "object with empty data",
			fields: fields{
				Object: &v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "test-ns",
						Name:      "test-secret",
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-secret",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.Error,
		},
		{
			name: "failed to get repository",
			fields: fields{
				Object: &v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "test-ns",
						Name:      "test-secret",
					},
					Data: map[string][]byte{
						"project":       []byte("project1"),
						"url":           []byte("https://argo.com"),
						"sshPrivateKey": []byte("sshPrivateSecretKey"),
						"type":          []byte("git"),
						"insecure":      []byte("true"),
					},
				},
				TransportMock: &transportMock{
					SubmitMock: func(operation *openApiRuntime.ClientOperation) (interface{}, error) {
						if operation.ID == "RepositoryService_Get" {
							return nil, repository_service.NewRepositoryServiceGetDefault(http.StatusBadRequest)
						}

						panic(fmt.Sprintf("need to mock method %s", operation.ID))
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-secret",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.Error,
		},
		{
			name: "failed to update repository",
			fields: fields{
				Object: &v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Namespace:  "test-ns",
						Name:       "test-secret",
						Finalizers: []string{argoFinalizer},
					},
					Data: map[string][]byte{
						"project":       []byte("project1"),
						"url":           []byte("https://argo.com"),
						"sshPrivateKey": []byte("sshPrivateSecretKey"),
						"type":          []byte("git"),
						"insecure":      []byte("true"),
					},
				},
				TransportMock: &transportMock{
					SubmitMock: func(operation *openApiRuntime.ClientOperation) (interface{}, error) {
						if operation.ID == "RepositoryService_Get" {
							return repository_service.NewRepositoryServiceGetOK(), nil
						}

						if operation.ID == "RepositoryService_UpdateRepository" {
							return nil, repository_service.NewRepositoryServiceUpdateRepositoryDefault(http.StatusBadRequest)
						}

						panic(fmt.Sprintf("need to mock method %s", operation.ID))
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-secret",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.Error,
		},
		{
			name: "failed to delete repository",
			fields: fields{
				Object: &v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Namespace:         "test-ns",
						Name:              "test-secret",
						Finalizers:        []string{argoFinalizer},
						DeletionTimestamp: &timeNow,
					},
					Data: map[string][]byte{
						"project":       []byte("project1"),
						"url":           []byte("https://argo.com"),
						"sshPrivateKey": []byte("sshPrivateSecretKey"),
						"type":          []byte("git"),
						"insecure":      []byte("true"),
					},
				},
				TransportMock: &transportMock{
					SubmitMock: func(operation *openApiRuntime.ClientOperation) (interface{}, error) {
						if operation.ID == "RepositoryService_DeleteRepository" {
							return nil, repository_service.NewRepositoryServiceDeleteRepositoryDefault(http.StatusBadRequest)
						}

						panic(fmt.Sprintf("need to mock method %s", operation.ID))
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-secret",
					},
				},
			},
			want:    controllerruntime.Result{},
			wantErr: assert.Error,
		},
		{
			name: "failed to create repository",
			fields: fields{
				Object: &v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "test-ns",
						Name:      "test-secret",
					},
					Data: map[string][]byte{
						"project":       []byte("project1"),
						"url":           []byte("https://argo.com"),
						"sshPrivateKey": []byte("sshPrivateSecretKey"),
						"type":          []byte("git"),
						"insecure":      []byte("true"),
					},
				},
				TransportMock: &transportMock{
					SubmitMock: func(operation *openApiRuntime.ClientOperation) (interface{}, error) {
						if operation.ID == "RepositoryService_Get" {
							return nil, repository_service.NewRepositoryServiceGetDefault(http.StatusNotFound)
						}

						if operation.ID == "RepositoryService_CreateRepository" {
							return nil, repository_service.NewRepositoryServiceCreateRepositoryDefault(http.StatusBadRequest)
						}

						panic(fmt.Sprintf("need to mock method %s", operation.ID))
					},
				},
			},
			args: args{
				req: controllerruntime.Request{
					NamespacedName: types.NamespacedName{
						Namespace: "test-ns",
						Name:      "test-secret",
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

			r := &RepositoryReconciler{
				Client: fakeClient,
				Scheme: scheme,
				RepositoryHTTPClient: &argoclient.RepositoryHTTPClient{
					Client:      repository_service.New(tt.fields.TransportMock, strfmt.Default),
					AuthOptions: repository_service.ClientOption(func(op *openApiRuntime.ClientOperation) {}),
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
