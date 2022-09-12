package argoclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetApplicationClient(t *testing.T) {
	tests := []struct {
		name       string
		setEnvFunc func(t *testing.T)
		wantNotNil assert.BoolAssertionFunc
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			setEnvFunc: func(t *testing.T) {
				t.Setenv("ARGOCD_URL", "https://argo.com")
				t.Setenv("ARGOCD_TOKEN", "token")
			},
			wantNotNil: assert.True,
			wantErr:    assert.NoError,
		},
		{
			name: "no env",
			setEnvFunc: func(t *testing.T) {
			},
			wantNotNil: assert.False,
			wantErr:    assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setEnvFunc(t)

			got, err := GetApplicationClient()
			if !tt.wantErr(t, err) {
				return
			}
			tt.wantNotNil(t, got != nil)
		})
	}
}

func TestGetRepositoryClient(t *testing.T) {
	tests := []struct {
		name       string
		setEnvFunc func(t *testing.T)
		wantNotNil assert.BoolAssertionFunc
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			setEnvFunc: func(t *testing.T) {
				t.Setenv("ARGOCD_URL", "https://argo.com")
				t.Setenv("ARGOCD_TOKEN", "token")
			},
			wantNotNil: assert.True,
			wantErr:    assert.NoError,
		},
		{
			name: "no env",
			setEnvFunc: func(t *testing.T) {
			},
			wantNotNil: assert.False,
			wantErr:    assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setEnvFunc(t)

			got, err := GetRepositoryClient()
			if !tt.wantErr(t, err) {
				return
			}
			tt.wantNotNil(t, got != nil)
		})
	}
}

func Test_getArgoConfig(t *testing.T) {
	tests := []struct {
		name       string
		setEnvFunc func(t *testing.T)
		want       *argoAccessData
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			setEnvFunc: func(t *testing.T) {
				t.Setenv("ARGOCD_URL", "https://argo.com")
				t.Setenv("ARGOCD_TOKEN", "token")
			},
			want: &argoAccessData{
				Host:   "argo.com",
				Schema: "https",
				Token:  "token",
			},
			wantErr: assert.NoError,
		},
		{
			name: "no ARGOCD_URL error",
			setEnvFunc: func(t *testing.T) {
				t.Setenv("ARGOCD_TOKEN", "token")
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "no ARGOCD_TOKEN error",
			setEnvFunc: func(t *testing.T) {
				t.Setenv("ARGOCD_URL", "https://argo.com")
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "invalid url error",
			setEnvFunc: func(t *testing.T) {
				t.Setenv("ARGOCD_URL", ":///argocom")
				t.Setenv("ARGOCD_TOKEN", "token")
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "invalid host error",
			setEnvFunc: func(t *testing.T) {
				t.Setenv("ARGOCD_URL", "argocom")
				t.Setenv("ARGOCD_TOKEN", "token")
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "invalid schema error",
			setEnvFunc: func(t *testing.T) {
				t.Setenv("ARGOCD_URL", "//argo.com")
				t.Setenv("ARGOCD_TOKEN", "token")
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setEnvFunc(t)

			got, err := getArgoConfig()
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getTransport(t *testing.T) {
	got := getTransport(&argoAccessData{
		Host:   "argo.com",
		Schema: "https",
		Token:  "token",
	})

	assert.NotNil(t, got)
}
