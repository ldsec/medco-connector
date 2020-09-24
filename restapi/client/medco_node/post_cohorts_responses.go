// Code generated by go-swagger; DO NOT EDIT.

package medco_node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostCohortsReader is a Reader for the PostCohorts structure.
type PostCohortsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCohortsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCohortsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPostCohortsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostCohortsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostCohortsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostCohortsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCohortsOK creates a PostCohortsOK with default headers values
func NewPostCohortsOK() *PostCohortsOK {
	return &PostCohortsOK{}
}

/*PostCohortsOK handles this case with default header values.

Updated cohort
*/
type PostCohortsOK struct {
	Payload string
}

func (o *PostCohortsOK) Error() string {
	return fmt.Sprintf("[POST /node/explore/cohorts][%d] postCohortsOK  %+v", 200, o.Payload)
}

func (o *PostCohortsOK) GetPayload() string {
	return o.Payload
}

func (o *PostCohortsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCohortsForbidden creates a PostCohortsForbidden with default headers values
func NewPostCohortsForbidden() *PostCohortsForbidden {
	return &PostCohortsForbidden{}
}

/*PostCohortsForbidden handles this case with default header values.

Not authorized
*/
type PostCohortsForbidden struct {
}

func (o *PostCohortsForbidden) Error() string {
	return fmt.Sprintf("[POST /node/explore/cohorts][%d] postCohortsForbidden ", 403)
}

func (o *PostCohortsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCohortsNotFound creates a PostCohortsNotFound with default headers values
func NewPostCohortsNotFound() *PostCohortsNotFound {
	return &PostCohortsNotFound{}
}

/*PostCohortsNotFound handles this case with default header values.

User not found
*/
type PostCohortsNotFound struct {
}

func (o *PostCohortsNotFound) Error() string {
	return fmt.Sprintf("[POST /node/explore/cohorts][%d] postCohortsNotFound ", 404)
}

func (o *PostCohortsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCohortsInternalServerError creates a PostCohortsInternalServerError with default headers values
func NewPostCohortsInternalServerError() *PostCohortsInternalServerError {
	return &PostCohortsInternalServerError{}
}

/*PostCohortsInternalServerError handles this case with default header values.

DB has been updated since last importation
*/
type PostCohortsInternalServerError struct {
}

func (o *PostCohortsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /node/explore/cohorts][%d] postCohortsInternalServerError ", 500)
}

func (o *PostCohortsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCohortsDefault creates a PostCohortsDefault with default headers values
func NewPostCohortsDefault(code int) *PostCohortsDefault {
	return &PostCohortsDefault{
		_statusCode: code,
	}
}

/*PostCohortsDefault handles this case with default header values.

Error response.
*/
type PostCohortsDefault struct {
	_statusCode int

	Payload *PostCohortsDefaultBody
}

// Code gets the status code for the post cohorts default response
func (o *PostCohortsDefault) Code() int {
	return o._statusCode
}

func (o *PostCohortsDefault) Error() string {
	return fmt.Sprintf("[POST /node/explore/cohorts][%d] postCohorts default  %+v", o._statusCode, o.Payload)
}

func (o *PostCohortsDefault) GetPayload() *PostCohortsDefaultBody {
	return o.Payload
}

func (o *PostCohortsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostCohortsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostCohortsBody post cohorts body
swagger:model PostCohortsBody
*/
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

/*PostCohortsDefaultBody post cohorts default body
swagger:model PostCohortsDefaultBody
*/
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
