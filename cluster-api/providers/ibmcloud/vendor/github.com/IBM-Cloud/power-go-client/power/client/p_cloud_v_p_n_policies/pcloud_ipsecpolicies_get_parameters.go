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

// NewPcloudIpsecpoliciesGetParams creates a new PcloudIpsecpoliciesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudIpsecpoliciesGetParams() *PcloudIpsecpoliciesGetParams {
	return &PcloudIpsecpoliciesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudIpsecpoliciesGetParamsWithTimeout creates a new PcloudIpsecpoliciesGetParams object
// with the ability to set a timeout on a request.
func NewPcloudIpsecpoliciesGetParamsWithTimeout(timeout time.Duration) *PcloudIpsecpoliciesGetParams {
	return &PcloudIpsecpoliciesGetParams{
		timeout: timeout,
	}
}

// NewPcloudIpsecpoliciesGetParamsWithContext creates a new PcloudIpsecpoliciesGetParams object
// with the ability to set a context for a request.
func NewPcloudIpsecpoliciesGetParamsWithContext(ctx context.Context) *PcloudIpsecpoliciesGetParams {
	return &PcloudIpsecpoliciesGetParams{
		Context: ctx,
	}
}

// NewPcloudIpsecpoliciesGetParamsWithHTTPClient creates a new PcloudIpsecpoliciesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudIpsecpoliciesGetParamsWithHTTPClient(client *http.Client) *PcloudIpsecpoliciesGetParams {
	return &PcloudIpsecpoliciesGetParams{
		HTTPClient: client,
	}
}

/*
PcloudIpsecpoliciesGetParams contains all the parameters to send to the API endpoint

	for the pcloud ipsecpolicies get operation.

	Typically these are written to a http.Request.
*/
type PcloudIpsecpoliciesGetParams struct {

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

// WithDefaults hydrates default values in the pcloud ipsecpolicies get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudIpsecpoliciesGetParams) WithDefaults() *PcloudIpsecpoliciesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud ipsecpolicies get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudIpsecpoliciesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud ipsecpolicies get params
func (o *PcloudIpsecpoliciesGetParams) WithTimeout(timeout time.Duration) *PcloudIpsecpoliciesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud ipsecpolicies get params
func (o *PcloudIpsecpoliciesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud ipsecpolicies get params
func (o *PcloudIpsecpoliciesGetParams) WithContext(ctx context.Context) *PcloudIpsecpoliciesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud ipsecpolicies get params
func (o *PcloudIpsecpoliciesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud ipsecpolicies get params
func (o *PcloudIpsecpoliciesGetParams) WithHTTPClient(client *http.Client) *PcloudIpsecpoliciesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud ipsecpolicies get params
func (o *PcloudIpsecpoliciesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud ipsecpolicies get params
func (o *PcloudIpsecpoliciesGetParams) WithCloudInstanceID(cloudInstanceID string) *PcloudIpsecpoliciesGetParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud ipsecpolicies get params
func (o *PcloudIpsecpoliciesGetParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithIpsecPolicyID adds the ipsecPolicyID to the pcloud ipsecpolicies get params
func (o *PcloudIpsecpoliciesGetParams) WithIpsecPolicyID(ipsecPolicyID string) *PcloudIpsecpoliciesGetParams {
	o.SetIpsecPolicyID(ipsecPolicyID)
	return o
}

// SetIpsecPolicyID adds the ipsecPolicyId to the pcloud ipsecpolicies get params
func (o *PcloudIpsecpoliciesGetParams) SetIpsecPolicyID(ipsecPolicyID string) {
	o.IpsecPolicyID = ipsecPolicyID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudIpsecpoliciesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
