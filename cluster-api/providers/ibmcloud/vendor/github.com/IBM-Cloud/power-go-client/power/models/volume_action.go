// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VolumeAction volume action
//
// swagger:model VolumeAction
type VolumeAction struct {

	// Indicates if the volume should be replication enabled or not
	ReplicationEnabled *bool `json:"replicationEnabled,omitempty"`

	// Target storage tier; used to change a volume's storage tier
	TargetStorageTier *string `json:"targetStorageTier,omitempty"`
}

// Validate validates this volume action
func (m *VolumeAction) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume action based on context it is used
func (m *VolumeAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeAction) UnmarshalBinary(b []byte) error {
	var res VolumeAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
