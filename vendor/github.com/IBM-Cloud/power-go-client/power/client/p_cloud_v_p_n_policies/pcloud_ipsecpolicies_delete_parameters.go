// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_policies

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

// NewPcloudIpsecpoliciesDeleteParams creates a new PcloudIpsecpoliciesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudIpsecpoliciesDeleteParams() *PcloudIpsecpoliciesDeleteParams {
	return &PcloudIpsecpoliciesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudIpsecpoliciesDeleteParamsWithTimeout creates a new PcloudIpsecpoliciesDeleteParams object
// with the ability to set a timeout on a request.
func NewPcloudIpsecpoliciesDeleteParamsWithTimeout(timeout time.Duration) *PcloudIpsecpoliciesDeleteParams {
	return &PcloudIpsecpoliciesDeleteParams{
		timeout: timeout,
	}
}

// NewPcloudIpsecpoliciesDeleteParamsWithContext creates a new PcloudIpsecpoliciesDeleteParams object
// with the ability to set a context for a request.
func NewPcloudIpsecpoliciesDeleteParamsWithContext(ctx context.Context) *PcloudIpsecpoliciesDeleteParams {
	return &PcloudIpsecpoliciesDeleteParams{
		Context: ctx,
	}
}

// NewPcloudIpsecpoliciesDeleteParamsWithHTTPClient creates a new PcloudIpsecpoliciesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudIpsecpoliciesDeleteParamsWithHTTPClient(client *http.Client) *PcloudIpsecpoliciesDeleteParams {
	return &PcloudIpsecpoliciesDeleteParams{
		HTTPClient: client,
	}
}

/*
PcloudIpsecpoliciesDeleteParams contains all the parameters to send to the API endpoint

	for the pcloud ipsecpolicies delete operation.

	Typically these are written to a http.Request.
*/
type PcloudIpsecpoliciesDeleteParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* IpsecPolicyID.

	   ID of a IPSec Policy
	*/
	IpsecPolicyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud ipsecpolicies delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudIpsecpoliciesDeleteParams) WithDefaults() *PcloudIpsecpoliciesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud ipsecpolicies delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudIpsecpoliciesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud ipsecpolicies delete params
func (o *PcloudIpsecpoliciesDeleteParams) WithTimeout(timeout time.Duration) *PcloudIpsecpoliciesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud ipsecpolicies delete params
func (o *PcloudIpsecpoliciesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud ipsecpolicies delete params
func (o *PcloudIpsecpoliciesDeleteParams) WithContext(ctx context.Context) *PcloudIpsecpoliciesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud ipsecpolicies delete params
func (o *PcloudIpsecpoliciesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud ipsecpolicies delete params
func (o *PcloudIpsecpoliciesDeleteParams) WithHTTPClient(client *http.Client) *PcloudIpsecpoliciesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud ipsecpolicies delete params
func (o *PcloudIpsecpoliciesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud ipsecpolicies delete params
func (o *PcloudIpsecpoliciesDeleteParams) WithCloudInstanceID(cloudInstanceID string) *PcloudIpsecpoliciesDeleteParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud ipsecpolicies delete params
func (o *PcloudIpsecpoliciesDeleteParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithIpsecPolicyID adds the ipsecPolicyID to the pcloud ipsecpolicies delete params
func (o *PcloudIpsecpoliciesDeleteParams) WithIpsecPolicyID(ipsecPolicyID string) *PcloudIpsecpoliciesDeleteParams {
	o.SetIpsecPolicyID(ipsecPolicyID)
	return o
}

// SetIpsecPolicyID adds the ipsecPolicyId to the pcloud ipsecpolicies delete params
func (o *PcloudIpsecpoliciesDeleteParams) SetIpsecPolicyID(ipsecPolicyID string) {
	o.IpsecPolicyID = ipsecPolicyID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudIpsecpoliciesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param ipsec_policy_id
	if err := r.SetPathParam("ipsec_policy_id", o.IpsecPolicyID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
