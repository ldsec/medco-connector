// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Panel panel
//
// swagger:model panel
type Panel struct {

	// i2b2 items (linked by an OR)
	Items []*PanelItemsItems0 `json:"items"`

	// exclude the i2b2 panel
	// Required: true
	Not *bool `json:"not"`

	// panel timing
	PanelTiming Timing `json:"panelTiming,omitempty"`
}

// Validate validates this panel
func (m *Panel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePanelTiming(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Panel) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Panel) validateNot(formats strfmt.Registry) error {

	if err := validate.Required("not", "body", m.Not); err != nil {
		return err
	}

	return nil
}

func (m *Panel) validatePanelTiming(formats strfmt.Registry) error {

	if swag.IsZero(m.PanelTiming) { // not required
		return nil
	}

	if err := m.PanelTiming.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("panelTiming")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Panel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Panel) UnmarshalBinary(b []byte) error {
	var res Panel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PanelItemsItems0 panel items items0
//
// swagger:model PanelItemsItems0
type PanelItemsItems0 struct {

	// encrypted
	// Required: true
	Encrypted *bool `json:"encrypted"`

	// modifier
	Modifier *PanelItemsItems0Modifier `json:"modifier,omitempty"`

	// # NUMBER operators EQ: equal NE: not equal GT: greater than GE: greater than or equal LT: less than LE: less than or equal BETWEEN: between (value syntax: "x and y") # TEXT operators IN: in (value syntax: "'x','y','z'") LIKE[exact]: equal LIKE[begin]: begins with LIKE[end]: ends with LIKE[contains]: contains
	//
	// Enum: [EQ NE GT GE LT LE BETWEEN IN LIKE[exact] LIKE[begin] LIKE[end] LIKE[contains]]
	Operator string `json:"operator,omitempty"`

	// query term
	// Required: true
	// Pattern: ^([\w=-]+)$|^((\/[^\/]+)+\/)$
	QueryTerm *string `json:"queryTerm"`

	// type
	// Enum: [NUMBER TEXT]
	Type string `json:"type,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this panel items items0
func (m *PanelItemsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEncrypted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryTerm(formats); err != nil {
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

func (m *PanelItemsItems0) validateEncrypted(formats strfmt.Registry) error {

	if err := validate.Required("encrypted", "body", m.Encrypted); err != nil {
		return err
	}

	return nil
}

func (m *PanelItemsItems0) validateModifier(formats strfmt.Registry) error {

	if swag.IsZero(m.Modifier) { // not required
		return nil
	}

	if m.Modifier != nil {
		if err := m.Modifier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifier")
			}
			return err
		}
	}

	return nil
}

var panelItemsItems0TypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EQ","NE","GT","GE","LT","LE","BETWEEN","IN","LIKE[exact]","LIKE[begin]","LIKE[end]","LIKE[contains]"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		panelItemsItems0TypeOperatorPropEnum = append(panelItemsItems0TypeOperatorPropEnum, v)
	}
}

const (

	// PanelItemsItems0OperatorEQ captures enum value "EQ"
	PanelItemsItems0OperatorEQ string = "EQ"

	// PanelItemsItems0OperatorNE captures enum value "NE"
	PanelItemsItems0OperatorNE string = "NE"

	// PanelItemsItems0OperatorGT captures enum value "GT"
	PanelItemsItems0OperatorGT string = "GT"

	// PanelItemsItems0OperatorGE captures enum value "GE"
	PanelItemsItems0OperatorGE string = "GE"

	// PanelItemsItems0OperatorLT captures enum value "LT"
	PanelItemsItems0OperatorLT string = "LT"

	// PanelItemsItems0OperatorLE captures enum value "LE"
	PanelItemsItems0OperatorLE string = "LE"

	// PanelItemsItems0OperatorBETWEEN captures enum value "BETWEEN"
	PanelItemsItems0OperatorBETWEEN string = "BETWEEN"

	// PanelItemsItems0OperatorIN captures enum value "IN"
	PanelItemsItems0OperatorIN string = "IN"

	// PanelItemsItems0OperatorLIKEExact captures enum value "LIKE[exact]"
	PanelItemsItems0OperatorLIKEExact string = "LIKE[exact]"

	// PanelItemsItems0OperatorLIKEBegin captures enum value "LIKE[begin]"
	PanelItemsItems0OperatorLIKEBegin string = "LIKE[begin]"

	// PanelItemsItems0OperatorLIKEEnd captures enum value "LIKE[end]"
	PanelItemsItems0OperatorLIKEEnd string = "LIKE[end]"

	// PanelItemsItems0OperatorLIKEContains captures enum value "LIKE[contains]"
	PanelItemsItems0OperatorLIKEContains string = "LIKE[contains]"
)

// prop value enum
func (m *PanelItemsItems0) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, panelItemsItems0TypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PanelItemsItems0) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *PanelItemsItems0) validateQueryTerm(formats strfmt.Registry) error {

	if err := validate.Required("queryTerm", "body", m.QueryTerm); err != nil {
		return err
	}

	if err := validate.Pattern("queryTerm", "body", string(*m.QueryTerm), `^([\w=-]+)$|^((\/[^\/]+)+\/)$`); err != nil {
		return err
	}

	return nil
}

var panelItemsItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NUMBER","TEXT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		panelItemsItems0TypeTypePropEnum = append(panelItemsItems0TypeTypePropEnum, v)
	}
}

const (

	// PanelItemsItems0TypeNUMBER captures enum value "NUMBER"
	PanelItemsItems0TypeNUMBER string = "NUMBER"

	// PanelItemsItems0TypeTEXT captures enum value "TEXT"
	PanelItemsItems0TypeTEXT string = "TEXT"
)

// prop value enum
func (m *PanelItemsItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, panelItemsItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PanelItemsItems0) validateType(formats strfmt.Registry) error {

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
func (m *PanelItemsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PanelItemsItems0) UnmarshalBinary(b []byte) error {
	var res PanelItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PanelItemsItems0Modifier panel items items0 modifier
//
// swagger:model PanelItemsItems0Modifier
type PanelItemsItems0Modifier struct {

	// applied path
	// Required: true
	// Pattern: ^((\/[^\/]+)+\/%?)$
	AppliedPath *string `json:"appliedPath"`

	// modifier key
	// Required: true
	// Pattern: ^((\/[^\/]+)+\/)$
	ModifierKey *string `json:"modifierKey"`
}

// Validate validates this panel items items0 modifier
func (m *PanelItemsItems0Modifier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppliedPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifierKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PanelItemsItems0Modifier) validateAppliedPath(formats strfmt.Registry) error {

	if err := validate.Required("modifier"+"."+"appliedPath", "body", m.AppliedPath); err != nil {
		return err
	}

	if err := validate.Pattern("modifier"+"."+"appliedPath", "body", string(*m.AppliedPath), `^((\/[^\/]+)+\/%?)$`); err != nil {
		return err
	}

	return nil
}

func (m *PanelItemsItems0Modifier) validateModifierKey(formats strfmt.Registry) error {

	if err := validate.Required("modifier"+"."+"modifierKey", "body", m.ModifierKey); err != nil {
		return err
	}

	if err := validate.Pattern("modifier"+"."+"modifierKey", "body", string(*m.ModifierKey), `^((\/[^\/]+)+\/)$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PanelItemsItems0Modifier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PanelItemsItems0Modifier) UnmarshalBinary(b []byte) error {
	var res PanelItemsItems0Modifier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
