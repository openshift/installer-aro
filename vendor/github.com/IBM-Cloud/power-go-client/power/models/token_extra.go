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

// TokenExtra token extra
//
// swagger:model TokenExtra
type TokenExtra struct {

	// Number of seconds token will expire
	// Required: true
	ExpiresIn *float64 `json:"expiresIn"`

	// Time on the service broker
	// Required: true
	// Format: date-time
	ServerTime *strfmt.DateTime `json:"serverTime"`

	// OAuth Token
	// Required: true
	Token *Token `json:"token"`

	// Is this token valid
	// Required: true
	Valid *bool `json:"valid"`
}

// Validate validates this token extra
func (m *TokenExtra) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiresIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenExtra) validateExpiresIn(formats strfmt.Registry) error {

	if err := validate.Required("expiresIn", "body", m.ExpiresIn); err != nil {
		return err
	}

	return nil
}

func (m *TokenExtra) validateServerTime(formats strfmt.Registry) error {

	if err := validate.Required("serverTime", "body", m.ServerTime); err != nil {
		return err
	}

	if err := validate.FormatOf("serverTime", "body", "date-time", m.ServerTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TokenExtra) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	if m.Token != nil {
		if err := m.Token.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token")
			}
			return err
		}
	}

	return nil
}

func (m *TokenExtra) validateValid(formats strfmt.Registry) error {

	if err := validate.Required("valid", "body", m.Valid); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this token extra based on the context it is used
func (m *TokenExtra) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenExtra) contextValidateToken(ctx context.Context, formats strfmt.Registry) error {

	if m.Token != nil {

		if err := m.Token.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TokenExtra) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenExtra) UnmarshalBinary(b []byte) error {
	var res TokenExtra
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
