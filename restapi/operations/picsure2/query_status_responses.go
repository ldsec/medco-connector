// Code generated by go-swagger; DO NOT EDIT.

package picsure2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/lca1/medco-connector/models"
)

// QueryStatusOKCode is the HTTP code returned for type QueryStatusOK
const QueryStatusOKCode int = 200

/*QueryStatusOK Query status.

swagger:response queryStatusOK
*/
type QueryStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.QueryStatus `json:"body,omitempty"`
}

// NewQueryStatusOK creates QueryStatusOK with default headers values
func NewQueryStatusOK() *QueryStatusOK {

	return &QueryStatusOK{}
}

// WithPayload adds the payload to the query status o k response
func (o *QueryStatusOK) WithPayload(payload *models.QueryStatus) *QueryStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the query status o k response
func (o *QueryStatusOK) SetPayload(payload *models.QueryStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QueryStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
