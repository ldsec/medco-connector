// Code generated by go-swagger; DO NOT EDIT.

package medco_node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ExploreSearchModifierChildrenOKCode is the HTTP code returned for type ExploreSearchModifierChildrenOK
const ExploreSearchModifierChildrenOKCode int = 200

/*ExploreSearchModifierChildrenOK MedCo-Explore search modifier children query response.

swagger:response exploreSearchModifierChildrenOK
*/
type ExploreSearchModifierChildrenOK struct {

	/*
	  In: Body
	*/
	Payload *ExploreSearchModifierChildrenOKBody `json:"body,omitempty"`
}

// NewExploreSearchModifierChildrenOK creates ExploreSearchModifierChildrenOK with default headers values
func NewExploreSearchModifierChildrenOK() *ExploreSearchModifierChildrenOK {

	return &ExploreSearchModifierChildrenOK{}
}

// WithPayload adds the payload to the explore search modifier children o k response
func (o *ExploreSearchModifierChildrenOK) WithPayload(payload *ExploreSearchModifierChildrenOKBody) *ExploreSearchModifierChildrenOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the explore search modifier children o k response
func (o *ExploreSearchModifierChildrenOK) SetPayload(payload *ExploreSearchModifierChildrenOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExploreSearchModifierChildrenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ExploreSearchModifierChildrenDefault Error response.

swagger:response exploreSearchModifierChildrenDefault
*/
type ExploreSearchModifierChildrenDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *ExploreSearchModifierChildrenDefaultBody `json:"body,omitempty"`
}

// NewExploreSearchModifierChildrenDefault creates ExploreSearchModifierChildrenDefault with default headers values
func NewExploreSearchModifierChildrenDefault(code int) *ExploreSearchModifierChildrenDefault {
	if code <= 0 {
		code = 500
	}

	return &ExploreSearchModifierChildrenDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the explore search modifier children default response
func (o *ExploreSearchModifierChildrenDefault) WithStatusCode(code int) *ExploreSearchModifierChildrenDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the explore search modifier children default response
func (o *ExploreSearchModifierChildrenDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the explore search modifier children default response
func (o *ExploreSearchModifierChildrenDefault) WithPayload(payload *ExploreSearchModifierChildrenDefaultBody) *ExploreSearchModifierChildrenDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the explore search modifier children default response
func (o *ExploreSearchModifierChildrenDefault) SetPayload(payload *ExploreSearchModifierChildrenDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExploreSearchModifierChildrenDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}