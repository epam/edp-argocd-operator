package argoclient

import (
	"net/http"

	openAPIerrors "github.com/go-openapi/errors"

	"github.com/epam/edp-argocd-operator/pkg/argoclient/client/application_service"
	"github.com/epam/edp-argocd-operator/pkg/argoclient/client/repository_service"
)

func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	if apiErr, ok := err.(openAPIerrors.Error); ok {
		return int(apiErr.Code()) == http.StatusNotFound
	}

	if apiErr, ok := err.(*repository_service.RepositoryServiceGetDefault); ok {
		return int(apiErr.Code()) == http.StatusNotFound ||
			//In case we have access only to some project
			//RBAC example: p, proj:team-foo:developer, repositories, get, team-foo/*, allow
			//ARGO CD returns 403 status on get request if repo does not exist
			int(apiErr.Code()) == http.StatusForbidden
	}

	if apiErr, ok := err.(*repository_service.RepositoryServiceDeleteRepositoryDefault); ok {
		return int(apiErr.Code()) == http.StatusNotFound ||
			int(apiErr.Code()) == http.StatusForbidden
	}

	if apiErr, ok := err.(*application_service.ApplicationServiceGetDefault); ok {
		return int(apiErr.Code()) == http.StatusNotFound
	}

	if apiErr, ok := err.(*application_service.ApplicationServiceDeleteDefault); ok {
		return int(apiErr.Code()) == http.StatusNotFound
	}

	return false
}
