// Code generated by go-swagger; DO NOT EDIT.

package survival_analysis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSurvivalAnalysisParams creates a new GetSurvivalAnalysisParams object
// with the default values initialized.
func NewGetSurvivalAnalysisParams() *GetSurvivalAnalysisParams {
	var ()
	return &GetSurvivalAnalysisParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSurvivalAnalysisParamsWithTimeout creates a new GetSurvivalAnalysisParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSurvivalAnalysisParamsWithTimeout(timeout time.Duration) *GetSurvivalAnalysisParams {
	var ()
	return &GetSurvivalAnalysisParams{

		timeout: timeout,
	}
}

// NewGetSurvivalAnalysisParamsWithContext creates a new GetSurvivalAnalysisParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSurvivalAnalysisParamsWithContext(ctx context.Context) *GetSurvivalAnalysisParams {
	var ()
	return &GetSurvivalAnalysisParams{

		Context: ctx,
	}
}

// NewGetSurvivalAnalysisParamsWithHTTPClient creates a new GetSurvivalAnalysisParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSurvivalAnalysisParamsWithHTTPClient(client *http.Client) *GetSurvivalAnalysisParams {
	var ()
	return &GetSurvivalAnalysisParams{
		HTTPClient: client,
	}
}

/*GetSurvivalAnalysisParams contains all the parameters to send to the API endpoint
for the get survival analysis operation typically these are written to a http.Request
*/
type GetSurvivalAnalysisParams struct {

	/*Body
	  User public key, patient list and time codes strings for the survival analysis

	*/
	Body GetSurvivalAnalysisBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get survival analysis params
func (o *GetSurvivalAnalysisParams) WithTimeout(timeout time.Duration) *GetSurvivalAnalysisParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get survival analysis params
func (o *GetSurvivalAnalysisParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get survival analysis params
func (o *GetSurvivalAnalysisParams) WithContext(ctx context.Context) *GetSurvivalAnalysisParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get survival analysis params
func (o *GetSurvivalAnalysisParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get survival analysis params
func (o *GetSurvivalAnalysisParams) WithHTTPClient(client *http.Client) *GetSurvivalAnalysisParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get survival analysis params
func (o *GetSurvivalAnalysisParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get survival analysis params
func (o *GetSurvivalAnalysisParams) WithBody(body GetSurvivalAnalysisBody) *GetSurvivalAnalysisParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get survival analysis params
func (o *GetSurvivalAnalysisParams) SetBody(body GetSurvivalAnalysisBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetSurvivalAnalysisParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}