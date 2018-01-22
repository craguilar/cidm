package models

import (
	"fmt"
	"github.com/go-openapi/strfmt"
	"time"
)

type UserToken struct {
	Id string `json:"tokenId"`
	// email
	UserId string
	// username
	TokenType string `json:"tokenType"`
	//Names
	Expiration time.Time
	// first name: json can come as given_name

}

func (ut *UserToken) String() string {
	return fmt.Sprintf("UserToken<%s %s %s>", ut.Id, ut.UserId, ut.TokenType)
}

// Validate validates this user
func (ut *UserToken) Validate(formats strfmt.Registry) error {

	return nil
}
