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

// TokenInfo token info
//
// swagger:model TokenInfo
type TokenInfo struct {

	// expires at
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expiresAt,omitempty"`

	// issued at
	// Format: date-time
	IssuedAt strfmt.DateTime `json:"issuedAt,omitempty"`

	// username
	Login string `json:"login,omitempty"`

	// roles
	Roles []string `json:"roles"`
}

// Validate validates this token info
func (m *TokenInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenInfo) validateExpiresAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expiresAt", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TokenInfo) validateIssuedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.IssuedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("issuedAt", "body", "date-time", m.IssuedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this token info based on context it is used
func (m *TokenInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TokenInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenInfo) UnmarshalBinary(b []byte) error {
	var res TokenInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
