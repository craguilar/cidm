package services

import (
	"github.com/craguilar/cidm/models"
	"github.com/craguilar/cidm/persistence"
	"github.com/craguilar/cidm/restapi/operations/login"
	"github.com/craguilar/cidm/utils"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

// GetVariable allows yu to get an Enviroment variable with a prefixed value: PREFIX
func GetVariable(key string) string {
	return persistence.GetEnvVariable(key)
}

// GetLoginConfigDetails allows you to get a persisted config in the attached storage
func GetLoginConfigDetails(params login.GetLoginConfigDetailsParams, principal interface{}) middleware.Responder {
	utils.Logger().Print("hit:" + params.HTTPRequest.URL.Path)
	name := params.ID
	config, err := persistence.GetPersistedConfig(name)
	//Get default values from the repository
	if err != nil {
		errorm := err.Error()
		utils.Logger().Error(errorm)
		return login.NewGetLoginConfigDetailsDefault(http.StatusNotFound).
			WithPayload(&models.Error{Code: http.StatusNotFound, Message: &errorm})
	}
	//Return the login details
	return login.NewGetLoginConfigDetailsOK().WithPayload(&config)

}

// UpdateLoginConfigDetails allows you to update an specific variable in attached storage: DB
func UpdateLoginConfigDetails(params login.UpdateLoginConfigDetailsParams, principal interface{}) middleware.Responder {
	utils.Logger().Print("hit:" + params.HTTPRequest.URL.Path)

	config := params.Configuration
	err := persistence.UpdatePersistedConfig(*config)
	if err != nil {
		errorm := err.Error()
		utils.Logger().Error(errorm)
		return login.NewUpdateLoginConfigDetailsDefault(http.StatusBadRequest).
			WithPayload(&models.Error{Code: http.StatusBadRequest, Message: &errorm})
	}
	return login.NewUpdateLoginConfigDetailsCreated()
}

// GetServerAddress get server address
func GetServerAddress() string {
	return persistence.GetServerAddress()
}
