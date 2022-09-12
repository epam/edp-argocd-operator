package argoclient

import (
	"errors"
	"net/http"

	openAPIErrors "github.com/go-openapi/errors"

	"github.com/epam/edp-argocd-operator/pkg/argoclient/client/application_service"
	"github.com/epam/edp-argocd-operator/pkg/argoclient/client/repository_service"
)

func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	notFoundOpenApiErr := openAPIErrors.New(http.StatusInternalServerError, "err")
	if errors.As(err, &notFoundOpenApiErr) {
		return int(notFoundOpenApiErr.Code()) == http.StatusNotFound
	}

	notFoundRepositoryGetErr := repository_service.NewRepositoryServiceGetDefault(http.StatusInternalServerError)
	if errors.As(err, &notFoundRepositoryGetErr) {
		return int(notFoundRepositoryGetErr.Code()) == http.StatusNotFound ||
			// In case we have access only to some project
			// RBAC example: p, proj:team-foo:developer, repositories, get, team-foo/*, allow
			// ARGO CD returns 403 status on get request if repo does not exist
			int(notFoundRepositoryGetErr.Code()) == http.StatusForbidden
	}

	notFoundRepositoryDeleteErr := repository_service.NewRepositoryServiceDeleteRepositoryDefault(http.StatusInternalServerError)
	if errors.As(err, &notFoundRepositoryDeleteErr) {
		return int(notFoundRepositoryDeleteErr.Code()) == http.StatusNotFound ||
			// argo-cd returns StatusForbidden in case we try to delete not existing repository
			int(notFoundRepositoryDeleteErr.Code()) == http.StatusForbidden
	}

	notFoundApplicationGetErr := application_service.NewApplicationServiceGetDefault(http.StatusInternalServerError)
	if errors.As(err, &notFoundApplicationGetErr) {
		return int(notFoundApplicationGetErr.Code()) == http.StatusNotFound
	}

	notFoundApplicationDeleteErr := application_service.NewApplicationServiceDeleteDefault(http.StatusInternalServerError)
	if errors.As(err, &notFoundApplicationDeleteErr) {
		return int(notFoundApplicationDeleteErr.Code()) == http.StatusNotFound
	}

	return false
}
