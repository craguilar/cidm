package models

import (
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"gopkg.in/pg.v5/orm"
	"log"
	"time"
)

type User struct {
	Id string `json:"email"`
}

func (u *User) String() string {
	return fmt.Sprintf("User<%s>", u.Id)
}

// Validate validates this user
func (u *User) Validate(formats strfmt.Registry) error {
	var res []error
	if u.Id == "" {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

//db Insert stmts
func (u *User) AfterInsert(db orm.DB) error {
	log.Println("Before Insert: User id: " + u.Id)
	conn := db
	_, err := conn.Model(u).
		Set("created_on = ?", time.Now()).
		Set("created_by = ", u.Id).
		Where("id = ?id").
		Update()
	if err != nil {
		return err
	}
	return nil
}
