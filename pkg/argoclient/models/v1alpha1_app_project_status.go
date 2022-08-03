// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1alpha1AppProjectStatus AppProjectStatus contains status information for AppProject CRs
// +kubebuilder:object:generate=true
//
// swagger:model v1alpha1AppProjectStatus
type V1alpha1AppProjectStatus struct {

	// JWTTokensByRole contains a list of JWT tokens issued for a given role
	JwtTokensByRole map[string]V1alpha1JWTTokens `json:"jwtTokensByRole,omitempty"`
}

// Validate validates this v1alpha1 app project status
func (m *V1alpha1AppProjectStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJwtTokensByRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1AppProjectStatus) validateJwtTokensByRole(formats strfmt.Registry) error {
	if swag.IsZero(m.JwtTokensByRole) { // not required
		return nil
	}

	for k := range m.JwtTokensByRole {

		if err := validate.Required("jwtTokensByRole"+"."+k, "body", m.JwtTokensByRole[k]); err != nil {
			return err
		}
		if val, ok := m.JwtTokensByRole[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jwtTokensByRole" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("jwtTokensByRole" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1alpha1 app project status based on the context it is used
func (m *V1alpha1AppProjectStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJwtTokensByRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1AppProjectStatus) contextValidateJwtTokensByRole(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.JwtTokensByRole {

		if val, ok := m.JwtTokensByRole[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1AppProjectStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1AppProjectStatus) UnmarshalBinary(b []byte) error {
	var res V1alpha1AppProjectStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}