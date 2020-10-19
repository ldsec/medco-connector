// Code generated by go-swagger; DO NOT EDIT.

package survival_analysis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/ldsec/medco/connector/restapi/models"
)

// SurvivalAnalysisReader is a Reader for the SurvivalAnalysis structure.
type SurvivalAnalysisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SurvivalAnalysisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSurvivalAnalysisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSurvivalAnalysisDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSurvivalAnalysisOK creates a SurvivalAnalysisOK with default headers values
func NewSurvivalAnalysisOK() *SurvivalAnalysisOK {
	return &SurvivalAnalysisOK{}
}

/*SurvivalAnalysisOK handles this case with default header values.

Queried survival analysis
*/
type SurvivalAnalysisOK struct {
	Payload *SurvivalAnalysisOKBody
}

func (o *SurvivalAnalysisOK) Error() string {
	return fmt.Sprintf("[POST /node/analysis/survival/query][%d] survivalAnalysisOK  %+v", 200, o.Payload)
}

func (o *SurvivalAnalysisOK) GetPayload() *SurvivalAnalysisOKBody {
	return o.Payload
}

func (o *SurvivalAnalysisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SurvivalAnalysisOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSurvivalAnalysisDefault creates a SurvivalAnalysisDefault with default headers values
func NewSurvivalAnalysisDefault(code int) *SurvivalAnalysisDefault {
	return &SurvivalAnalysisDefault{
		_statusCode: code,
	}
}

/*SurvivalAnalysisDefault handles this case with default header values.

Error response.
*/
type SurvivalAnalysisDefault struct {
	_statusCode int

	Payload *SurvivalAnalysisDefaultBody
}

// Code gets the status code for the survival analysis default response
func (o *SurvivalAnalysisDefault) Code() int {
	return o._statusCode
}

func (o *SurvivalAnalysisDefault) Error() string {
	return fmt.Sprintf("[POST /node/analysis/survival/query][%d] survivalAnalysis default  %+v", o._statusCode, o.Payload)
}

func (o *SurvivalAnalysisDefault) GetPayload() *SurvivalAnalysisDefaultBody {
	return o.Payload
}

func (o *SurvivalAnalysisDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SurvivalAnalysisDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SurvivalAnalysisBody survival analysis body
swagger:model SurvivalAnalysisBody
*/
type SurvivalAnalysisBody struct {

	// ID
	ID string `json:"ID,omitempty"`

	// cohort name
	// Pattern: ^\w+$
	CohortName string `json:"cohortName,omitempty"`

	// end concept
	EndConcept string `json:"endConcept,omitempty"`

	// end modifier
	EndModifier string `json:"endModifier,omitempty"`

	// start concept
	StartConcept string `json:"startConcept,omitempty"`

	// start modifier
	StartModifier string `json:"startModifier,omitempty"`

	// sub group definitions
	// Max Items: 4
	SubGroupDefinitions []*SurvivalAnalysisParamsBodySubGroupDefinitionsItems0 `json:"subGroupDefinitions"`

	// time granularity
	// Enum: [day week month year]
	TimeGranularity string `json:"timeGranularity,omitempty"`

	// time limit
	TimeLimit int64 `json:"timeLimit,omitempty"`

	// user public key
	// Pattern: ^[\w=-]+$
	UserPublicKey string `json:"userPublicKey,omitempty"`
}

// Validate validates this survival analysis body
func (o *SurvivalAnalysisBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCohortName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubGroupDefinitions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTimeGranularity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUserPublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SurvivalAnalysisBody) validateCohortName(formats strfmt.Registry) error {

	if swag.IsZero(o.CohortName) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"cohortName", "body", string(o.CohortName), `^\w+$`); err != nil {
		return err
	}

	return nil
}

func (o *SurvivalAnalysisBody) validateSubGroupDefinitions(formats strfmt.Registry) error {

	if swag.IsZero(o.SubGroupDefinitions) { // not required
		return nil
	}

	iSubGroupDefinitionsSize := int64(len(o.SubGroupDefinitions))

	if err := validate.MaxItems("body"+"."+"subGroupDefinitions", "body", iSubGroupDefinitionsSize, 4); err != nil {
		return err
	}

	for i := 0; i < len(o.SubGroupDefinitions); i++ {
		if swag.IsZero(o.SubGroupDefinitions[i]) { // not required
			continue
		}

		if o.SubGroupDefinitions[i] != nil {
			if err := o.SubGroupDefinitions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "subGroupDefinitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var survivalAnalysisBodyTypeTimeGranularityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["day","week","month","year"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		survivalAnalysisBodyTypeTimeGranularityPropEnum = append(survivalAnalysisBodyTypeTimeGranularityPropEnum, v)
	}
}

const (

	// SurvivalAnalysisBodyTimeGranularityDay captures enum value "day"
	SurvivalAnalysisBodyTimeGranularityDay string = "day"

	// SurvivalAnalysisBodyTimeGranularityWeek captures enum value "week"
	SurvivalAnalysisBodyTimeGranularityWeek string = "week"

	// SurvivalAnalysisBodyTimeGranularityMonth captures enum value "month"
	SurvivalAnalysisBodyTimeGranularityMonth string = "month"

	// SurvivalAnalysisBodyTimeGranularityYear captures enum value "year"
	SurvivalAnalysisBodyTimeGranularityYear string = "year"
)

// prop value enum
func (o *SurvivalAnalysisBody) validateTimeGranularityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, survivalAnalysisBodyTypeTimeGranularityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *SurvivalAnalysisBody) validateTimeGranularity(formats strfmt.Registry) error {

	if swag.IsZero(o.TimeGranularity) { // not required
		return nil
	}

	// value enum
	if err := o.validateTimeGranularityEnum("body"+"."+"timeGranularity", "body", o.TimeGranularity); err != nil {
		return err
	}

	return nil
}

func (o *SurvivalAnalysisBody) validateUserPublicKey(formats strfmt.Registry) error {

	if swag.IsZero(o.UserPublicKey) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"userPublicKey", "body", string(o.UserPublicKey), `^[\w=-]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SurvivalAnalysisBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SurvivalAnalysisBody) UnmarshalBinary(b []byte) error {
	var res SurvivalAnalysisBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SurvivalAnalysisDefaultBody survival analysis default body
swagger:model SurvivalAnalysisDefaultBody
*/
type SurvivalAnalysisDefaultBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this survival analysis default body
func (o *SurvivalAnalysisDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SurvivalAnalysisDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SurvivalAnalysisDefaultBody) UnmarshalBinary(b []byte) error {
	var res SurvivalAnalysisDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SurvivalAnalysisOKBody survival analysis o k body
swagger:model SurvivalAnalysisOKBody
*/
type SurvivalAnalysisOKBody struct {

	// results
	Results []*SurvivalAnalysisOKBodyResultsItems0 `json:"results"`

	// timers
	Timers models.Timers `json:"timers"`
}

// Validate validates this survival analysis o k body
func (o *SurvivalAnalysisOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTimers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SurvivalAnalysisOKBody) validateResults(formats strfmt.Registry) error {

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
					return ve.ValidateName("survivalAnalysisOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SurvivalAnalysisOKBody) validateTimers(formats strfmt.Registry) error {

	if swag.IsZero(o.Timers) { // not required
		return nil
	}

	if err := o.Timers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("survivalAnalysisOK" + "." + "timers")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SurvivalAnalysisOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SurvivalAnalysisOKBody) UnmarshalBinary(b []byte) error {
	var res SurvivalAnalysisOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SurvivalAnalysisOKBodyResultsItems0 survival analysis o k body results items0
swagger:model SurvivalAnalysisOKBodyResultsItems0
*/
type SurvivalAnalysisOKBodyResultsItems0 struct {

	// group ID
	GroupID string `json:"groupID,omitempty"`

	// group results
	GroupResults []*SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0 `json:"groupResults"`

	// initial count
	InitialCount string `json:"initialCount,omitempty"`
}

// Validate validates this survival analysis o k body results items0
func (o *SurvivalAnalysisOKBodyResultsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGroupResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SurvivalAnalysisOKBodyResultsItems0) validateGroupResults(formats strfmt.Registry) error {

	if swag.IsZero(o.GroupResults) { // not required
		return nil
	}

	for i := 0; i < len(o.GroupResults); i++ {
		if swag.IsZero(o.GroupResults[i]) { // not required
			continue
		}

		if o.GroupResults[i] != nil {
			if err := o.GroupResults[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groupResults" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SurvivalAnalysisOKBodyResultsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SurvivalAnalysisOKBodyResultsItems0) UnmarshalBinary(b []byte) error {
	var res SurvivalAnalysisOKBodyResultsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0 survival analysis o k body results items0 group results items0
swagger:model SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0
*/
type SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0 struct {

	// events
	Events *SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0Events `json:"events,omitempty"`

	// timepoint
	Timepoint int64 `json:"timepoint,omitempty"`
}

// Validate validates this survival analysis o k body results items0 group results items0
func (o *SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0) validateEvents(formats strfmt.Registry) error {

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
func (o *SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0) UnmarshalBinary(b []byte) error {
	var res SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0Events survival analysis o k body results items0 group results items0 events
swagger:model SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0Events
*/
type SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0Events struct {

	// censoringevent
	Censoringevent string `json:"censoringevent,omitempty"`

	// eventofinterest
	Eventofinterest string `json:"eventofinterest,omitempty"`
}

// Validate validates this survival analysis o k body results items0 group results items0 events
func (o *SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0Events) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0Events) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0Events) UnmarshalBinary(b []byte) error {
	var res SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0Events
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SurvivalAnalysisParamsBodySubGroupDefinitionsItems0 survival analysis params body sub group definitions items0
swagger:model SurvivalAnalysisParamsBodySubGroupDefinitionsItems0
*/
type SurvivalAnalysisParamsBodySubGroupDefinitionsItems0 struct {

	// cohort name
	CohortName string `json:"cohortName,omitempty"`

	// panels
	Panels []*models.Panel `json:"panels"`
}

// Validate validates this survival analysis params body sub group definitions items0
func (o *SurvivalAnalysisParamsBodySubGroupDefinitionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePanels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SurvivalAnalysisParamsBodySubGroupDefinitionsItems0) validatePanels(formats strfmt.Registry) error {

	if swag.IsZero(o.Panels) { // not required
		return nil
	}

	for i := 0; i < len(o.Panels); i++ {
		if swag.IsZero(o.Panels[i]) { // not required
			continue
		}

		if o.Panels[i] != nil {
			if err := o.Panels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("panels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SurvivalAnalysisParamsBodySubGroupDefinitionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SurvivalAnalysisParamsBodySubGroupDefinitionsItems0) UnmarshalBinary(b []byte) error {
	var res SurvivalAnalysisParamsBodySubGroupDefinitionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}