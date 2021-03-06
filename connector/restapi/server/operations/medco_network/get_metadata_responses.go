// Code generated by go-swagger; DO NOT EDIT.

package medco_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetMetadataOKCode is the HTTP code returned for type GetMetadataOK
const GetMetadataOKCode int = 200

/*GetMetadataOK Network metadata (public key and nodes list).

swagger:response getMetadataOK
*/
type GetMetadataOK struct {

	/*
	  In: Body
	*/
	Payload *GetMetadataOKBody `json:"body,omitempty"`
}

// NewGetMetadataOK creates GetMetadataOK with default headers values
func NewGetMetadataOK() *GetMetadataOK {

	return &GetMetadataOK{}
}

// WithPayload adds the payload to the get metadata o k response
func (o *GetMetadataOK) WithPayload(payload *GetMetadataOKBody) *GetMetadataOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get metadata o k response
func (o *GetMetadataOK) SetPayload(payload *GetMetadataOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMetadataOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetMetadataDefault Error response.

swagger:response getMetadataDefault
*/
type GetMetadataDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *GetMetadataDefaultBody `json:"body,omitempty"`
}

// NewGetMetadataDefault creates GetMetadataDefault with default headers values
func NewGetMetadataDefault(code int) *GetMetadataDefault {
	if code <= 0 {
		code = 500
	}

	return &GetMetadataDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get metadata default response
func (o *GetMetadataDefault) WithStatusCode(code int) *GetMetadataDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get metadata default response
func (o *GetMetadataDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get metadata default response
func (o *GetMetadataDefault) WithPayload(payload *GetMetadataDefaultBody) *GetMetadataDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get metadata default response
func (o *GetMetadataDefault) SetPayload(payload *GetMetadataDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMetadataDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
