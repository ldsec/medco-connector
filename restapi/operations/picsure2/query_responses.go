// Code generated by go-swagger; DO NOT EDIT.

package picsure2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/lca1/medco-connector/models"
)

// QueryOKCode is the HTTP code returned for type QueryOK
const QueryOKCode int = 200

/*QueryOK Query status.

swagger:response queryOK
*/
type QueryOK struct {

	/*
	  In: Body
	*/
	Payload *models.QueryStatus `json:"body,omitempty"`
}

// NewQueryOK creates QueryOK with default headers values
func NewQueryOK() *QueryOK {

	return &QueryOK{}
}

// WithPayload adds the payload to the query o k response
func (o *QueryOK) WithPayload(payload *models.QueryStatus) *QueryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the query o k response
func (o *QueryOK) SetPayload(payload *models.QueryStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QueryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
