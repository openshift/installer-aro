// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

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

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudNetworksPortsPostParams creates a new PcloudNetworksPortsPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudNetworksPortsPostParams() *PcloudNetworksPortsPostParams {
	return &PcloudNetworksPortsPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudNetworksPortsPostParamsWithTimeout creates a new PcloudNetworksPortsPostParams object
// with the ability to set a timeout on a request.
func NewPcloudNetworksPortsPostParamsWithTimeout(timeout time.Duration) *PcloudNetworksPortsPostParams {
	return &PcloudNetworksPortsPostParams{
		timeout: timeout,
	}
}

// NewPcloudNetworksPortsPostParamsWithContext creates a new PcloudNetworksPortsPostParams object
// with the ability to set a context for a request.
func NewPcloudNetworksPortsPostParamsWithContext(ctx context.Context) *PcloudNetworksPortsPostParams {
	return &PcloudNetworksPortsPostParams{
		Context: ctx,
	}
}

// NewPcloudNetworksPortsPostParamsWithHTTPClient creates a new PcloudNetworksPortsPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudNetworksPortsPostParamsWithHTTPClient(client *http.Client) *PcloudNetworksPortsPostParams {
	return &PcloudNetworksPortsPostParams{
		HTTPClient: client,
	}
}

/*
PcloudNetworksPortsPostParams contains all the parameters to send to the API endpoint

	for the pcloud networks ports post operation.

	Typically these are written to a http.Request.
*/
type PcloudNetworksPortsPostParams struct {

	/* Body.

	   Create a Network Port
	*/
	Body *models.NetworkPortCreate

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud networks ports post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudNetworksPortsPostParams) WithDefaults() *PcloudNetworksPortsPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud networks ports post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudNetworksPortsPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud networks ports post params
func (o *PcloudNetworksPortsPostParams) WithTimeout(timeout time.Duration) *PcloudNetworksPortsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud networks ports post params
func (o *PcloudNetworksPortsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud networks ports post params
func (o *PcloudNetworksPortsPostParams) WithContext(ctx context.Context) *PcloudNetworksPortsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud networks ports post params
func (o *PcloudNetworksPortsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud networks ports post params
func (o *PcloudNetworksPortsPostParams) WithHTTPClient(client *http.Client) *PcloudNetworksPortsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud networks ports post params
func (o *PcloudNetworksPortsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud networks ports post params
func (o *PcloudNetworksPortsPostParams) WithBody(body *models.NetworkPortCreate) *PcloudNetworksPortsPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud networks ports post params
func (o *PcloudNetworksPortsPostParams) SetBody(body *models.NetworkPortCreate) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud networks ports post params
func (o *PcloudNetworksPortsPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudNetworksPortsPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud networks ports post params
func (o *PcloudNetworksPortsPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithNetworkID adds the networkID to the pcloud networks ports post params
func (o *PcloudNetworksPortsPostParams) WithNetworkID(networkID string) *PcloudNetworksPortsPostParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the pcloud networks ports post params
func (o *PcloudNetworksPortsPostParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudNetworksPortsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
