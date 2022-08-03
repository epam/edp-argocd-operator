// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1ApplicationSourceKustomize ApplicationSourceKustomize holds options specific to an Application source specific to Kustomize
//
// swagger:model v1alpha1ApplicationSourceKustomize
type V1alpha1ApplicationSourceKustomize struct {

	// CommonAnnotations is a list of additional annotations to add to rendered manifests
	CommonAnnotations map[string]string `json:"commonAnnotations,omitempty"`

	// CommonLabels is a list of additional labels to add to rendered manifests
	CommonLabels map[string]string `json:"commonLabels,omitempty"`

	// ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps
	ForceCommonAnnotations bool `json:"forceCommonAnnotations,omitempty"`

	// ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps
	ForceCommonLabels bool `json:"forceCommonLabels,omitempty"`

	// Images is a list of Kustomize image override specifications
	Images []string `json:"images"`

	// NamePrefix is a prefix appended to resources for Kustomize apps
	NamePrefix string `json:"namePrefix,omitempty"`

	// NameSuffix is a suffix appended to resources for Kustomize apps
	NameSuffix string `json:"nameSuffix,omitempty"`

	// Version controls which version of Kustomize to use for rendering manifests
	Version string `json:"version,omitempty"`
}

// Validate validates this v1alpha1 application source kustomize
func (m *V1alpha1ApplicationSourceKustomize) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1alpha1 application source kustomize based on context it is used
func (m *V1alpha1ApplicationSourceKustomize) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1ApplicationSourceKustomize) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1ApplicationSourceKustomize) UnmarshalBinary(b []byte) error {
	var res V1alpha1ApplicationSourceKustomize
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}