// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SSHKeys SSH keys
//
// swagger:model SSHKeys
type SSHKeys struct {

	// SSH Keys
	// Required: true
	SSHKeys []*SSHKey `json:"sshKeys"`
}

// Validate validates this SSH keys
func (m *SSHKeys) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSSHKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SSHKeys) validateSSHKeys(formats strfmt.Registry) error {

	if err := validate.Required("sshKeys", "body", m.SSHKeys); err != nil {
		return err
	}

	for i := 0; i < len(m.SSHKeys); i++ {
		if swag.IsZero(m.SSHKeys[i]) { // not required
			continue
		}

		if m.SSHKeys[i] != nil {
			if err := m.SSHKeys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sshKeys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sshKeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this SSH keys based on the context it is used
func (m *SSHKeys) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSSHKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SSHKeys) contextValidateSSHKeys(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SSHKeys); i++ {

		if m.SSHKeys[i] != nil {

			if swag.IsZero(m.SSHKeys[i]) { // not required
				return nil
			}

			if err := m.SSHKeys[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sshKeys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sshKeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SSHKeys) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSHKeys) UnmarshalBinary(b []byte) error {
	var res SSHKeys
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
