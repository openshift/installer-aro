// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Service service
//
// swagger:model Service
type Service struct {

	// bindable
	// Required: true
	Bindable *bool `json:"bindable"`

	// dashboard client
	DashboardClient *DashboardClient `json:"dashboard_client,omitempty"`

	// description
	// Required: true
	Description *string `json:"description"`

	// iam compatible
	IamCompatible bool `json:"iam_compatible,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// metadata
	Metadata Metadata `json:"metadata,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// plan updateable
	PlanUpdateable bool `json:"plan_updateable,omitempty"`

	// plans
	// Required: true
	Plans []*Plan `json:"plans"`

	// provisionable
	Provisionable bool `json:"provisionable,omitempty"`

	// rc compatible
	RcCompatible bool `json:"rc_compatible,omitempty"`

	// requires
	Requires []string `json:"requires"`

	// tags
	Tags []string `json:"tags"`

	// unique api key
	UniqueAPIKey bool `json:"unique_api_key,omitempty"`
}

// Validate validates this service
func (m *Service) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBindable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDashboardClient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlans(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequires(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Service) validateBindable(formats strfmt.Registry) error {

	if err := validate.Required("bindable", "body", m.Bindable); err != nil {
		return err
	}

	return nil
}

func (m *Service) validateDashboardClient(formats strfmt.Registry) error {
	if swag.IsZero(m.DashboardClient) { // not required
		return nil
	}

	if m.DashboardClient != nil {
		if err := m.DashboardClient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dashboard_client")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dashboard_client")
			}
			return err
		}
	}

	return nil
}

func (m *Service) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Service) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Service) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Service) validatePlans(formats strfmt.Registry) error {

	if err := validate.Required("plans", "body", m.Plans); err != nil {
		return err
	}

	for i := 0; i < len(m.Plans); i++ {
		if swag.IsZero(m.Plans[i]) { // not required
			continue
		}

		if m.Plans[i] != nil {
			if err := m.Plans[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("plans" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("plans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var serviceRequiresItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["syslog_drain","route_forwarding","volume_mount"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceRequiresItemsEnum = append(serviceRequiresItemsEnum, v)
	}
}

func (m *Service) validateRequiresItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serviceRequiresItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Service) validateRequires(formats strfmt.Registry) error {
	if swag.IsZero(m.Requires) { // not required
		return nil
	}

	for i := 0; i < len(m.Requires); i++ {

		// value enum
		if err := m.validateRequiresItemsEnum("requires"+"."+strconv.Itoa(i), "body", m.Requires[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validate this service based on the context it is used
func (m *Service) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDashboardClient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlans(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Service) contextValidateDashboardClient(ctx context.Context, formats strfmt.Registry) error {

	if m.DashboardClient != nil {

		if swag.IsZero(m.DashboardClient) { // not required
			return nil
		}

		if err := m.DashboardClient.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dashboard_client")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dashboard_client")
			}
			return err
		}
	}

	return nil
}

func (m *Service) contextValidatePlans(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Plans); i++ {

		if m.Plans[i] != nil {

			if swag.IsZero(m.Plans[i]) { // not required
				return nil
			}

			if err := m.Plans[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("plans" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("plans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Service) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Service) UnmarshalBinary(b []byte) error {
	var res Service
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
