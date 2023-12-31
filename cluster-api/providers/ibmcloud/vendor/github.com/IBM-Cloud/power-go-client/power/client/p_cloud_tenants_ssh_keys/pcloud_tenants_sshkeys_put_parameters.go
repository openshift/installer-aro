// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_tenants_ssh_keys

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

// NewPcloudTenantsSshkeysPutParams creates a new PcloudTenantsSshkeysPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudTenantsSshkeysPutParams() *PcloudTenantsSshkeysPutParams {
	return &PcloudTenantsSshkeysPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudTenantsSshkeysPutParamsWithTimeout creates a new PcloudTenantsSshkeysPutParams object
// with the ability to set a timeout on a request.
func NewPcloudTenantsSshkeysPutParamsWithTimeout(timeout time.Duration) *PcloudTenantsSshkeysPutParams {
	return &PcloudTenantsSshkeysPutParams{
		timeout: timeout,
	}
}

// NewPcloudTenantsSshkeysPutParamsWithContext creates a new PcloudTenantsSshkeysPutParams object
// with the ability to set a context for a request.
func NewPcloudTenantsSshkeysPutParamsWithContext(ctx context.Context) *PcloudTenantsSshkeysPutParams {
	return &PcloudTenantsSshkeysPutParams{
		Context: ctx,
	}
}

// NewPcloudTenantsSshkeysPutParamsWithHTTPClient creates a new PcloudTenantsSshkeysPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudTenantsSshkeysPutParamsWithHTTPClient(client *http.Client) *PcloudTenantsSshkeysPutParams {
	return &PcloudTenantsSshkeysPutParams{
		HTTPClient: client,
	}
}

/*
PcloudTenantsSshkeysPutParams contains all the parameters to send to the API endpoint

	for the pcloud tenants sshkeys put operation.

	Typically these are written to a http.Request.
*/
type PcloudTenantsSshkeysPutParams struct {

	/* Body.

	   Parameters for updating a Tenant's SSH Key
	*/
	Body *models.SSHKey

	/* SshkeyName.

	   SSH key name for a pcloud tenant
	*/
	SshkeyName string

	/* TenantID.

	   Tenant ID of a pcloud tenant
	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud tenants sshkeys put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudTenantsSshkeysPutParams) WithDefaults() *PcloudTenantsSshkeysPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud tenants sshkeys put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudTenantsSshkeysPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud tenants sshkeys put params
func (o *PcloudTenantsSshkeysPutParams) WithTimeout(timeout time.Duration) *PcloudTenantsSshkeysPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud tenants sshkeys put params
func (o *PcloudTenantsSshkeysPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud tenants sshkeys put params
func (o *PcloudTenantsSshkeysPutParams) WithContext(ctx context.Context) *PcloudTenantsSshkeysPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud tenants sshkeys put params
func (o *PcloudTenantsSshkeysPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud tenants sshkeys put params
func (o *PcloudTenantsSshkeysPutParams) WithHTTPClient(client *http.Client) *PcloudTenantsSshkeysPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud tenants sshkeys put params
func (o *PcloudTenantsSshkeysPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud tenants sshkeys put params
func (o *PcloudTenantsSshkeysPutParams) WithBody(body *models.SSHKey) *PcloudTenantsSshkeysPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud tenants sshkeys put params
func (o *PcloudTenantsSshkeysPutParams) SetBody(body *models.SSHKey) {
	o.Body = body
}

// WithSshkeyName adds the sshkeyName to the pcloud tenants sshkeys put params
func (o *PcloudTenantsSshkeysPutParams) WithSshkeyName(sshkeyName string) *PcloudTenantsSshkeysPutParams {
	o.SetSshkeyName(sshkeyName)
	return o
}

// SetSshkeyName adds the sshkeyName to the pcloud tenants sshkeys put params
func (o *PcloudTenantsSshkeysPutParams) SetSshkeyName(sshkeyName string) {
	o.SshkeyName = sshkeyName
}

// WithTenantID adds the tenantID to the pcloud tenants sshkeys put params
func (o *PcloudTenantsSshkeysPutParams) WithTenantID(tenantID string) *PcloudTenantsSshkeysPutParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the pcloud tenants sshkeys put params
func (o *PcloudTenantsSshkeysPutParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudTenantsSshkeysPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param sshkey_name
	if err := r.SetPathParam("sshkey_name", o.SshkeyName); err != nil {
		return err
	}

	// path param tenant_id
	if err := r.SetPathParam("tenant_id", o.TenantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
