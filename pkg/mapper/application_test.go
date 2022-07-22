package mapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/epam/edp-argocd-operator/api/v1alpha1"
	"github.com/epam/edp-argocd-operator/pkg/argoclient/models"
)

var (
	restModel = &models.V1alpha1ApplicationSpec{
		Source: &models.V1alpha1ApplicationSource{
			RepoURL:        "git@github.com:profile/some-repo.git",
			Path:           "dev",
			TargetRevision: "main",
			Helm: &models.V1alpha1ApplicationSourceHelm{
				ValueFiles: []string{"valuefile"},
				Parameters: []*models.V1alpha1HelmParameter{
					{
						Name:        "helnparam",
						Value:       "helnparam1",
						ForceString: true,
					},
				},
				ReleaseName: "release1",
				Values:      "val1",
				FileParameters: []*models.V1alpha1HelmFileParameter{
					{
						Name: "param1",
						Path: "/dev",
					},
				},
				Version:                 "v1",
				PassCredentials:         true,
				IgnoreMissingValueFiles: true,
				SkipCrds:                true,
			},
			Kustomize: &models.V1alpha1ApplicationSourceKustomize{
				NamePrefix: "edp",
				NameSuffix: "EDP",
				Images:     []string{"img1"},
				CommonLabels: map[string]string{
					"lab1": "val1",
				},
				Version: "v2",
				CommonAnnotations: map[string]string{
					"annotation1": "val2",
				},
				ForceCommonLabels:      true,
				ForceCommonAnnotations: true,
			},
			Directory: &models.V1alpha1ApplicationSourceDirectory{
				Recurse: true,
				Jsonnet: &models.V1alpha1ApplicationSourceJsonnet{
					ExtVars: []*models.V1alpha1JsonnetVar{
						{
							Name:  "jsonvar1",
							Value: "val1",
							Code:  true,
						},
					},
					Libs: []string{"lib1"},
				},
				Exclude: "test",
				Include: "go",
			},
			Plugin: &models.V1alpha1ApplicationSourcePlugin{
				Name: "plug1",
				Env: []*models.Applicationv1alpha1EnvEntry{
					{
						Name:  "ENV_1",
						Value: "data1",
					},
				},
			},
			Chart: "edp-chart",
		},
		Destination: &models.V1alpha1ApplicationDestination{
			Server:    "https://kubernetes.default.svc",
			Namespace: "edp",
			Name:      "argo",
		},
		Project: "default",
		SyncPolicy: &models.V1alpha1SyncPolicy{
			Automated: &models.V1alpha1SyncPolicyAutomated{
				Prune:      true,
				SelfHeal:   true,
				AllowEmpty: true,
			},
			SyncOptions: []string{"command"},
			Retry: &models.V1alpha1RetryStrategy{
				Limit: 1,
				Backoff: &models.V1alpha1Backoff{
					Duration:    "100",
					Factor:      0,
					MaxDuration: "100",
				},
			},
		},
		IgnoreDifferences: []*models.V1alpha1ResourceIgnoreDifferences{
			{
				Group:                 "dev",
				Kind:                  "APP",
				Name:                  "edp",
				Namespace:             "argo",
				JSONPointers:          []string{"val1"},
				ManagedFieldsManagers: []string{"field1"},
			},
		},
		Info: []*models.V1alpha1Info{
			{
				Name:  "info1",
				Value: "val1",
			},
		},
		RevisionHistoryLimit: 0,
	}

	argoModel = &v1alpha1.ApplicationSpec{
		Source: v1alpha1.ApplicationSource{
			RepoURL:        "git@github.com:profile/some-repo.git",
			Path:           "dev",
			TargetRevision: "main",
			Helm: &v1alpha1.ApplicationSourceHelm{
				ValueFiles: []string{"valuefile"},
				Parameters: []v1alpha1.HelmParameter{
					{
						Name:        "helnparam",
						Value:       "helnparam1",
						ForceString: true,
					},
				},
				ReleaseName: "release1",
				Values:      "val1",
				FileParameters: []v1alpha1.HelmFileParameter{
					{
						Name: "param1",
						Path: "/dev",
					},
				},
				Version:                 "v1",
				PassCredentials:         true,
				IgnoreMissingValueFiles: true,
				SkipCrds:                true,
			},
			Kustomize: &v1alpha1.ApplicationSourceKustomize{
				NamePrefix: "edp",
				NameSuffix: "EDP",
				Images: v1alpha1.KustomizeImages{
					v1alpha1.KustomizeImage("img1"),
				},
				CommonLabels: map[string]string{
					"lab1": "val1",
				},
				Version: "v2",
				CommonAnnotations: map[string]string{
					"annotation1": "val2",
				},
				ForceCommonLabels:      true,
				ForceCommonAnnotations: true,
			},
			Directory: &v1alpha1.ApplicationSourceDirectory{
				Recurse: true,
				Jsonnet: v1alpha1.ApplicationSourceJsonnet{
					ExtVars: []v1alpha1.JsonnetVar{
						{
							Name:  "jsonvar1",
							Value: "val1",
							Code:  true,
						},
					},
					Libs: []string{"lib1"},
				},
				Exclude: "test",
				Include: "go",
			},
			Plugin: &v1alpha1.ApplicationSourcePlugin{
				Name: "plug1",
				Env: []*v1alpha1.EnvEntry{
					{
						Name:  "ENV_1",
						Value: "data1",
					},
				},
			},
			Chart: "edp-chart",
		},
		Destination: v1alpha1.ApplicationDestination{
			Server:    "https://kubernetes.default.svc",
			Namespace: "edp",
			Name:      "argo",
		},
		Project: "default",
		SyncPolicy: &v1alpha1.SyncPolicy{
			Automated: &v1alpha1.SyncPolicyAutomated{
				Prune:      true,
				SelfHeal:   true,
				AllowEmpty: true,
			},
			SyncOptions: []string{"command"},
			Retry: &v1alpha1.RetryStrategy{
				Limit: 1,
				Backoff: &v1alpha1.Backoff{
					Duration:    "100",
					Factor:      nil,
					MaxDuration: "100",
				},
			},
		},
		IgnoreDifferences: []v1alpha1.ResourceIgnoreDifferences{
			{
				Group:                 "dev",
				Kind:                  "APP",
				Name:                  "edp",
				Namespace:             "argo",
				JSONPointers:          []string{"val1"},
				ManagedFieldsManagers: []string{"field1"},
			},
		},
		Info: []v1alpha1.Info{
			{
				Name:  "info1",
				Value: "val1",
			},
		},
		RevisionHistoryLimit: nil,
	}
)

func TestApplicationSpecToRestModel(t *testing.T) {
	t.Parallel()

	type args struct {
		applicationSpec *v1alpha1.ApplicationSpec
	}
	tests := []struct {
		name    string
		args    args
		want    *models.V1alpha1ApplicationSpec
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success full model",
			args: args{
				applicationSpec: argoModel,
			},
			want:    restModel,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ApplicationSpecToRestModel(tt.args.applicationSpec)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestApplicationToRestModel(t *testing.T) {
	t.Parallel()

	type args struct {
		argoApplication *v1alpha1.ArgoApplication
	}
	tests := []struct {
		name    string
		args    args
		want    *models.V1alpha1Application
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success full model",
			args: args{
				argoApplication: &v1alpha1.ArgoApplication{
					ObjectMeta: metav1.ObjectMeta{
						Name: "argoapp",
					},
					Spec: *argoModel,
				},
			},
			want: &models.V1alpha1Application{
				Metadata: &models.V1ObjectMeta{Name: "argoapp"},
				Spec:     restModel,
				Status:   &models.V1alpha1ApplicationStatus{},
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ApplicationToRestModel(tt.args.argoApplication)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, got, tt.want)
		})
	}
}
