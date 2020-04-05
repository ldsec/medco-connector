// Code generated by go-swagger; DO NOT EDIT.

package medco_node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "github.com/ldsec/medco-connector/restapi/models"
)

// GetStatusHandlerFunc turns a function with the right signature into a get status handler
type GetStatusHandlerFunc func(GetStatusParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStatusHandlerFunc) Handle(params GetStatusParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetStatusHandler interface for that can handle valid get status params
type GetStatusHandler interface {
	Handle(GetStatusParams, *models.User) middleware.Responder
}

// NewGetStatus creates a new http.Handler for the get status operation
func NewGetStatus(ctx *middleware.Context, handler GetStatusHandler) *GetStatus {
	return &GetStatus{Context: ctx, Handler: handler}
}

/*GetStatus swagger:route GET /node/status medco-node getStatus

Get info about node status.

*/
type GetStatus struct {
	Context *middleware.Context
	Handler GetStatusHandler
}

func (o *GetStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetStatusParams()

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

// GetStatusDefaultBody get status default body
// swagger:model GetStatusDefaultBody
type GetStatusDefaultBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get status default body
func (o *GetStatusDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetStatusDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStatusDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetStatusDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetStatusOKBody get status o k body
// swagger:model GetStatusOKBody
type GetStatusOKBody struct {

	// message
	Message string `json:"message,omitempty"`

	// status o k
	StatusOK bool `json:"statusOK,omitempty"`
}

// Validate validates this get status o k body
func (o *GetStatusOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStatusOKBody) UnmarshalBinary(b []byte) error {
	var res GetStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
