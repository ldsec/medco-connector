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

// NewGetExploreQueryParams creates a new GetExploreQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExploreQueryParams() *GetExploreQueryParams {
	return &GetExploreQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExploreQueryParamsWithTimeout creates a new GetExploreQueryParams object
// with the ability to set a timeout on a request.
func NewGetExploreQueryParamsWithTimeout(timeout time.Duration) *GetExploreQueryParams {
	return &GetExploreQueryParams{
		timeout: timeout,
	}
}

// NewGetExploreQueryParamsWithContext creates a new GetExploreQueryParams object
// with the ability to set a context for a request.
func NewGetExploreQueryParamsWithContext(ctx context.Context) *GetExploreQueryParams {
	return &GetExploreQueryParams{
		Context: ctx,
	}
}

// NewGetExploreQueryParamsWithHTTPClient creates a new GetExploreQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExploreQueryParamsWithHTTPClient(client *http.Client) *GetExploreQueryParams {
	return &GetExploreQueryParams{
		HTTPClient: client,
	}
}

/* GetExploreQueryParams contains all the parameters to send to the API endpoint
   for the get explore query operation.

   Typically these are written to a http.Request.
*/
type GetExploreQueryParams struct {

	/* QueryID.

	   Query ID
	*/
	QueryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get explore query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExploreQueryParams) WithDefaults() *GetExploreQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get explore query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExploreQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get explore query params
func (o *GetExploreQueryParams) WithTimeout(timeout time.Duration) *GetExploreQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get explore query params
func (o *GetExploreQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get explore query params
func (o *GetExploreQueryParams) WithContext(ctx context.Context) *GetExploreQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get explore query params
func (o *GetExploreQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get explore query params
func (o *GetExploreQueryParams) WithHTTPClient(client *http.Client) *GetExploreQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get explore query params
func (o *GetExploreQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQueryID adds the queryID to the get explore query params
func (o *GetExploreQueryParams) WithQueryID(queryID string) *GetExploreQueryParams {
	o.SetQueryID(queryID)
	return o
}

// SetQueryID adds the queryId to the get explore query params
func (o *GetExploreQueryParams) SetQueryID(queryID string) {
	o.QueryID = queryID
}

// WriteToRequest writes these params to a swagger request
func (o *GetExploreQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param queryId
	if err := r.SetPathParam("queryId", o.QueryID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
