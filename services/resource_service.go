package services

import (
	"github.com/craguilar/cidm/persistence"
	"github.com/craguilar/cidm/restapi/operations/resource"
	"github.com/craguilar/cidm/utils"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

func GetResourceById(parameter resource.GetResourceByIDParams, principal interface{}) middleware.Responder {
	id := parameter.ID
	resourceObj, err := persistence.GetResourceByID(id)
	if err != nil {
		return resource.NewGetResourceByIDDefault(http.StatusBadRequest).
			WithPayload(utils.GetError(http.StatusBadRequest, err.Error()))
	}
	return resource.NewGetResourceByIDOK().WithPayload(&resourceObj)
}
func AddResource(parameter resource.AddResourceParams, principal interface{}) middleware.Responder {
	resourceObj := parameter.Resource
	_, err := persistence.AddResource(*resourceObj)
	if err != nil {
		return resource.NewAddResourceDefault(http.StatusBadRequest).
			WithPayload(utils.GetError(http.StatusBadRequest, err.Error()))
	}
	return resource.NewAddResourceOK().WithPayload(resourceObj)
}
func UpdateResource(parameter resource.UpdateResourceParams, principal interface{}) middleware.Responder {

	resourceObj := parameter.Resource
	_, err := persistence.UpdateResource(*resourceObj)
	if err != nil {
		return resource.NewUpdateResourceDefault(http.StatusBadRequest).
			WithPayload(utils.GetError(http.StatusBadRequest, err.Error()))
	}
	return resource.NewUpdateResourceOK().WithPayload(resourceObj)
}
func DeleteResource(parameter resource.DeleteResourceParams, principal interface{}) middleware.Responder {
	resourceObj := parameter.Resource
	_, err := persistence.DeleteResource(*resourceObj)
	if err != nil {
		return resource.NewDeleteResourceDefault(http.StatusBadRequest).
			WithPayload(utils.GetError(http.StatusBadRequest, err.Error()))
	}
	return resource.NewDeleteResourceOK().WithPayload(resourceObj)
}
