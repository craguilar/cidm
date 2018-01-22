package utils

import (
	"errors"
	"github.com/craguilar/cidm/models"
)

var (
	ErrorNoDB = errors.New("not able to connect to db")
)

//
func GetError(code int32, message string) *models.Error {
	return &models.Error{Code: code, Message: &message}
}
