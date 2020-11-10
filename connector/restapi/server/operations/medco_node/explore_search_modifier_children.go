// Code generated by go-swagger; DO NOT EDIT.

package medco_node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ldsec/medco/connector/restapi/models"
)

// ExploreSearchModifierChildrenHandlerFunc turns a function with the right signature into a explore search modifier children handler
type ExploreSearchModifierChildrenHandlerFunc func(ExploreSearchModifierChildrenParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn ExploreSearchModifierChildrenHandlerFunc) Handle(params ExploreSearchModifierChildrenParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// ExploreSearchModifierChildrenHandler interface for that can handle valid explore search modifier children params
type ExploreSearchModifierChildrenHandler interface {
	Handle(ExploreSearchModifierChildrenParams, *models.User) middleware.Responder
}

// NewExploreSearchModifierChildren creates a new http.Handler for the explore search modifier children operation
func NewExploreSearchModifierChildren(ctx *middleware.Context, handler ExploreSearchModifierChildrenHandler) *ExploreSearchModifierChildren {
	return &ExploreSearchModifierChildren{Context: ctx, Handler: handler}
}

/*ExploreSearchModifierChildren swagger:route POST /node/explore/search/modifier-children medco-node exploreSearchModifierChildren

Returns the modifier children

*/
type ExploreSearchModifierChildren struct {
	Context *middleware.Context
	Handler ExploreSearchModifierChildrenHandler
}

func (o *ExploreSearchModifierChildren) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewExploreSearchModifierChildrenParams()

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

// ExploreSearchModifierChildrenDefaultBody explore search modifier children default body
//
// swagger:model ExploreSearchModifierChildrenDefaultBody
type ExploreSearchModifierChildrenDefaultBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this explore search modifier children default body
func (o *ExploreSearchModifierChildrenDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExploreSearchModifierChildrenDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExploreSearchModifierChildrenDefaultBody) UnmarshalBinary(b []byte) error {
	var res ExploreSearchModifierChildrenDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ExploreSearchModifierChildrenOKBody explore search modifier children o k body
//
// swagger:model ExploreSearchModifierChildrenOKBody
type ExploreSearchModifierChildrenOKBody struct {

	// results
	Results []*models.ExploreSearchResultElement `json:"results"`

	// search
	Search *models.ExploreSearchModifierChildren `json:"search,omitempty"`
}

// Validate validates this explore search modifier children o k body
func (o *ExploreSearchModifierChildrenOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSearch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExploreSearchModifierChildrenOKBody) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(o.Results) { // not required
		return nil
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("exploreSearchModifierChildrenOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ExploreSearchModifierChildrenOKBody) validateSearch(formats strfmt.Registry) error {

	if swag.IsZero(o.Search) { // not required
		return nil
	}

	if o.Search != nil {
		if err := o.Search.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exploreSearchModifierChildrenOK" + "." + "search")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExploreSearchModifierChildrenOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExploreSearchModifierChildrenOKBody) UnmarshalBinary(b []byte) error {
	var res ExploreSearchModifierChildrenOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
