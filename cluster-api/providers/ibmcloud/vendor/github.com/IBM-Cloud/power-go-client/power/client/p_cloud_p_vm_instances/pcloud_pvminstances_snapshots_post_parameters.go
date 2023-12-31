// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

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

// NewPcloudPvminstancesSnapshotsPostParams creates a new PcloudPvminstancesSnapshotsPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudPvminstancesSnapshotsPostParams() *PcloudPvminstancesSnapshotsPostParams {
	return &PcloudPvminstancesSnapshotsPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesSnapshotsPostParamsWithTimeout creates a new PcloudPvminstancesSnapshotsPostParams object
// with the ability to set a timeout on a request.
func NewPcloudPvminstancesSnapshotsPostParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesSnapshotsPostParams {
	return &PcloudPvminstancesSnapshotsPostParams{
		timeout: timeout,
	}
}

// NewPcloudPvminstancesSnapshotsPostParamsWithContext creates a new PcloudPvminstancesSnapshotsPostParams object
// with the ability to set a context for a request.
func NewPcloudPvminstancesSnapshotsPostParamsWithContext(ctx context.Context) *PcloudPvminstancesSnapshotsPostParams {
	return &PcloudPvminstancesSnapshotsPostParams{
		Context: ctx,
	}
}

// NewPcloudPvminstancesSnapshotsPostParamsWithHTTPClient creates a new PcloudPvminstancesSnapshotsPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudPvminstancesSnapshotsPostParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesSnapshotsPostParams {
	return &PcloudPvminstancesSnapshotsPostParams{
		HTTPClient: client,
	}
}

/*
PcloudPvminstancesSnapshotsPostParams contains all the parameters to send to the API endpoint

	for the pcloud pvminstances snapshots post operation.

	Typically these are written to a http.Request.
*/
type PcloudPvminstancesSnapshotsPostParams struct {

	/* Body.

	   PVM Instance snapshot create parameters
	*/
	Body *models.SnapshotCreate

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* PvmInstanceID.

	   PCloud PVM Instance ID
	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud pvminstances snapshots post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesSnapshotsPostParams) WithDefaults() *PcloudPvminstancesSnapshotsPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud pvminstances snapshots post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesSnapshotsPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud pvminstances snapshots post params
func (o *PcloudPvminstancesSnapshotsPostParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesSnapshotsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances snapshots post params
func (o *PcloudPvminstancesSnapshotsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances snapshots post params
func (o *PcloudPvminstancesSnapshotsPostParams) WithContext(ctx context.Context) *PcloudPvminstancesSnapshotsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances snapshots post params
func (o *PcloudPvminstancesSnapshotsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances snapshots post params
func (o *PcloudPvminstancesSnapshotsPostParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesSnapshotsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances snapshots post params
func (o *PcloudPvminstancesSnapshotsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud pvminstances snapshots post params
func (o *PcloudPvminstancesSnapshotsPostParams) WithBody(body *models.SnapshotCreate) *PcloudPvminstancesSnapshotsPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud pvminstances snapshots post params
func (o *PcloudPvminstancesSnapshotsPostParams) SetBody(body *models.SnapshotCreate) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances snapshots post params
func (o *PcloudPvminstancesSnapshotsPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesSnapshotsPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances snapshots post params
func (o *PcloudPvminstancesSnapshotsPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances snapshots post params
func (o *PcloudPvminstancesSnapshotsPostParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesSnapshotsPostParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances snapshots post params
func (o *PcloudPvminstancesSnapshotsPostParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesSnapshotsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
