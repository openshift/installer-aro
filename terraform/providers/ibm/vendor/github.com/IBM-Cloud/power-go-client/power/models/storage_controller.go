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

// StorageController Description of a Storage Controller
//
// swagger:model StorageController
type StorageController struct {

	// Display Name of the Storage Controller
	// Required: true
	DisplayName *string `json:"displayName"`

	// Health status of this storage controller
	// Required: true
	Health *string `json:"health"`

	// Free storage in user pools on this storage controller (GB)
	// Required: true
	PoolStorage *float64 `json:"poolStorage"`

	// Total storage capacity of user pools in this storage controller (GB)
	// Required: true
	PoolTotalStorage *float64 `json:"poolTotalStorage"`

	// List of storage pools within this storage controller
	// Required: true
	Pools map[string]StoragePoolCombined `json:"pools"`
}

// Validate validates this storage controller
func (m *StorageController) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoolStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoolTotalStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePools(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageController) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *StorageController) validateHealth(formats strfmt.Registry) error {

	if err := validate.Required("health", "body", m.Health); err != nil {
		return err
	}

	return nil
}

func (m *StorageController) validatePoolStorage(formats strfmt.Registry) error {

	if err := validate.Required("poolStorage", "body", m.PoolStorage); err != nil {
		return err
	}

	return nil
}

func (m *StorageController) validatePoolTotalStorage(formats strfmt.Registry) error {

	if err := validate.Required("poolTotalStorage", "body", m.PoolTotalStorage); err != nil {
		return err
	}

	return nil
}

func (m *StorageController) validatePools(formats strfmt.Registry) error {

	if err := validate.Required("pools", "body", m.Pools); err != nil {
		return err
	}

	for k := range m.Pools {

		if err := validate.Required("pools"+"."+k, "body", m.Pools[k]); err != nil {
			return err
		}
		if val, ok := m.Pools[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pools" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pools" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this storage controller based on the context it is used
func (m *StorageController) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePools(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageController) contextValidatePools(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("pools", "body", m.Pools); err != nil {
		return err
	}

	for k := range m.Pools {

		if val, ok := m.Pools[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageController) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageController) UnmarshalBinary(b []byte) error {
	var res StorageController
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
