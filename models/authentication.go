package models

import (
	"github.com/go-openapi/strfmt"
)

type Authentication struct {
	Id int `json:"id"`
	// additional properties
	AuthenticationType string `json:"authenticationType"`
	//
	AuthenticationProvider string `json:"authenticationProvider"`
	// login title
	LoginUrl string `json:"loginUrl,omitempty"`

	// login type
	RedirectUrl string `json:"redirectUrl,omitempty"`

	OauthClientId string

	OauthClientSecret string

	Enabled bool
}

// Validate validates this login
func (m *Authentication) Validate(formats strfmt.Registry) error {

	return nil
}
