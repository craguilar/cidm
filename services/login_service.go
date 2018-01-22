package services

import (
	"errors"
	"github.com/craguilar/cidm/persistence"
	"github.com/craguilar/cidm/restapi/operations/login"
	"github.com/craguilar/cidm/utils"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
	"strings"
)

var (
	ErrorBodyNull             = errors.New("body for request is null")
	ErrorBodyNotAllowedMethod = errors.New("no method supported")
	ErrorInvalidOauthState    = errors.New("invalid oauth state")
)

//Get Login details provide secure login details for consumers to use this API in front end  applications, this function
// returns an object of type models.
func GetLoginDetails(params login.GetLoginDetailsParams) middleware.Responder {
	utils.Logger().Print("hit:" + params.HTTPRequest.URL.Path)
	//Get default values from the repository
	details, err := persistence.GetDefaultLogin()
	if err != nil {
		utils.Logger().Error(err)
		return login.NewGetLoginDefault(http.StatusNotFound).
			WithPayload(utils.GetError(http.StatusNotFound, err.Error()))
	}
	//Return the login details
	return login.NewGetLoginDetailsOK().WithPayload(&details)
}
func GetLoginDetailsById(params login.GetLoginDetailsByIDParams) middleware.Responder {
	utils.Logger().Print("hit:" + params.HTTPRequest.URL.Path)

	details, err := persistence.GetLoginByID(params.ID)
	if err != nil {
		utils.Logger().Error(err)
		return login.NewGetLoginDetailsByIDDefault(http.StatusNotFound).
			WithPayload(utils.GetError(http.StatusNotFound, err.Error()))
	}

	return login.NewGetLoginDetailsByIDOK().WithPayload(&details)
}

func AddLoginDetails(params login.AddLoginDetailsParams, principal interface{}) middleware.Responder {
	utils.Logger().Print("hit:" + params.HTTPRequest.URL.Path)
	_, err := persistence.Insert(*params.Login)
	if err != nil {
		return login.NewAddLoginDetailsDefault(http.StatusBadRequest).
			WithPayload(utils.GetError(http.StatusNotFound, err.Error()))
	}
	return login.NewAddLoginDetailsCreated()
}
func UpdateLoginDetails(params login.UpdateLoginDetailsParams, principal interface{}) middleware.Responder {
	utils.Logger().Print("hit:" + params.HTTPRequest.URL.Path)
	_, err := persistence.Update(*params.Login)
	if err != nil {
		return login.NewUpdateLoginDetailsDefault(http.StatusBadRequest).
			WithPayload(utils.GetError(http.StatusBadRequest, err.Error()))
	}
	return login.NewUpdateLoginDetailsCreated()
}

//HandleOAuth2Login handles the OAuth for google  or any other supported OAuth service.
func HandleLoginRequest(params login.GetLoginParams) middleware.Responder {
	utils.Logger().Print("hit:" + params.HTTPRequest.URL.Path)

	if params.Body == nil {
		utils.Logger().Error(ErrorBodyNull)
		return login.NewGetLoginDefault(http.StatusBadRequest).WithPayload(utils.GetError(http.StatusBadRequest, ErrorBodyNull.Error()))
	}

	details := params.Body
	if strings.ToLower(details.Type) == strings.ToLower(persistence.OauthGoogle) {
		url := persistence.OauthAuthenticationURLByType(persistence.OauthGoogle)
		utils.Logger().Print(url)
		return login.NewGetLoginFound().WithLocation(url)
	}

	return login.NewGetLoginDefault(http.StatusNotFound).
		WithPayload(utils.GetError(http.StatusNotFound, ErrorBodyNotAllowedMethod.Error()))

}

//Handle login for tes enviroments with a mock server
func HandleLoginMockRequest(params login.GetLoginMockParams) middleware.Responder {
	utils.Logger().Print("hit:" + params.HTTPRequest.URL.Path)
	url := persistence.OauthAuthenticationURLByType(persistence.OauthGoogle)
	utils.Logger().Print(url)
	return login.NewGetLoginFound().WithLocation(url)
}

func HandleCallback(params login.GetLoginCallbackParams) middleware.Responder {
	utils.Logger().Print("hit:" + params.HTTPRequest.URL.Path)
	state := params.HTTPRequest.FormValue("state")
	if state != persistence.GetOauthStateString() {

		utils.Logger().Error(ErrorInvalidOauthState)
		login.NewGetLoginCallbackDefault(http.StatusBadRequest).
			WithStatusCode(http.StatusBadRequest).
			WithPayload(utils.GetError(http.StatusBadRequest, ErrorInvalidOauthState.Error()))
	}
	code := params.HTTPRequest.FormValue("code")
	token, err := persistence.OauthAuthenticationTokenByType(persistence.OauthGoogle, state, code)
	if err != nil {

		login.NewGetLoginCallbackDefault(http.StatusBadRequest).
			WithStatusCode(http.StatusBadRequest).
			WithPayload(utils.GetError(http.StatusBadRequest, err.Error()))
	}

	user, err := LoginUser(*token)

	if err != nil {

		login.NewGetLoginCallbackDefault(http.StatusBadRequest).
			WithStatusCode(http.StatusBadRequest).
			WithPayload(utils.GetError(http.StatusBadRequest, err.Error()))
	}
	jwt, err := GetTokenString(strings.ToLower(user.Id))
	if err != nil {

		login.NewGetLoginCallbackDefault(http.StatusBadRequest).
			WithStatusCode(http.StatusBadRequest).
			WithPayload(utils.GetError(http.StatusBadRequest, err.Error()))
	}
	utils.Logger().Print("User " + strings.ToLower(user.Id) + " Jwt :" + jwt)
	//Get callback url
	callback, err := persistence.GetPersistedConfig("oauth2callbackurl")
	if err != nil {

		login.NewGetLoginCallbackDefault(http.StatusBadRequest).
			WithStatusCode(http.StatusBadRequest).
			WithPayload(utils.GetError(http.StatusBadRequest, err.Error()))
	}
	return login.NewGetLoginCallbackFound().
		WithAuthorization("Bearer " + jwt).
		WithLocation(callback.Value)
}
