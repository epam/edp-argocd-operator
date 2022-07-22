package mapper

import (
	"errors"
	"fmt"
	"strconv"

	corev1 "k8s.io/api/core/v1"

	"github.com/epam/edp-argocd-operator/pkg/argoclient/models"
)

func SecretToRestRepositoryModel(secret *corev1.Secret) (*models.V1alpha1Repository, error) {
	restRepository := &models.V1alpha1Repository{}

	restRepository.Name = secret.Name

	project, ok := secret.Data["project"]
	if !ok {
		return nil, errors.New("repository project can't be empty")
	}
	restRepository.Project = string(project)

	url, ok := secret.Data["url"]
	if !ok {
		return nil, errors.New("repository url can't be empty")
	}
	restRepository.Repo = string(url)

	if sshPrivate, ok := secret.Data["sshPrivateKey"]; ok {
		restRepository.SSHPrivateKey = string(sshPrivate)
	}

	if repType, ok := secret.Data["type"]; ok {
		restRepository.Type = string(repType)
	}

	if _, ok := secret.Data["insecure"]; ok {
		insecure, err := strconv.ParseBool(string(secret.Data["insecure"]))
		if err != nil {
			return nil, fmt.Errorf("failde to get insecure value, error %w", err)
		}
		restRepository.Insecure = insecure
	}

	return restRepository, nil
}
