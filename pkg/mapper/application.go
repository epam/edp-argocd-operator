package mapper

import (
	"fmt"

	"github.com/mitchellh/mapstructure"

	"github.com/epam/edp-argocd-operator/api/v1alpha1"
	"github.com/epam/edp-argocd-operator/pkg/argoclient/models"
)

func ApplicationToRestModel(argoApplication *v1alpha1.ArgoApplication) (*models.V1alpha1Application, error) {
	restApplication := &models.V1alpha1Application{}

	err := mapstructure.Decode(argoApplication, restApplication)
	if err != nil {
		return nil, fmt.Errorf("failde to map application to rest model, error: %w", err)
	}

	restApplication.Metadata = &models.V1ObjectMeta{
		Name: argoApplication.ObjectMeta.Name,
	}

	return restApplication, nil
}

func ApplicationSpecToRestModel(applicationSpec *v1alpha1.ApplicationSpec) (*models.V1alpha1ApplicationSpec, error) {
	restApplicationSpec := &models.V1alpha1ApplicationSpec{}

	err := mapstructure.Decode(applicationSpec, restApplicationSpec)
	if err != nil {
		return nil, fmt.Errorf("failde to map application spec to rest model, error: %w", err)
	}

	return restApplicationSpec, nil
}
