// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Login login
// swagger:model Login

type Login struct {

	// canonical Url
	CanonicalURL string `json:"canonicalUrl,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// is default
	IsDefault bool `json:"isDefault,omitempty"`

	// title
	// Required: true
	Title *string `json:"title"`

	// type
	Type string `json:"type,omitempty"`
}

/* polymorph Login canonicalUrl false */

/* polymorph Login enabled false */

/* polymorph Login id false */

/* polymorph Login isDefault false */

/* polymorph Login title false */

/* polymorph Login type false */

// Validate validates this login
func (m *Login) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTitle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Login) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Login) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Login) UnmarshalBinary(b []byte) error {
	var res Login
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
