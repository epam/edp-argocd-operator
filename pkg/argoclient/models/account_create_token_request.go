// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountCreateTokenRequest account create token request
//
// swagger:model accountCreateTokenRequest
type AccountCreateTokenRequest struct {

	// expiresIn represents a duration in seconds
	ExpiresIn int64 `json:"expiresIn,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this account create token request
func (m *AccountCreateTokenRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account create token request based on context it is used
func (m *AccountCreateTokenRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountCreateTokenRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountCreateTokenRequest) UnmarshalBinary(b []byte) error {
	var res AccountCreateTokenRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
