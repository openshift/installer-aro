// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

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

// NewPcloudVolumesClonePostParams creates a new PcloudVolumesClonePostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudVolumesClonePostParams() *PcloudVolumesClonePostParams {
	return &PcloudVolumesClonePostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudVolumesClonePostParamsWithTimeout creates a new PcloudVolumesClonePostParams object
// with the ability to set a timeout on a request.
func NewPcloudVolumesClonePostParamsWithTimeout(timeout time.Duration) *PcloudVolumesClonePostParams {
	return &PcloudVolumesClonePostParams{
		timeout: timeout,
	}
}

// NewPcloudVolumesClonePostParamsWithContext creates a new PcloudVolumesClonePostParams object
// with the ability to set a context for a request.
func NewPcloudVolumesClonePostParamsWithContext(ctx context.Context) *PcloudVolumesClonePostParams {
	return &PcloudVolumesClonePostParams{
		Context: ctx,
	}
}

// NewPcloudVolumesClonePostParamsWithHTTPClient creates a new PcloudVolumesClonePostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudVolumesClonePostParamsWithHTTPClient(client *http.Client) *PcloudVolumesClonePostParams {
	return &PcloudVolumesClonePostParams{
		HTTPClient: client,
	}
}

/*
PcloudVolumesClonePostParams contains all the parameters to send to the API endpoint

	for the pcloud volumes clone post operation.

	Typically these are written to a http.Request.
*/
type PcloudVolumesClonePostParams struct {

	/* Body.

	   Parameters for the cloning of volumes
	*/
	Body *models.VolumesCloneRequest

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud volumes clone post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudVolumesClonePostParams) WithDefaults() *PcloudVolumesClonePostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud volumes clone post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudVolumesClonePostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud volumes clone post params
func (o *PcloudVolumesClonePostParams) WithTimeout(timeout time.Duration) *PcloudVolumesClonePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud volumes clone post params
func (o *PcloudVolumesClonePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud volumes clone post params
func (o *PcloudVolumesClonePostParams) WithContext(ctx context.Context) *PcloudVolumesClonePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud volumes clone post params
func (o *PcloudVolumesClonePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud volumes clone post params
func (o *PcloudVolumesClonePostParams) WithHTTPClient(client *http.Client) *PcloudVolumesClonePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud volumes clone post params
func (o *PcloudVolumesClonePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud volumes clone post params
func (o *PcloudVolumesClonePostParams) WithBody(body *models.VolumesCloneRequest) *PcloudVolumesClonePostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud volumes clone post params
func (o *PcloudVolumesClonePostParams) SetBody(body *models.VolumesCloneRequest) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud volumes clone post params
func (o *PcloudVolumesClonePostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudVolumesClonePostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud volumes clone post params
func (o *PcloudVolumesClonePostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudVolumesClonePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
