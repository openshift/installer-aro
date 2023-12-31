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

// CloudConnectionVirtualPrivateClouds cloud connection virtual private clouds
//
// swagger:model CloudConnectionVirtualPrivateClouds
type CloudConnectionVirtualPrivateClouds struct {

	// list of available virtual private clouds
	// Required: true
	VirtualPrivateClouds []*CloudConnectionVirtualPrivateCloud `json:"virtualPrivateClouds"`
}

// Validate validates this cloud connection virtual private clouds
func (m *CloudConnectionVirtualPrivateClouds) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVirtualPrivateClouds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudConnectionVirtualPrivateClouds) validateVirtualPrivateClouds(formats strfmt.Registry) error {

	if err := validate.Required("virtualPrivateClouds", "body", m.VirtualPrivateClouds); err != nil {
		return err
	}

	for i := 0; i < len(m.VirtualPrivateClouds); i++ {
		if swag.IsZero(m.VirtualPrivateClouds[i]) { // not required
			continue
		}

		if m.VirtualPrivateClouds[i] != nil {
			if err := m.VirtualPrivateClouds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualPrivateClouds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("virtualPrivateClouds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cloud connection virtual private clouds based on the context it is used
func (m *CloudConnectionVirtualPrivateClouds) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVirtualPrivateClouds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudConnectionVirtualPrivateClouds) contextValidateVirtualPrivateClouds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VirtualPrivateClouds); i++ {

		if m.VirtualPrivateClouds[i] != nil {

			if swag.IsZero(m.VirtualPrivateClouds[i]) { // not required
				return nil
			}

			if err := m.VirtualPrivateClouds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualPrivateClouds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("virtualPrivateClouds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudConnectionVirtualPrivateClouds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudConnectionVirtualPrivateClouds) UnmarshalBinary(b []byte) error {
	var res CloudConnectionVirtualPrivateClouds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CloudConnectionVirtualPrivateCloud cloud connection virtual private cloud
//
// swagger:model CloudConnectionVirtualPrivateCloud
type CloudConnectionVirtualPrivateCloud struct {

	// indicates if vpc uses classic architecture
	// Required: true
	ClassicAccess *bool `json:"classicAccess"`

	// name for the vpc
	// Required: true
	Name *string `json:"name"`

	// status of this vpc
	// Required: true
	Status *string `json:"status"`

	// virtual private cloud id
	// Required: true
	VpcID *string `json:"vpcID"`
}

// Validate validates this cloud connection virtual private cloud
func (m *CloudConnectionVirtualPrivateCloud) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClassicAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudConnectionVirtualPrivateCloud) validateClassicAccess(formats strfmt.Registry) error {

	if err := validate.Required("classicAccess", "body", m.ClassicAccess); err != nil {
		return err
	}

	return nil
}

func (m *CloudConnectionVirtualPrivateCloud) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CloudConnectionVirtualPrivateCloud) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *CloudConnectionVirtualPrivateCloud) validateVpcID(formats strfmt.Registry) error {

	if err := validate.Required("vpcID", "body", m.VpcID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cloud connection virtual private cloud based on context it is used
func (m *CloudConnectionVirtualPrivateCloud) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloudConnectionVirtualPrivateCloud) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudConnectionVirtualPrivateCloud) UnmarshalBinary(b []byte) error {
	var res CloudConnectionVirtualPrivateCloud
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
