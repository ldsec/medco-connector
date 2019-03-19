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

// SearchResultElement search result element
// swagger:model searchResultElement
type SearchResultElement struct {

	// code
	Code string `json:"code,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// medco encryption
	MedcoEncryption *SearchResultElementMedcoEncryption `json:"medcoEncryption,omitempty"`

	// metadata
	Metadata interface{} `json:"metadata,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// value type
	// Enum: [none numeric enum text genomic_annotation]
	ValueType string `json:"valueType,omitempty"`
}

// Validate validates this search result element
func (m *SearchResultElement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMedcoEncryption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchResultElement) validateMedcoEncryption(formats strfmt.Registry) error {

	if swag.IsZero(m.MedcoEncryption) { // not required
		return nil
	}

	if m.MedcoEncryption != nil {
		if err := m.MedcoEncryption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("medcoEncryption")
			}
			return err
		}
	}

	return nil
}

var searchResultElementTypeValueTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","numeric","enum","text","genomic_annotation"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchResultElementTypeValueTypePropEnum = append(searchResultElementTypeValueTypePropEnum, v)
	}
}

const (

	// SearchResultElementValueTypeNone captures enum value "none"
	SearchResultElementValueTypeNone string = "none"

	// SearchResultElementValueTypeNumeric captures enum value "numeric"
	SearchResultElementValueTypeNumeric string = "numeric"

	// SearchResultElementValueTypeEnum captures enum value "enum"
	SearchResultElementValueTypeEnum string = "enum"

	// SearchResultElementValueTypeText captures enum value "text"
	SearchResultElementValueTypeText string = "text"

	// SearchResultElementValueTypeGenomicAnnotation captures enum value "genomic_annotation"
	SearchResultElementValueTypeGenomicAnnotation string = "genomic_annotation"
)

// prop value enum
func (m *SearchResultElement) validateValueTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, searchResultElementTypeValueTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SearchResultElement) validateValueType(formats strfmt.Registry) error {

	if swag.IsZero(m.ValueType) { // not required
		return nil
	}

	// value enum
	if err := m.validateValueTypeEnum("valueType", "body", m.ValueType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchResultElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchResultElement) UnmarshalBinary(b []byte) error {
	var res SearchResultElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SearchResultElementMedcoEncryption search result element medco encryption
// swagger:model SearchResultElementMedcoEncryption
type SearchResultElementMedcoEncryption struct {

	// children ids
	ChildrenIds []int64 `json:"childrenIds"`

	// encrypted
	Encrypted bool `json:"encrypted,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// type
	// Enum: [CONCEPT_PARENT_NODE CONCEPT_INTERNAL_NODE CONCEPT_LEAF]
	Type string `json:"type,omitempty"`
}

// Validate validates this search result element medco encryption
func (m *SearchResultElementMedcoEncryption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var searchResultElementMedcoEncryptionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CONCEPT_PARENT_NODE","CONCEPT_INTERNAL_NODE","CONCEPT_LEAF"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchResultElementMedcoEncryptionTypeTypePropEnum = append(searchResultElementMedcoEncryptionTypeTypePropEnum, v)
	}
}

const (

	// SearchResultElementMedcoEncryptionTypeCONCEPTPARENTNODE captures enum value "CONCEPT_PARENT_NODE"
	SearchResultElementMedcoEncryptionTypeCONCEPTPARENTNODE string = "CONCEPT_PARENT_NODE"

	// SearchResultElementMedcoEncryptionTypeCONCEPTINTERNALNODE captures enum value "CONCEPT_INTERNAL_NODE"
	SearchResultElementMedcoEncryptionTypeCONCEPTINTERNALNODE string = "CONCEPT_INTERNAL_NODE"

	// SearchResultElementMedcoEncryptionTypeCONCEPTLEAF captures enum value "CONCEPT_LEAF"
	SearchResultElementMedcoEncryptionTypeCONCEPTLEAF string = "CONCEPT_LEAF"
)

// prop value enum
func (m *SearchResultElementMedcoEncryption) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, searchResultElementMedcoEncryptionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SearchResultElementMedcoEncryption) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("medcoEncryption"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchResultElementMedcoEncryption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchResultElementMedcoEncryption) UnmarshalBinary(b []byte) error {
	var res SearchResultElementMedcoEncryption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
