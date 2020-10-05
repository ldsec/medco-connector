// Code generated by go-swagger; DO NOT EDIT.

package medco_node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ldsec/medco-connector/restapi/models"
)

// PostCohortsHandlerFunc turns a function with the right signature into a post cohorts handler
type PostCohortsHandlerFunc func(PostCohortsParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn PostCohortsHandlerFunc) Handle(params PostCohortsParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// PostCohortsHandler interface for that can handle valid post cohorts params
type PostCohortsHandler interface {
	Handle(PostCohortsParams, *models.User) middleware.Responder
}

// NewPostCohorts creates a new http.Handler for the post cohorts operation
func NewPostCohorts(ctx *middleware.Context, handler PostCohortsHandler) *PostCohorts {
	return &PostCohorts{Context: ctx, Handler: handler}
}

/*PostCohorts swagger:route POST /node/explore/cohorts medco-node postCohorts

Add a new cohort or update an existing one

*/
type PostCohorts struct {
	Context *middleware.Context
	Handler PostCohortsHandler
}

func (o *PostCohorts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostCohortsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostCohortsBody post cohorts body
//
// swagger:model PostCohortsBody
type PostCohortsBody struct {

	// cohort name
	CohortName string `json:"cohortName,omitempty"`

	// creation date
	CreationDate string `json:"creationDate,omitempty"`

	// patient set ID
	PatientSetID int64 `json:"patientSetID,omitempty"`

	// update date
	UpdateDate string `json:"updateDate,omitempty"`
}

// Validate validates this post cohorts body
func (o *PostCohortsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCohortsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCohortsBody) UnmarshalBinary(b []byte) error {
	var res PostCohortsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostCohortsDefaultBody post cohorts default body
//
// swagger:model PostCohortsDefaultBody
type PostCohortsDefaultBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this post cohorts default body
func (o *PostCohortsDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCohortsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCohortsDefaultBody) UnmarshalBinary(b []byte) error {
	var res PostCohortsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}