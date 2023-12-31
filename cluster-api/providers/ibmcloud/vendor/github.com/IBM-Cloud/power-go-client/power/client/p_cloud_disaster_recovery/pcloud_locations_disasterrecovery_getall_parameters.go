// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_disaster_recovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPcloudLocationsDisasterrecoveryGetallParams creates a new PcloudLocationsDisasterrecoveryGetallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudLocationsDisasterrecoveryGetallParams() *PcloudLocationsDisasterrecoveryGetallParams {
	return &PcloudLocationsDisasterrecoveryGetallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudLocationsDisasterrecoveryGetallParamsWithTimeout creates a new PcloudLocationsDisasterrecoveryGetallParams object
// with the ability to set a timeout on a request.
func NewPcloudLocationsDisasterrecoveryGetallParamsWithTimeout(timeout time.Duration) *PcloudLocationsDisasterrecoveryGetallParams {
	return &PcloudLocationsDisasterrecoveryGetallParams{
		timeout: timeout,
	}
}

// NewPcloudLocationsDisasterrecoveryGetallParamsWithContext creates a new PcloudLocationsDisasterrecoveryGetallParams object
// with the ability to set a context for a request.
func NewPcloudLocationsDisasterrecoveryGetallParamsWithContext(ctx context.Context) *PcloudLocationsDisasterrecoveryGetallParams {
	return &PcloudLocationsDisasterrecoveryGetallParams{
		Context: ctx,
	}
}

// NewPcloudLocationsDisasterrecoveryGetallParamsWithHTTPClient creates a new PcloudLocationsDisasterrecoveryGetallParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudLocationsDisasterrecoveryGetallParamsWithHTTPClient(client *http.Client) *PcloudLocationsDisasterrecoveryGetallParams {
	return &PcloudLocationsDisasterrecoveryGetallParams{
		HTTPClient: client,
	}
}

/*
PcloudLocationsDisasterrecoveryGetallParams contains all the parameters to send to the API endpoint

	for the pcloud locations disasterrecovery getall operation.

	Typically these are written to a http.Request.
*/
type PcloudLocationsDisasterrecoveryGetallParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud locations disasterrecovery getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudLocationsDisasterrecoveryGetallParams) WithDefaults() *PcloudLocationsDisasterrecoveryGetallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud locations disasterrecovery getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudLocationsDisasterrecoveryGetallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud locations disasterrecovery getall params
func (o *PcloudLocationsDisasterrecoveryGetallParams) WithTimeout(timeout time.Duration) *PcloudLocationsDisasterrecoveryGetallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud locations disasterrecovery getall params
func (o *PcloudLocationsDisasterrecoveryGetallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud locations disasterrecovery getall params
func (o *PcloudLocationsDisasterrecoveryGetallParams) WithContext(ctx context.Context) *PcloudLocationsDisasterrecoveryGetallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud locations disasterrecovery getall params
func (o *PcloudLocationsDisasterrecoveryGetallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud locations disasterrecovery getall params
func (o *PcloudLocationsDisasterrecoveryGetallParams) WithHTTPClient(client *http.Client) *PcloudLocationsDisasterrecoveryGetallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud locations disasterrecovery getall params
func (o *PcloudLocationsDisasterrecoveryGetallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudLocationsDisasterrecoveryGetallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
