// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// User user
//
// swagger:model user
type User struct {

	// authorizations
	Authorizations *UserAuthorizations `json:"authorizations,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorizations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateAuthorizations(formats strfmt.Registry) error {

	if swag.IsZero(m.Authorizations) { // not required
		return nil
	}

	if m.Authorizations != nil {
		if err := m.Authorizations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authorizations")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UserAuthorizations user authorizations
//
// swagger:model UserAuthorizations
type UserAuthorizations struct {

	// explore query
	ExploreQuery []ExploreQueryType `json:"exploreQuery"`

	// rest Api
	RestAPI []RestAPIAuthorization `json:"restApi"`
}

// Validate validates this user authorizations
func (m *UserAuthorizations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExploreQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestAPI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserAuthorizations) validateExploreQuery(formats strfmt.Registry) error {

	if swag.IsZero(m.ExploreQuery) { // not required
		return nil
	}

	for i := 0; i < len(m.ExploreQuery); i++ {

		if err := m.ExploreQuery[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authorizations" + "." + "exploreQuery" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserAuthorizations) validateRestAPI(formats strfmt.Registry) error {

	if swag.IsZero(m.RestAPI) { // not required
		return nil
	}

	for i := 0; i < len(m.RestAPI); i++ {

		if err := m.RestAPI[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authorizations" + "." + "restApi" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserAuthorizations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserAuthorizations) UnmarshalBinary(b []byte) error {
	var res UserAuthorizations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
