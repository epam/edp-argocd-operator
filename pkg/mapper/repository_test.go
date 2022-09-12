package mapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"

	"github.com/epam/edp-argocd-operator/pkg/argoclient/models"
)

func TestSecretToRestRepositoryModel(t *testing.T) {
	type args struct {
		secret *corev1.Secret
	}

	tests := []struct {
		name    string
		args    args
		want    *models.V1alpha1Repository
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				secret: &corev1.Secret{
					Data: map[string][]byte{
						"type":          []byte("git"),
						"url":           []byte("git@github.com:profile/some-repo.git"),
						"project":       []byte("default"),
						"insecure":      []byte("true"),
						"sshPrivateKey": []byte("PRIVATE_SSH_KEY"),
					},
				},
			},
			want: &models.V1alpha1Repository{
				Type:          "git",
				Repo:          "git@github.com:profile/some-repo.git",
				Project:       "default",
				Insecure:      true,
				SSHPrivateKey: "PRIVATE_SSH_KEY",
			},
			wantErr: assert.NoError,
		},
		{
			name: "project not found",
			args: args{
				secret: &corev1.Secret{
					Data: map[string][]byte{
						"type":          []byte("git"),
						"url":           []byte("git@github.com:profile/some-repo.git"),
						"insecure":      []byte("true"),
						"sshPrivateKey": []byte("PRIVATE_SSH_KEY"),
					},
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "insecure field can't be converted to bool",
			args: args{
				secret: &corev1.Secret{
					Data: map[string][]byte{
						"type":          []byte("git"),
						"url":           []byte("git@github.com:profile/some-repo.git"),
						"project":       []byte("default"),
						"insecure":      []byte("invalid"),
						"sshPrivateKey": []byte("PRIVATE_SSH_KEY"),
					},
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "url not found",
			args: args{
				secret: &corev1.Secret{
					Data: map[string][]byte{
						"type":          []byte("git"),
						"project":       []byte("default"),
						"insecure":      []byte("true"),
						"sshPrivateKey": []byte("PRIVATE_SSH_KEY"),
					},
				},
			},
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SecretToRestRepositoryModel(tt.args.secret)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
