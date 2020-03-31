// Code generated by go-swagger; DO NOT EDIT.

package survival_analysis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"

	models "github.com/ldsec/medco-connector/restapi/models"
)

// GetSurvivalAnalysisHandlerFunc turns a function with the right signature into a get survival analysis handler
type GetSurvivalAnalysisHandlerFunc func(GetSurvivalAnalysisParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSurvivalAnalysisHandlerFunc) Handle(params GetSurvivalAnalysisParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetSurvivalAnalysisHandler interface for that can handle valid get survival analysis params
type GetSurvivalAnalysisHandler interface {
	Handle(GetSurvivalAnalysisParams, *models.User) middleware.Responder
}

// NewGetSurvivalAnalysis creates a new http.Handler for the get survival analysis operation
func NewGetSurvivalAnalysis(ctx *middleware.Context, handler GetSurvivalAnalysisHandler) *GetSurvivalAnalysis {
	return &GetSurvivalAnalysis{Context: ctx, Handler: handler}
}

/*GetSurvivalAnalysis swagger:route POST /survival-analysis survival-analysis getSurvivalAnalysis

Send a query to run a survival analysis

*/
type GetSurvivalAnalysis struct {
	Context *middleware.Context
	Handler GetSurvivalAnalysisHandler
}

func (o *GetSurvivalAnalysis) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSurvivalAnalysisParams()

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

// GetSurvivalAnalysisBody get survival analysis body
// swagger:model GetSurvivalAnalysisBody
type GetSurvivalAnalysisBody struct {

	// ID
	ID string `json:"ID,omitempty"`

	// patient set ID
	PatientSetID string `json:"patientSetID,omitempty"`

	// time codes
	TimeCodes []string `json:"timeCodes"`

	// user public key
	// Pattern: ^[\w=-]+$
	UserPublicKey string `json:"userPublicKey,omitempty"`
}

// Validate validates this get survival analysis body
func (o *GetSurvivalAnalysisBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUserPublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSurvivalAnalysisBody) validateUserPublicKey(formats strfmt.Registry) error {

	if swag.IsZero(o.UserPublicKey) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"userPublicKey", "body", string(o.UserPublicKey), `^[\w=-]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSurvivalAnalysisBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSurvivalAnalysisBody) UnmarshalBinary(b []byte) error {
	var res GetSurvivalAnalysisBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetSurvivalAnalysisDefaultBody get survival analysis default body
// swagger:model GetSurvivalAnalysisDefaultBody
type GetSurvivalAnalysisDefaultBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get survival analysis default body
func (o *GetSurvivalAnalysisDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSurvivalAnalysisDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSurvivalAnalysisDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetSurvivalAnalysisDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetSurvivalAnalysisOKBody get survival analysis o k body
// swagger:model GetSurvivalAnalysisOKBody
type GetSurvivalAnalysisOKBody struct {

	// results
	Results []*ResultsItems0 `json:"results"`

	// timers
	Timers map[string]float64 `json:"timers,omitempty"`
}

// Validate validates this get survival analysis o k body
func (o *GetSurvivalAnalysisOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSurvivalAnalysisOKBody) validateResults(formats strfmt.Registry) error {

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
					return ve.ValidateName("getSurvivalAnalysisOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSurvivalAnalysisOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSurvivalAnalysisOKBody) UnmarshalBinary(b []byte) error {
	var res GetSurvivalAnalysisOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ResultsItems0 results items0
// swagger:model ResultsItems0
type ResultsItems0 struct {

	// events
	Events *ResultsItems0Events `json:"events,omitempty"`

	// timepoint
	Timepoint string `json:"timepoint,omitempty"`
}

// Validate validates this results items0
func (o *ResultsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ResultsItems0) validateEvents(formats strfmt.Registry) error {

	if swag.IsZero(o.Events) { // not required
		return nil
	}

	if o.Events != nil {
		if err := o.Events.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("events")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ResultsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ResultsItems0) UnmarshalBinary(b []byte) error {
	var res ResultsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ResultsItems0Events results items0 events
// swagger:model ResultsItems0Events
type ResultsItems0Events struct {

	// censoringevent
	Censoringevent string `json:"censoringevent,omitempty"`

	// eventofinterest
	Eventofinterest string `json:"eventofinterest,omitempty"`
}

// Validate validates this results items0 events
func (o *ResultsItems0Events) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ResultsItems0Events) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ResultsItems0Events) UnmarshalBinary(b []byte) error {
	var res ResultsItems0Events
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
