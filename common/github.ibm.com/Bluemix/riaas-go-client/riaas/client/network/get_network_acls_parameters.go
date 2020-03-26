// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNetworkAclsParams creates a new GetNetworkAclsParams object
// with the default values initialized.
func NewGetNetworkAclsParams() *GetNetworkAclsParams {
	var (
		limitDefault = int32(50)
	)
	return &GetNetworkAclsParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkAclsParamsWithTimeout creates a new GetNetworkAclsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkAclsParamsWithTimeout(timeout time.Duration) *GetNetworkAclsParams {
	var (
		limitDefault = int32(50)
	)
	return &GetNetworkAclsParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewGetNetworkAclsParamsWithContext creates a new GetNetworkAclsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkAclsParamsWithContext(ctx context.Context) *GetNetworkAclsParams {
	var (
		limitDefault = int32(50)
	)
	return &GetNetworkAclsParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewGetNetworkAclsParamsWithHTTPClient creates a new GetNetworkAclsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkAclsParamsWithHTTPClient(client *http.Client) *GetNetworkAclsParams {
	var (
		limitDefault = int32(50)
	)
	return &GetNetworkAclsParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*GetNetworkAclsParams contains all the parameters to send to the API endpoint
for the get network acls operation typically these are written to a http.Request
*/
type GetNetworkAclsParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*Limit
	  The number of resources to return on a page

	*/
	Limit *int32
	/*ResourceGroupID
	  Filters the collection to resources within the resource group of the specified identifier

	*/
	ResourceGroupID *string
	/*Start
	  A server-supplied token determining what resource to start the page on

	*/
	Start *string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network acls params
func (o *GetNetworkAclsParams) WithTimeout(timeout time.Duration) *GetNetworkAclsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network acls params
func (o *GetNetworkAclsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network acls params
func (o *GetNetworkAclsParams) WithContext(ctx context.Context) *GetNetworkAclsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network acls params
func (o *GetNetworkAclsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network acls params
func (o *GetNetworkAclsParams) WithHTTPClient(client *http.Client) *GetNetworkAclsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network acls params
func (o *GetNetworkAclsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get network acls params
func (o *GetNetworkAclsParams) WithGeneration(generation int64) *GetNetworkAclsParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get network acls params
func (o *GetNetworkAclsParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithLimit adds the limit to the get network acls params
func (o *GetNetworkAclsParams) WithLimit(limit *int32) *GetNetworkAclsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get network acls params
func (o *GetNetworkAclsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithResourceGroupID adds the resourceGroupID to the get network acls params
func (o *GetNetworkAclsParams) WithResourceGroupID(resourceGroupID *string) *GetNetworkAclsParams {
	o.SetResourceGroupID(resourceGroupID)
	return o
}

// SetResourceGroupID adds the resourceGroupId to the get network acls params
func (o *GetNetworkAclsParams) SetResourceGroupID(resourceGroupID *string) {
	o.ResourceGroupID = resourceGroupID
}

// WithStart adds the start to the get network acls params
func (o *GetNetworkAclsParams) WithStart(start *string) *GetNetworkAclsParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get network acls params
func (o *GetNetworkAclsParams) SetStart(start *string) {
	o.Start = start
}

// WithVersion adds the version to the get network acls params
func (o *GetNetworkAclsParams) WithVersion(version string) *GetNetworkAclsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get network acls params
func (o *GetNetworkAclsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkAclsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.ResourceGroupID != nil {

		// query param resource_group.id
		var qrResourceGroupID string
		if o.ResourceGroupID != nil {
			qrResourceGroupID = *o.ResourceGroupID
		}
		qResourceGroupID := qrResourceGroupID
		if qResourceGroupID != "" {
			if err := r.SetQueryParam("resource_group.id", qResourceGroupID); err != nil {
				return err
			}
		}

	}

	if o.Start != nil {

		// query param start
		var qrStart string
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := qrStart
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}