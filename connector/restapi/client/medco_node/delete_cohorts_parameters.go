// Code generated by go-swagger; DO NOT EDIT.

package medco_node

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

// NewDeleteCohortsParams creates a new DeleteCohortsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCohortsParams() *DeleteCohortsParams {
	return &DeleteCohortsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCohortsParamsWithTimeout creates a new DeleteCohortsParams object
// with the ability to set a timeout on a request.
func NewDeleteCohortsParamsWithTimeout(timeout time.Duration) *DeleteCohortsParams {
	return &DeleteCohortsParams{
		timeout: timeout,
	}
}

// NewDeleteCohortsParamsWithContext creates a new DeleteCohortsParams object
// with the ability to set a context for a request.
func NewDeleteCohortsParamsWithContext(ctx context.Context) *DeleteCohortsParams {
	return &DeleteCohortsParams{
		Context: ctx,
	}
}

// NewDeleteCohortsParamsWithHTTPClient creates a new DeleteCohortsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCohortsParamsWithHTTPClient(client *http.Client) *DeleteCohortsParams {
	return &DeleteCohortsParams{
		HTTPClient: client,
	}
}

/* DeleteCohortsParams contains all the parameters to send to the API endpoint
   for the delete cohorts operation.

   Typically these are written to a http.Request.
*/
type DeleteCohortsParams struct {

	/* Name.

	   Name of the cohort to delete
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete cohorts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCohortsParams) WithDefaults() *DeleteCohortsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete cohorts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCohortsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete cohorts params
func (o *DeleteCohortsParams) WithTimeout(timeout time.Duration) *DeleteCohortsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cohorts params
func (o *DeleteCohortsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cohorts params
func (o *DeleteCohortsParams) WithContext(ctx context.Context) *DeleteCohortsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cohorts params
func (o *DeleteCohortsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cohorts params
func (o *DeleteCohortsParams) WithHTTPClient(client *http.Client) *DeleteCohortsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cohorts params
func (o *DeleteCohortsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete cohorts params
func (o *DeleteCohortsParams) WithName(name string) *DeleteCohortsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete cohorts params
func (o *DeleteCohortsParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCohortsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
