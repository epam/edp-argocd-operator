// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1RepoCreds RepoCreds holds the definition for repository credentials
//
// swagger:model v1alpha1RepoCreds
type V1alpha1RepoCreds struct {

	// EnableOCI specifies whether helm-oci support should be enabled for this repo
	EnableOCI bool `json:"enableOCI,omitempty"`

	// GithubAppEnterpriseBaseURL specifies the GitHub API URL for GitHub app authentication. If empty will default to https://api.github.com
	GithubAppEnterpriseBaseURL string `json:"githubAppEnterpriseBaseUrl,omitempty"`

	// GithubAppId specifies the Github App ID of the app used to access the repo for GitHub app authentication
	GithubAppID int64 `json:"githubAppID,omitempty"`

	// GithubAppInstallationId specifies the ID of the installed GitHub App for GitHub app authentication
	GithubAppInstallationID int64 `json:"githubAppInstallationID,omitempty"`

	// GithubAppPrivateKey specifies the private key PEM data for authentication via GitHub app
	GithubAppPrivateKey string `json:"githubAppPrivateKey,omitempty"`

	// Password for authenticating at the repo server
	Password string `json:"password,omitempty"`

	// SSHPrivateKey contains the private key data for authenticating at the repo server using SSH (only Git repos)
	SSHPrivateKey string `json:"sshPrivateKey,omitempty"`

	// TLSClientCertData specifies the TLS client cert data for authenticating at the repo server
	TLSClientCertData string `json:"tlsClientCertData,omitempty"`

	// TLSClientCertKey specifies the TLS client cert key for authenticating at the repo server
	TLSClientCertKey string `json:"tlsClientCertKey,omitempty"`

	// Type specifies the type of the repoCreds. Can be either "git" or "helm. "git" is assumed if empty or absent.
	Type string `json:"type,omitempty"`

	// URL is the URL that this credentials matches to
	URL string `json:"url,omitempty"`

	// Username for authenticating at the repo server
	Username string `json:"username,omitempty"`
}

// Validate validates this v1alpha1 repo creds
func (m *V1alpha1RepoCreds) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1alpha1 repo creds based on context it is used
func (m *V1alpha1RepoCreds) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1RepoCreds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1RepoCreds) UnmarshalBinary(b []byte) error {
	var res V1alpha1RepoCreds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
