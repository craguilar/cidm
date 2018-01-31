package services

import (
	"github.com/craguilar/cidm/models"
	"github.com/craguilar/cidm/persistence"
	"github.com/craguilar/cidm/restapi/operations/token"
	"github.com/craguilar/cidm/utils"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

func Authorization(a http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.Logger().Info("Authorization handler")
		a.ServeHTTP(w, r)
	})
}

//GetUserPermission
func GetUserPermission(parameter token.GetUserPermissionParams, principal interface{}) middleware.Responder {
	user, err := getUserFromBearer(parameter.Authorization)
	if err != nil {
		return token.NewGetUserPermissionDefault(http.StatusBadRequest).
			WithPayload(utils.GetError(http.StatusBadRequest, err.Error()))
	}
	_, err = persistence.GetPermissionsByUser(user)
	var permissions models.GetUserPermissionOKBody
	//TODO Needs to translate to permissions[0] = &obj
	if err != nil {
		return token.NewGetUserPermissionDefault(http.StatusBadRequest).
			WithPayload(utils.GetError(http.StatusBadRequest, err.Error()))
	}
	return token.NewGetUserPermissionOK().WithPayload(permissions)
}

//AddUserPermission
func AddUserPermission(parameter token.AddUserPermissionParams, principal interface{}) middleware.Responder {
	user, err := getUserFromBearer(parameter.Authorization)
	if err != nil {
		return token.NewAddUserPermissionDefault(http.StatusBadRequest).
			WithPayload(utils.GetError(http.StatusBadRequest, err.Error()))
	}
	permission := parameter.Permission
	_, err = persistence.AddResourceToUser(user, *permission.Resource, *permission.Operation)
	if err != nil {
		return token.NewAddUserPermissionDefault(http.StatusBadRequest).
			WithPayload(utils.GetError(http.StatusBadRequest, err.Error()))
	}
	return token.NewAddUserPermissionOK()
}

//UpdateUserPermission
func UpdateUserPermission(parameter token.UpdateUserPermissionParams, principal interface{}) middleware.Responder {
	permission := parameter.Permission
	_, err := persistence.UpdatePermission(*permission)
	if err != nil {
		return token.NewUpdateUserPermissionDefault(http.StatusBadRequest).
			WithPayload(utils.GetError(http.StatusBadRequest, err.Error()))
	}
	return token.NewUpdateUserPermissionOK()
}

//DeleteUserPermission
func DeleteUserPermission(parameter token.DeleteUserPermissionParams, principal interface{}) middleware.Responder {
	permission := parameter.Permission
	_, err := persistence.DeletePermission(*permission)
	if err != nil {
		return token.NewDeleteUserPermissionDefault(http.StatusBadRequest).
			WithPayload(utils.GetError(http.StatusBadRequest, err.Error()))
	}
	return token.NewDeleteUserPermissionOK()
}
func getUserFromBearer(bearer string) (models.User, error) {
	bearer, err := SplitBearerAuth(bearer)
	if err != nil {

	}
	claims, err := GetTokenClaims(bearer)
	if err != nil {
		return models.User{}, nil
	}
	return models.User{Id: claims.Username}, nil
}
