// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1ResourceActionParam TODO: describe this type
// TODO: describe members of this type
//
// swagger:model v1alpha1ResourceActionParam
type V1alpha1ResourceActionParam struct {

	// default
	Default string `json:"default,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this v1alpha1 resource action param
func (m *V1alpha1ResourceActionParam) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1alpha1 resource action param based on context it is used
func (m *V1alpha1ResourceActionParam) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1ResourceActionParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1ResourceActionParam) UnmarshalBinary(b []byte) error {
	var res V1alpha1ResourceActionParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
