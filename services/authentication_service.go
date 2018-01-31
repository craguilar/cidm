package services

import (
	"errors"
	"github.com/craguilar/cidm/models"
	"github.com/craguilar/cidm/restapi/operations/token"
	"github.com/craguilar/cidm/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	jwtTokenExpiration = "JWTEXPIRATION"
	jtwTokenSubject    = "JWTSUBJECT"
	jwtConst           = "JWT"
)

var (
	ErrorParseToken    = errors.New("not able to parse Authentication header")
	ErrorSplitToken    = errors.New("error in split bearer auth")
	ErrorSigningMethod = errors.New("unexpected signing method")
)

func HandleBearerAuth(token string) (interface{}, error) {
	utils.Logger().Info("Authentication Handler")
	jwtToken, err := SplitBearerAuth(token)
	if err != nil {
		return false, err
	}
	if _, err := validToken(jwtToken); err != nil {
		return false, err
	}
	return true, nil
}

func HandleTokenValidate(params token.TokenValidateParams) middleware.Responder {
	utils.Logger().Print("hit:" + params.HTTPRequest.URL.Path)
	authorization := params.Authorization
	jwtToken, err := SplitBearerAuth(authorization)

	if err != nil {
		e := err.Error()
		return token.NewTokenValidateDefault(http.StatusBadRequest).
			WithPayload(&models.Error{Message: &e, Code: -1})
	}
	if exist, _ := FindKey(jwtToken); exist {
		utils.Logger().Print("Cached token:" + jwtToken)
		return token.NewTokenValidateOK()
	}
	if _, err := validToken(jwtToken); err != nil {
		e := err.Error()
		return token.NewTokenValidateDefault(http.StatusBadRequest).
			WithPayload(&models.Error{Message: &e, Code: -1})
	} else {
		SetKey(jwtToken, "true")
		return token.NewTokenValidateOK()
	}
}

func HandleTokenPayloadValidate(params token.TokenValidatPayloadParams) middleware.Responder {
	utils.Logger().Print("hit:" + params.HTTPRequest.URL.Path)
	//TODO HandleTokenPayloadValidate not implemented
	return middleware.NotImplemented("Not implemented")
}

// GetTokenString
func GetTokenString(username string) (string, error) {
	// Create a new token object, specifying signing method and the claims you would like it to contain.
	expiration := time.Now().AddDate(0, 0, getExpirationTime()).Unix()
	subject := GetVariable(jtwTokenSubject)
	// Create the Claims
	claims := models.AuthClaims{}
	claims.Username = username
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: int64(expiration),
		Issuer:    GetServerAddress(),
		Subject:   subject,
		IssuedAt:  int64(time.Now().Unix()),
	}
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	return tokens.SignedString(getSecret())
}

func validToken(tokenString string) (bool, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrorSigningMethod
		}
		return getSecret(), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil

}

// GetTokenClaims decode token claims
func GetTokenClaims(tokenStr string) (models.AuthClaims, error) {

	tokens, err := jwt.ParseWithClaims(tokenStr, &models.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrorSigningMethod
		}
		return getSecret(), nil
	})

	if claims, ok := tokens.Claims.(*models.AuthClaims); ok && tokens.Valid {
		return *claims, nil
	} else {
		return *claims, err
	}
}

func SplitBearerAuth(token string) (string, error) {
	tokenSplit := strings.Split(token, " ")
	if len(tokenSplit) != 2 {
		return "", ErrorParseToken
	}
	if strings.ToLower(tokenSplit[0]) == "bearer" {
		return tokenSplit[1], nil
	}
	return "", ErrorSplitToken
}

func MultiFactor() {

}

func getSecret() []byte {
	key, _ := GetSecret(jwtConst)
	return []byte(key)
}

func getExpirationTime() int {
	days, err := strconv.Atoi(GetVariable(jwtTokenExpiration))
	if err != nil {
		days = 1
	}
	return days
}
