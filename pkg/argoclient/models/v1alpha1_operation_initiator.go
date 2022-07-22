// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1OperationInitiator OperationInitiator contains information about the initiator of an operation
//
// swagger:model v1alpha1OperationInitiator
type V1alpha1OperationInitiator struct {

	// Automated is set to true if operation was initiated automatically by the application controller.
	Automated bool `json:"automated,omitempty"`

	// Username contains the name of a user who started operation
	// +kubebuilder:object:generate=true
	Username string `json:"username,omitempty"`
}

// Validate validates this v1alpha1 operation initiator
func (m *V1alpha1OperationInitiator) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1alpha1 operation initiator based on context it is used
func (m *V1alpha1OperationInitiator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1OperationInitiator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1OperationInitiator) UnmarshalBinary(b []byte) error {
	var res V1alpha1OperationInitiator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
