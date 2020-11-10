// Code generated by go-swagger; DO NOT EDIT.

package medco_node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ldsec/medco/connector/restapi/models"
)

// ExploreSearchConceptInfoReader is a Reader for the ExploreSearchConceptInfo structure.
type ExploreSearchConceptInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExploreSearchConceptInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExploreSearchConceptInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExploreSearchConceptInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExploreSearchConceptInfoOK creates a ExploreSearchConceptInfoOK with default headers values
func NewExploreSearchConceptInfoOK() *ExploreSearchConceptInfoOK {
	return &ExploreSearchConceptInfoOK{}
}

/*ExploreSearchConceptInfoOK handles this case with default header values.

MedCo-Explore search concept info query response.
*/
type ExploreSearchConceptInfoOK struct {
	Payload *ExploreSearchConceptInfoOKBody
}

func (o *ExploreSearchConceptInfoOK) Error() string {
	return fmt.Sprintf("[POST /node/explore/search/concept-info][%d] exploreSearchConceptInfoOK  %+v", 200, o.Payload)
}

func (o *ExploreSearchConceptInfoOK) GetPayload() *ExploreSearchConceptInfoOKBody {
	return o.Payload
}

func (o *ExploreSearchConceptInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExploreSearchConceptInfoOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExploreSearchConceptInfoDefault creates a ExploreSearchConceptInfoDefault with default headers values
func NewExploreSearchConceptInfoDefault(code int) *ExploreSearchConceptInfoDefault {
	return &ExploreSearchConceptInfoDefault{
		_statusCode: code,
	}
}

/*ExploreSearchConceptInfoDefault handles this case with default header values.

Error response.
*/
type ExploreSearchConceptInfoDefault struct {
	_statusCode int

	Payload *ExploreSearchConceptInfoDefaultBody
}

// Code gets the status code for the explore search concept info default response
func (o *ExploreSearchConceptInfoDefault) Code() int {
	return o._statusCode
}

func (o *ExploreSearchConceptInfoDefault) Error() string {
	return fmt.Sprintf("[POST /node/explore/search/concept-info][%d] exploreSearchConceptInfo default  %+v", o._statusCode, o.Payload)
}

func (o *ExploreSearchConceptInfoDefault) GetPayload() *ExploreSearchConceptInfoDefaultBody {
	return o.Payload
}

func (o *ExploreSearchConceptInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExploreSearchConceptInfoDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ExploreSearchConceptInfoDefaultBody explore search concept info default body
swagger:model ExploreSearchConceptInfoDefaultBody
*/
type ExploreSearchConceptInfoDefaultBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this explore search concept info default body
func (o *ExploreSearchConceptInfoDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExploreSearchConceptInfoDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExploreSearchConceptInfoDefaultBody) UnmarshalBinary(b []byte) error {
	var res ExploreSearchConceptInfoDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ExploreSearchConceptInfoOKBody explore search concept info o k body
swagger:model ExploreSearchConceptInfoOKBody
*/
type ExploreSearchConceptInfoOKBody struct {

	// results
	Results []*models.ExploreSearchResultElement `json:"results"`

	// search
	Search *models.ExploreSearchConceptInfo `json:"search,omitempty"`
}

// Validate validates this explore search concept info o k body
func (o *ExploreSearchConceptInfoOKBody) Validate(formats strfmt.Registry) error {
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

func (o *ExploreSearchConceptInfoOKBody) validateResults(formats strfmt.Registry) error {

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
					return ve.ValidateName("exploreSearchConceptInfoOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ExploreSearchConceptInfoOKBody) validateSearch(formats strfmt.Registry) error {

	if swag.IsZero(o.Search) { // not required
		return nil
	}

	if o.Search != nil {
		if err := o.Search.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exploreSearchConceptInfoOK" + "." + "search")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExploreSearchConceptInfoOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExploreSearchConceptInfoOKBody) UnmarshalBinary(b []byte) error {
	var res ExploreSearchConceptInfoOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
