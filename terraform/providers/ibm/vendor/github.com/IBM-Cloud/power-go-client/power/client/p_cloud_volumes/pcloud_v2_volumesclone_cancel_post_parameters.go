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

// NewPcloudV2VolumescloneCancelPostParams creates a new PcloudV2VolumescloneCancelPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudV2VolumescloneCancelPostParams() *PcloudV2VolumescloneCancelPostParams {
	return &PcloudV2VolumescloneCancelPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudV2VolumescloneCancelPostParamsWithTimeout creates a new PcloudV2VolumescloneCancelPostParams object
// with the ability to set a timeout on a request.
func NewPcloudV2VolumescloneCancelPostParamsWithTimeout(timeout time.Duration) *PcloudV2VolumescloneCancelPostParams {
	return &PcloudV2VolumescloneCancelPostParams{
		timeout: timeout,
	}
}

// NewPcloudV2VolumescloneCancelPostParamsWithContext creates a new PcloudV2VolumescloneCancelPostParams object
// with the ability to set a context for a request.
func NewPcloudV2VolumescloneCancelPostParamsWithContext(ctx context.Context) *PcloudV2VolumescloneCancelPostParams {
	return &PcloudV2VolumescloneCancelPostParams{
		Context: ctx,
	}
}

// NewPcloudV2VolumescloneCancelPostParamsWithHTTPClient creates a new PcloudV2VolumescloneCancelPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudV2VolumescloneCancelPostParamsWithHTTPClient(client *http.Client) *PcloudV2VolumescloneCancelPostParams {
	return &PcloudV2VolumescloneCancelPostParams{
		HTTPClient: client,
	}
}

/*
PcloudV2VolumescloneCancelPostParams contains all the parameters to send to the API endpoint

	for the pcloud v2 volumesclone cancel post operation.

	Typically these are written to a http.Request.
*/
type PcloudV2VolumescloneCancelPostParams struct {

	/* Body.

	   Parameters for cancelling a volumes-clone request
	*/
	Body *models.VolumesCloneCancel

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* VolumesCloneID.

	   Volumes Clone ID
	*/
	VolumesCloneID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud v2 volumesclone cancel post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudV2VolumescloneCancelPostParams) WithDefaults() *PcloudV2VolumescloneCancelPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud v2 volumesclone cancel post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudV2VolumescloneCancelPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud v2 volumesclone cancel post params
func (o *PcloudV2VolumescloneCancelPostParams) WithTimeout(timeout time.Duration) *PcloudV2VolumescloneCancelPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud v2 volumesclone cancel post params
func (o *PcloudV2VolumescloneCancelPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud v2 volumesclone cancel post params
func (o *PcloudV2VolumescloneCancelPostParams) WithContext(ctx context.Context) *PcloudV2VolumescloneCancelPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud v2 volumesclone cancel post params
func (o *PcloudV2VolumescloneCancelPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud v2 volumesclone cancel post params
func (o *PcloudV2VolumescloneCancelPostParams) WithHTTPClient(client *http.Client) *PcloudV2VolumescloneCancelPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud v2 volumesclone cancel post params
func (o *PcloudV2VolumescloneCancelPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud v2 volumesclone cancel post params
func (o *PcloudV2VolumescloneCancelPostParams) WithBody(body *models.VolumesCloneCancel) *PcloudV2VolumescloneCancelPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud v2 volumesclone cancel post params
func (o *PcloudV2VolumescloneCancelPostParams) SetBody(body *models.VolumesCloneCancel) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud v2 volumesclone cancel post params
func (o *PcloudV2VolumescloneCancelPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudV2VolumescloneCancelPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud v2 volumesclone cancel post params
func (o *PcloudV2VolumescloneCancelPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithVolumesCloneID adds the volumesCloneID to the pcloud v2 volumesclone cancel post params
func (o *PcloudV2VolumescloneCancelPostParams) WithVolumesCloneID(volumesCloneID string) *PcloudV2VolumescloneCancelPostParams {
	o.SetVolumesCloneID(volumesCloneID)
	return o
}

// SetVolumesCloneID adds the volumesCloneId to the pcloud v2 volumesclone cancel post params
func (o *PcloudV2VolumescloneCancelPostParams) SetVolumesCloneID(volumesCloneID string) {
	o.VolumesCloneID = volumesCloneID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudV2VolumescloneCancelPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param volumes_clone_id
	if err := r.SetPathParam("volumes_clone_id", o.VolumesCloneID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
