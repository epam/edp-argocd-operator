// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GpgkeyGnuPGPublicKeyCreateResponse Response to a public key creation request
//
// swagger:model gpgkeyGnuPGPublicKeyCreateResponse
type GpgkeyGnuPGPublicKeyCreateResponse struct {

	// created
	Created *V1alpha1GnuPGPublicKeyList `json:"created,omitempty"`

	// List of key IDs that haven been skipped because they already exist on the server
	Skipped []string `json:"skipped"`
}

// Validate validates this gpgkey gnu p g public key create response
func (m *GpgkeyGnuPGPublicKeyCreateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GpgkeyGnuPGPublicKeyCreateResponse) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if m.Created != nil {
		if err := m.Created.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this gpgkey gnu p g public key create response based on the context it is used
func (m *GpgkeyGnuPGPublicKeyCreateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GpgkeyGnuPGPublicKeyCreateResponse) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if m.Created != nil {
		if err := m.Created.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GpgkeyGnuPGPublicKeyCreateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GpgkeyGnuPGPublicKeyCreateResponse) UnmarshalBinary(b []byte) error {
	var res GpgkeyGnuPGPublicKeyCreateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}