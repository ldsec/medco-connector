// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExploreQuery MedCo-Explore query
//
// swagger:model exploreQuery
type ExploreQuery struct {

	// i2b2 panels (linked by an AND)
	Panels []*Panel `json:"panels"`

	// query timing
	QueryTiming Timing `json:"queryTiming,omitempty"`

	// user public key
	// Pattern: ^[\w=-]+$
	UserPublicKey string `json:"userPublicKey,omitempty"`
}

// Validate validates this explore query
func (m *ExploreQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePanels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryTiming(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserPublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExploreQuery) validatePanels(formats strfmt.Registry) error {

	if swag.IsZero(m.Panels) { // not required
		return nil
	}

	for i := 0; i < len(m.Panels); i++ {
		if swag.IsZero(m.Panels[i]) { // not required
			continue
		}

		if m.Panels[i] != nil {
			if err := m.Panels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("panels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExploreQuery) validateQueryTiming(formats strfmt.Registry) error {

	if swag.IsZero(m.QueryTiming) { // not required
		return nil
	}

	if err := m.QueryTiming.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("queryTiming")
		}
		return err
	}

	return nil
}

func (m *ExploreQuery) validateUserPublicKey(formats strfmt.Registry) error {

	if swag.IsZero(m.UserPublicKey) { // not required
		return nil
	}

	if err := validate.Pattern("userPublicKey", "body", string(m.UserPublicKey), `^[\w=-]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExploreQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExploreQuery) UnmarshalBinary(b []byte) error {
	var res ExploreQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
