// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1ResourceDiff ResourceDiff holds the diff of a live and target resource object
// TODO: describe members of this type
//
// swagger:model v1alpha1ResourceDiff
type V1alpha1ResourceDiff struct {

	// Diff contains the JSON patch between target and live resource
	// Deprecated: use NormalizedLiveState and PredictedLiveState to render the difference
	Diff string `json:"diff,omitempty"`

	// group
	Group string `json:"group,omitempty"`

	// hook
	Hook bool `json:"hook,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`

	// TargetState contains the JSON live resource manifest
	LiveState string `json:"liveState,omitempty"`

	// modified
	Modified bool `json:"modified,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// NormalizedLiveState contains JSON serialized live resource state with applied normalizations
	NormalizedLiveState string `json:"normalizedLiveState,omitempty"`

	// PredictedLiveState contains JSON serialized resource state that is calculated based on normalized and target resource state
	PredictedLiveState string `json:"predictedLiveState,omitempty"`

	// resource version
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// TargetState contains the JSON serialized resource manifest defined in the Git/Helm
	// +kubebuilder:object:generate=true
	TargetState string `json:"targetState,omitempty"`
}

// Validate validates this v1alpha1 resource diff
func (m *V1alpha1ResourceDiff) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1alpha1 resource diff based on context it is used
func (m *V1alpha1ResourceDiff) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1ResourceDiff) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1ResourceDiff) UnmarshalBinary(b []byte) error {
	var res V1alpha1ResourceDiff
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
