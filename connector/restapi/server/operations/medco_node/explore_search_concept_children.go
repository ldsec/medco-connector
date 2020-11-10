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

// ExploreSearchConceptChildrenHandlerFunc turns a function with the right signature into a explore search concept children handler
type ExploreSearchConceptChildrenHandlerFunc func(ExploreSearchConceptChildrenParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn ExploreSearchConceptChildrenHandlerFunc) Handle(params ExploreSearchConceptChildrenParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// ExploreSearchConceptChildrenHandler interface for that can handle valid explore search concept children params
type ExploreSearchConceptChildrenHandler interface {
	Handle(ExploreSearchConceptChildrenParams, *models.User) middleware.Responder
}

// NewExploreSearchConceptChildren creates a new http.Handler for the explore search concept children operation
func NewExploreSearchConceptChildren(ctx *middleware.Context, handler ExploreSearchConceptChildrenHandler) *ExploreSearchConceptChildren {
	return &ExploreSearchConceptChildren{Context: ctx, Handler: handler}
}

/*ExploreSearchConceptChildren swagger:route POST /node/explore/search/concept-children medco-node exploreSearchConceptChildren

Returns the concept children (both concepts and modifiers)

*/
type ExploreSearchConceptChildren struct {
	Context *middleware.Context
	Handler ExploreSearchConceptChildrenHandler
}

func (o *ExploreSearchConceptChildren) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewExploreSearchConceptChildrenParams()

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

// ExploreSearchConceptChildrenDefaultBody explore search concept children default body
//
// swagger:model ExploreSearchConceptChildrenDefaultBody
type ExploreSearchConceptChildrenDefaultBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this explore search concept children default body
func (o *ExploreSearchConceptChildrenDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExploreSearchConceptChildrenDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExploreSearchConceptChildrenDefaultBody) UnmarshalBinary(b []byte) error {
	var res ExploreSearchConceptChildrenDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ExploreSearchConceptChildrenOKBody explore search concept children o k body
//
// swagger:model ExploreSearchConceptChildrenOKBody
type ExploreSearchConceptChildrenOKBody struct {

	// results
	Results []*models.ExploreSearchResultElement `json:"results"`

	// search
	Search *models.ExploreSearchConceptChildren `json:"search,omitempty"`
}

// Validate validates this explore search concept children o k body
func (o *ExploreSearchConceptChildrenOKBody) Validate(formats strfmt.Registry) error {
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

func (o *ExploreSearchConceptChildrenOKBody) validateResults(formats strfmt.Registry) error {

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
					return ve.ValidateName("exploreSearchConceptChildrenOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ExploreSearchConceptChildrenOKBody) validateSearch(formats strfmt.Registry) error {

	if swag.IsZero(o.Search) { // not required
		return nil
	}

	if o.Search != nil {
		if err := o.Search.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exploreSearchConceptChildrenOK" + "." + "search")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExploreSearchConceptChildrenOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExploreSearchConceptChildrenOKBody) UnmarshalBinary(b []byte) error {
	var res ExploreSearchConceptChildrenOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
