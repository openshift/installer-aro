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

// Volumes volumes
//
// swagger:model Volumes
type Volumes struct {

	// Volumes
	// Required: true
	Volumes []*VolumeReference `json:"volumes"`
}

// Validate validates this volumes
func (m *Volumes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVolumes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Volumes) validateVolumes(formats strfmt.Registry) error {

	if err := validate.Required("volumes", "body", m.Volumes); err != nil {
		return err
	}

	for i := 0; i < len(m.Volumes); i++ {
		if swag.IsZero(m.Volumes[i]) { // not required
			continue
		}

		if m.Volumes[i] != nil {
			if err := m.Volumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this volumes based on the context it is used
func (m *Volumes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Volumes) contextValidateVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Volumes); i++ {

		if m.Volumes[i] != nil {

			if swag.IsZero(m.Volumes[i]) { // not required
				return nil
			}

			if err := m.Volumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Volumes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Volumes) UnmarshalBinary(b []byte) error {
	var res Volumes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
