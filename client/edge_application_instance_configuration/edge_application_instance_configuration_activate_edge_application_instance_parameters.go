// Code generated by go-swagger; DO NOT EDIT.

package edge_application_instance_configuration

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

// NewEdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams creates a new EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams() *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams {
	return &EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParamsWithTimeout creates a new EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams object
// with the ability to set a timeout on a request.
func NewEdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParamsWithTimeout(timeout time.Duration) *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams {
	return &EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams{
		timeout: timeout,
	}
}

// NewEdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParamsWithContext creates a new EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams object
// with the ability to set a context for a request.
func NewEdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParamsWithContext(ctx context.Context) *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams {
	return &EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams{
		Context: ctx,
	}
}

// NewEdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParamsWithHTTPClient creates a new EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParamsWithHTTPClient(client *http.Client) *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams {
	return &EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams{
		HTTPClient: client,
	}
}

/*
EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams contains all the parameters to send to the API endpoint

	for the edge application instance configuration activate edge application instance operation.

	Typically these are written to a http.Request.
*/
type EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the app instance
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge application instance configuration activate edge application instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams) WithDefaults() *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge application instance configuration activate edge application instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge application instance configuration activate edge application instance params
func (o *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams) WithTimeout(timeout time.Duration) *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge application instance configuration activate edge application instance params
func (o *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge application instance configuration activate edge application instance params
func (o *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams) WithContext(ctx context.Context) *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge application instance configuration activate edge application instance params
func (o *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge application instance configuration activate edge application instance params
func (o *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams) WithHTTPClient(client *http.Client) *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge application instance configuration activate edge application instance params
func (o *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge application instance configuration activate edge application instance params
func (o *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams) WithXRequestID(xRequestID *string) *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge application instance configuration activate edge application instance params
func (o *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge application instance configuration activate edge application instance params
func (o *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams) WithID(id string) *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge application instance configuration activate edge application instance params
func (o *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}