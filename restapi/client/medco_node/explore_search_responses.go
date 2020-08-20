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

	"github.com/ldsec/medco-connector/restapi/models"
)

// ExploreSearchReader is a Reader for the ExploreSearch structure.
type ExploreSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExploreSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExploreSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExploreSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExploreSearchOK creates a ExploreSearchOK with default headers values
func NewExploreSearchOK() *ExploreSearchOK {
	return &ExploreSearchOK{}
}

/*ExploreSearchOK handles this case with default header values.

MedCo-Explore search query response.
*/
type ExploreSearchOK struct {
	Payload *ExploreSearchOKBody
}

func (o *ExploreSearchOK) Error() string {
	return fmt.Sprintf("[POST /node/explore/search][%d] exploreSearchOK  %+v", 200, o.Payload)
}

func (o *ExploreSearchOK) GetPayload() *ExploreSearchOKBody {
	return o.Payload
}

func (o *ExploreSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExploreSearchOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExploreSearchDefault creates a ExploreSearchDefault with default headers values
func NewExploreSearchDefault(code int) *ExploreSearchDefault {
	return &ExploreSearchDefault{
		_statusCode: code,
	}
}

/*ExploreSearchDefault handles this case with default header values.

Error response.
*/
type ExploreSearchDefault struct {
	_statusCode int

	Payload *ExploreSearchDefaultBody
}

// Code gets the status code for the explore search default response
func (o *ExploreSearchDefault) Code() int {
	return o._statusCode
}

func (o *ExploreSearchDefault) Error() string {
	return fmt.Sprintf("[POST /node/explore/search][%d] exploreSearch default  %+v", o._statusCode, o.Payload)
}

func (o *ExploreSearchDefault) GetPayload() *ExploreSearchDefaultBody {
	return o.Payload
}

func (o *ExploreSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExploreSearchDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ExploreSearchDefaultBody explore search default body
swagger:model ExploreSearchDefaultBody
*/
type ExploreSearchDefaultBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this explore search default body
func (o *ExploreSearchDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExploreSearchDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExploreSearchDefaultBody) UnmarshalBinary(b []byte) error {
	var res ExploreSearchDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ExploreSearchOKBody explore search o k body
swagger:model ExploreSearchOKBody
*/
type ExploreSearchOKBody struct {

	// results
	Results []*models.ExploreSearchResultElement `json:"results"`

	// search
	Search *models.ExploreSearch `json:"search,omitempty"`
}

// Validate validates this explore search o k body
func (o *ExploreSearchOKBody) Validate(formats strfmt.Registry) error {
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

func (o *ExploreSearchOKBody) validateResults(formats strfmt.Registry) error {

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
					return ve.ValidateName("exploreSearchOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ExploreSearchOKBody) validateSearch(formats strfmt.Registry) error {

	if swag.IsZero(o.Search) { // not required
		return nil
	}

	if o.Search != nil {
		if err := o.Search.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exploreSearchOK" + "." + "search")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExploreSearchOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExploreSearchOKBody) UnmarshalBinary(b []byte) error {
	var res ExploreSearchOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
