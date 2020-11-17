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

// NewPostCohortsParams creates a new PostCohortsParams object
// with the default values initialized.
func NewPostCohortsParams() *PostCohortsParams {
	var ()
	return &PostCohortsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCohortsParamsWithTimeout creates a new PostCohortsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCohortsParamsWithTimeout(timeout time.Duration) *PostCohortsParams {
	var ()
	return &PostCohortsParams{

		timeout: timeout,
	}
}

// NewPostCohortsParamsWithContext creates a new PostCohortsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCohortsParamsWithContext(ctx context.Context) *PostCohortsParams {
	var ()
	return &PostCohortsParams{

		Context: ctx,
	}
}

// NewPostCohortsParamsWithHTTPClient creates a new PostCohortsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCohortsParamsWithHTTPClient(client *http.Client) *PostCohortsParams {
	var ()
	return &PostCohortsParams{
		HTTPClient: client,
	}
}

/*PostCohortsParams contains all the parameters to send to the API endpoint
for the post cohorts operation typically these are written to a http.Request
*/
type PostCohortsParams struct {

	/*CohortRequest
	  Cohort that has been updated or created

	*/
	CohortRequest PostCohortsBody
	/*Name
	  Name of the cohort to update

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cohorts params
func (o *PostCohortsParams) WithTimeout(timeout time.Duration) *PostCohortsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cohorts params
func (o *PostCohortsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cohorts params
func (o *PostCohortsParams) WithContext(ctx context.Context) *PostCohortsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cohorts params
func (o *PostCohortsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cohorts params
func (o *PostCohortsParams) WithHTTPClient(client *http.Client) *PostCohortsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cohorts params
func (o *PostCohortsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCohortRequest adds the cohortRequest to the post cohorts params
func (o *PostCohortsParams) WithCohortRequest(cohortRequest PostCohortsBody) *PostCohortsParams {
	o.SetCohortRequest(cohortRequest)
	return o
}

// SetCohortRequest adds the cohortRequest to the post cohorts params
func (o *PostCohortsParams) SetCohortRequest(cohortRequest PostCohortsBody) {
	o.CohortRequest = cohortRequest
}

// WithName adds the name to the post cohorts params
func (o *PostCohortsParams) WithName(name string) *PostCohortsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post cohorts params
func (o *PostCohortsParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PostCohortsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CohortRequest); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
