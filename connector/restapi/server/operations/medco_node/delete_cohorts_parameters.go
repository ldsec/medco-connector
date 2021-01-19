// Code generated by go-swagger; DO NOT EDIT.

package medco_node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewDeleteCohortsParams creates a new DeleteCohortsParams object
//
// There are no default values defined in the spec.
func NewDeleteCohortsParams() DeleteCohortsParams {

	return DeleteCohortsParams{}
}

// DeleteCohortsParams contains all the bound params for the delete cohorts operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteCohorts
type DeleteCohortsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Name of the cohort to delete
	  Required: true
	  Pattern: ^\w+$
	  In: path
	*/
	Name string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteCohortsParams() beforehand.
func (o *DeleteCohortsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindName binds and validates parameter Name from path.
func (o *DeleteCohortsParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Name = raw

	if err := o.validateName(formats); err != nil {
		return err
	}

	return nil
}

// validateName carries on validations for parameter Name
func (o *DeleteCohortsParams) validateName(formats strfmt.Registry) error {

	if err := validate.Pattern("name", "path", o.Name, `^\w+$`); err != nil {
		return err
	}

	return nil
}
