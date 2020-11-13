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

// ExploreSearchModifierChildrenReader is a Reader for the ExploreSearchModifierChildren structure.
type ExploreSearchModifierChildrenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExploreSearchModifierChildrenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExploreSearchModifierChildrenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExploreSearchModifierChildrenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExploreSearchModifierChildrenOK creates a ExploreSearchModifierChildrenOK with default headers values
func NewExploreSearchModifierChildrenOK() *ExploreSearchModifierChildrenOK {
	return &ExploreSearchModifierChildrenOK{}
}

/*ExploreSearchModifierChildrenOK handles this case with default header values.

MedCo-Explore search modifier children query response.
*/
type ExploreSearchModifierChildrenOK struct {
	Payload *ExploreSearchModifierChildrenOKBody
}

func (o *ExploreSearchModifierChildrenOK) Error() string {
	return fmt.Sprintf("[POST /node/explore/search/modifier-children][%d] exploreSearchModifierChildrenOK  %+v", 200, o.Payload)
}

func (o *ExploreSearchModifierChildrenOK) GetPayload() *ExploreSearchModifierChildrenOKBody {
	return o.Payload
}

func (o *ExploreSearchModifierChildrenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExploreSearchModifierChildrenOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExploreSearchModifierChildrenDefault creates a ExploreSearchModifierChildrenDefault with default headers values
func NewExploreSearchModifierChildrenDefault(code int) *ExploreSearchModifierChildrenDefault {
	return &ExploreSearchModifierChildrenDefault{
		_statusCode: code,
	}
}

/*ExploreSearchModifierChildrenDefault handles this case with default header values.

Error response.
*/
type ExploreSearchModifierChildrenDefault struct {
	_statusCode int

	Payload *ExploreSearchModifierChildrenDefaultBody
}

// Code gets the status code for the explore search modifier children default response
func (o *ExploreSearchModifierChildrenDefault) Code() int {
	return o._statusCode
}

func (o *ExploreSearchModifierChildrenDefault) Error() string {
	return fmt.Sprintf("[POST /node/explore/search/modifier-children][%d] exploreSearchModifierChildren default  %+v", o._statusCode, o.Payload)
}

func (o *ExploreSearchModifierChildrenDefault) GetPayload() *ExploreSearchModifierChildrenDefaultBody {
	return o.Payload
}

func (o *ExploreSearchModifierChildrenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExploreSearchModifierChildrenDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ExploreSearchModifierChildrenDefaultBody explore search modifier children default body
swagger:model ExploreSearchModifierChildrenDefaultBody
*/
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

/*ExploreSearchModifierChildrenOKBody explore search modifier children o k body
swagger:model ExploreSearchModifierChildrenOKBody
*/
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