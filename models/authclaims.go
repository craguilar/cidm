package models

import (
	"github.com/dgrijalva/jwt-go"
)

type AuthClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
