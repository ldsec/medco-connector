// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// QueryStatus query status
// swagger:model queryStatus
type QueryStatus struct {

	// duration
	Duration int64 `json:"duration,omitempty"`

	// expiration
	Expiration int64 `json:"expiration,omitempty"`

	// picsure result Id
	PicsureResultID string `json:"picsureResultId,omitempty"`

	// resource ID
	ResourceID string `json:"resourceID,omitempty"`

	// resource result Id
	ResourceResultID string `json:"resourceResultId,omitempty"`

	// resource status
	ResourceStatus string `json:"resourceStatus,omitempty"`

	// result metadata
	// Format: byte
	ResultMetadata strfmt.Base64 `json:"resultMetadata,omitempty"`

	// size in bytes
	SizeInBytes int64 `json:"sizeInBytes,omitempty"`

	// start time
	StartTime int64 `json:"startTime,omitempty"`

	// status
	// Enum: [QUEUED PENDING ERROR AVAILABLE]
	Status string `json:"status,omitempty"`
}

// Validate validates this query status
func (m *QueryStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResultMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryStatus) validateResultMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.ResultMetadata) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

var queryStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["QUEUED","PENDING","ERROR","AVAILABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		queryStatusTypeStatusPropEnum = append(queryStatusTypeStatusPropEnum, v)
	}
}

const (

	// QueryStatusStatusQUEUED captures enum value "QUEUED"
	QueryStatusStatusQUEUED string = "QUEUED"

	// QueryStatusStatusPENDING captures enum value "PENDING"
	QueryStatusStatusPENDING string = "PENDING"

	// QueryStatusStatusERROR captures enum value "ERROR"
	QueryStatusStatusERROR string = "ERROR"

	// QueryStatusStatusAVAILABLE captures enum value "AVAILABLE"
	QueryStatusStatusAVAILABLE string = "AVAILABLE"
)

// prop value enum
func (m *QueryStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, queryStatusTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *QueryStatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryStatus) UnmarshalBinary(b []byte) error {
	var res QueryStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
