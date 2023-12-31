// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_cloud_connections

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

// NewPcloudCloudconnectionsPostParams creates a new PcloudCloudconnectionsPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudCloudconnectionsPostParams() *PcloudCloudconnectionsPostParams {
	return &PcloudCloudconnectionsPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudCloudconnectionsPostParamsWithTimeout creates a new PcloudCloudconnectionsPostParams object
// with the ability to set a timeout on a request.
func NewPcloudCloudconnectionsPostParamsWithTimeout(timeout time.Duration) *PcloudCloudconnectionsPostParams {
	return &PcloudCloudconnectionsPostParams{
		timeout: timeout,
	}
}

// NewPcloudCloudconnectionsPostParamsWithContext creates a new PcloudCloudconnectionsPostParams object
// with the ability to set a context for a request.
func NewPcloudCloudconnectionsPostParamsWithContext(ctx context.Context) *PcloudCloudconnectionsPostParams {
	return &PcloudCloudconnectionsPostParams{
		Context: ctx,
	}
}

// NewPcloudCloudconnectionsPostParamsWithHTTPClient creates a new PcloudCloudconnectionsPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudCloudconnectionsPostParamsWithHTTPClient(client *http.Client) *PcloudCloudconnectionsPostParams {
	return &PcloudCloudconnectionsPostParams{
		HTTPClient: client,
	}
}

/*
PcloudCloudconnectionsPostParams contains all the parameters to send to the API endpoint

	for the pcloud cloudconnections post operation.

	Typically these are written to a http.Request.
*/
type PcloudCloudconnectionsPostParams struct {

	/* Body.

	   Parameters for the creation of a new cloud connection
	*/
	Body *models.CloudConnectionCreate

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud cloudconnections post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudconnectionsPostParams) WithDefaults() *PcloudCloudconnectionsPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud cloudconnections post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudconnectionsPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud cloudconnections post params
func (o *PcloudCloudconnectionsPostParams) WithTimeout(timeout time.Duration) *PcloudCloudconnectionsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud cloudconnections post params
func (o *PcloudCloudconnectionsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud cloudconnections post params
func (o *PcloudCloudconnectionsPostParams) WithContext(ctx context.Context) *PcloudCloudconnectionsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud cloudconnections post params
func (o *PcloudCloudconnectionsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud cloudconnections post params
func (o *PcloudCloudconnectionsPostParams) WithHTTPClient(client *http.Client) *PcloudCloudconnectionsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud cloudconnections post params
func (o *PcloudCloudconnectionsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud cloudconnections post params
func (o *PcloudCloudconnectionsPostParams) WithBody(body *models.CloudConnectionCreate) *PcloudCloudconnectionsPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud cloudconnections post params
func (o *PcloudCloudconnectionsPostParams) SetBody(body *models.CloudConnectionCreate) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud cloudconnections post params
func (o *PcloudCloudconnectionsPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudCloudconnectionsPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud cloudconnections post params
func (o *PcloudCloudconnectionsPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudCloudconnectionsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
