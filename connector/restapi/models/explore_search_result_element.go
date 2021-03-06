// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExploreSearchResultElement explore search result element
//
// swagger:model exploreSearchResultElement
type ExploreSearchResultElement struct {

	// applied path
	AppliedPath string `json:"appliedPath,omitempty"`

	// code
	Code string `json:"code,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// leaf
	// Required: true
	Leaf *bool `json:"leaf"`

	// medco encryption
	MedcoEncryption *ExploreSearchResultElementMedcoEncryption `json:"medcoEncryption,omitempty"`

	// metadata
	Metadata *Metadataxml `json:"metadata,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// type
	// Enum: [concept concept_container concept_folder modifier modifier_container modifier_folder genomic_annotation]
	Type string `json:"type,omitempty"`
}

// Validate validates this explore search result element
func (m *ExploreSearchResultElement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLeaf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMedcoEncryption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExploreSearchResultElement) validateLeaf(formats strfmt.Registry) error {

	if err := validate.Required("leaf", "body", m.Leaf); err != nil {
		return err
	}

	return nil
}

func (m *ExploreSearchResultElement) validateMedcoEncryption(formats strfmt.Registry) error {

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

func (m *ExploreSearchResultElement) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

var exploreSearchResultElementTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["concept","concept_container","concept_folder","modifier","modifier_container","modifier_folder","genomic_annotation"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		exploreSearchResultElementTypeTypePropEnum = append(exploreSearchResultElementTypeTypePropEnum, v)
	}
}

const (

	// ExploreSearchResultElementTypeConcept captures enum value "concept"
	ExploreSearchResultElementTypeConcept string = "concept"

	// ExploreSearchResultElementTypeConceptContainer captures enum value "concept_container"
	ExploreSearchResultElementTypeConceptContainer string = "concept_container"

	// ExploreSearchResultElementTypeConceptFolder captures enum value "concept_folder"
	ExploreSearchResultElementTypeConceptFolder string = "concept_folder"

	// ExploreSearchResultElementTypeModifier captures enum value "modifier"
	ExploreSearchResultElementTypeModifier string = "modifier"

	// ExploreSearchResultElementTypeModifierContainer captures enum value "modifier_container"
	ExploreSearchResultElementTypeModifierContainer string = "modifier_container"

	// ExploreSearchResultElementTypeModifierFolder captures enum value "modifier_folder"
	ExploreSearchResultElementTypeModifierFolder string = "modifier_folder"

	// ExploreSearchResultElementTypeGenomicAnnotation captures enum value "genomic_annotation"
	ExploreSearchResultElementTypeGenomicAnnotation string = "genomic_annotation"
)

// prop value enum
func (m *ExploreSearchResultElement) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, exploreSearchResultElementTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExploreSearchResultElement) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExploreSearchResultElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExploreSearchResultElement) UnmarshalBinary(b []byte) error {
	var res ExploreSearchResultElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ExploreSearchResultElementMedcoEncryption explore search result element medco encryption
//
// swagger:model ExploreSearchResultElementMedcoEncryption
type ExploreSearchResultElementMedcoEncryption struct {

	// children ids
	ChildrenIds []int64 `json:"childrenIds"`

	// encrypted
	// Required: true
	Encrypted *bool `json:"encrypted"`

	// id
	// Required: true
	ID *int64 `json:"id"`
}

// Validate validates this explore search result element medco encryption
func (m *ExploreSearchResultElementMedcoEncryption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEncrypted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExploreSearchResultElementMedcoEncryption) validateEncrypted(formats strfmt.Registry) error {

	if err := validate.Required("medcoEncryption"+"."+"encrypted", "body", m.Encrypted); err != nil {
		return err
	}

	return nil
}

func (m *ExploreSearchResultElementMedcoEncryption) validateID(formats strfmt.Registry) error {

	if err := validate.Required("medcoEncryption"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExploreSearchResultElementMedcoEncryption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExploreSearchResultElementMedcoEncryption) UnmarshalBinary(b []byte) error {
	var res ExploreSearchResultElementMedcoEncryption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
