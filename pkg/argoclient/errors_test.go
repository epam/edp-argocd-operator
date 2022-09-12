package argoclient

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	openAPIErrors "github.com/go-openapi/errors"
	"github.com/stretchr/testify/assert"

	"github.com/epam/edp-argocd-operator/pkg/argoclient/client/application_service"
	"github.com/epam/edp-argocd-operator/pkg/argoclient/client/repository_service"
)

func TestIsNotFoundError(t *testing.T) {
	type args struct {
		err error
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success openAPI error",
			args: args{
				err: openAPIErrors.New(http.StatusNotFound, "not found"),
			},
			want: true,
		},
		{
			name: "success openAPI error, wrapped",
			args: args{
				err: fmt.Errorf("%w", openAPIErrors.New(http.StatusNotFound, "not found")),
			},
			want: true,
		},
		{
			name: "openAPI error, wrong code",
			args: args{
				err: fmt.Errorf("%w", openAPIErrors.New(http.StatusInternalServerError, "err")),
			},
			want: false,
		},
		{
			name: "success repository get error",
			args: args{
				err: repository_service.NewRepositoryServiceGetDefault(http.StatusNotFound),
			},
			want: true,
		},
		{
			name: "success repository get error, wrapped",
			args: args{
				err: fmt.Errorf("%w", repository_service.NewRepositoryServiceGetDefault(http.StatusNotFound)),
			},
			want: true,
		},
		{
			name: "repository get error, wrong code",
			args: args{
				err: fmt.Errorf("%w", repository_service.NewRepositoryServiceGetDefault(http.StatusBadRequest)),
			},
			want: false,
		},
		{
			name: "success repository delete error",
			args: args{
				err: repository_service.NewRepositoryServiceDeleteRepositoryDefault(http.StatusNotFound),
			},
			want: true,
		},
		{
			name: "success repository delete error, wrapped",
			args: args{
				err: fmt.Errorf("%w", repository_service.NewRepositoryServiceDeleteRepositoryDefault(http.StatusNotFound)),
			},
			want: true,
		},
		{
			name: "repository delete error, wrong code",
			args: args{
				err: fmt.Errorf("%w", repository_service.NewRepositoryServiceDeleteRepositoryDefault(http.StatusBadRequest)),
			},
			want: false,
		},
		{
			name: "success application get error",
			args: args{
				err: application_service.NewApplicationServiceGetDefault(http.StatusNotFound),
			},
			want: true,
		},
		{
			name: "success application get error, wrapped",
			args: args{
				err: fmt.Errorf("%w", application_service.NewApplicationServiceGetDefault(http.StatusNotFound)),
			},
			want: true,
		},
		{
			name: "application get error, wrong code",
			args: args{
				err: fmt.Errorf("%w", application_service.NewApplicationServiceGetDefault(http.StatusBadRequest)),
			},
			want: false,
		},
		{
			name: "success application delete error",
			args: args{
				err: application_service.NewApplicationServiceDeleteDefault(http.StatusNotFound),
			},
			want: true,
		},
		{
			name: "success application delete error, wrapped",
			args: args{
				err: fmt.Errorf("%w", application_service.NewApplicationServiceDeleteDefault(http.StatusNotFound)),
			},
			want: true,
		},
		{
			name: "application delete error, wrong code",
			args: args{
				err: fmt.Errorf("%w", application_service.NewApplicationServiceDeleteDefault(http.StatusBadRequest)),
			},
			want: false,
		},
		{
			name: "nil error",
			args: args{
				err: nil,
			},
			want: false,
		},
		{
			name: "wrong error type",
			args: args{
				errors.New("error"),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotFoundError(tt.args.err); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
