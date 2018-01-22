package persistence

import (
	"errors"
	"github.com/craguilar/cidm/models"
	"github.com/craguilar/cidm/persistence/db"
	"github.com/craguilar/cidm/utils"
)

var (
	errorAddResource    = errors.New("error when adding resource, ether a db issue or the resource already exist")
	errorUpdateResource = errors.New("error when updating resource, ether a db issue or the resource not exist")
)

// GetResourceByID get resource by its Id that could be represented as a canonical uri.
func GetResourceByID(id string) (models.Resource, error) {
	resource := models.Resource{ID: &id}
	conn := db.GetConnection()
	defer conn.Close()
	err := conn.Model(&resource).
		Select()
	if err != nil {
		utils.Logger().Error(err)
		return models.Resource{}, err
	}
	return resource, nil
}

// AddResource to the global data base
func AddResource(resource models.Resource) (models.Resource, error) {
	conn := db.GetConnection()
	defer conn.Close()
	if conn == nil {
		return resource, utils.ErrorNoDB
	}
	count, err := conn.Model(&resource).Count()
	if err != nil || count > 0 {
		utils.Logger().Error(errorAddResource)
		return resource, errorAddResource
	}

	err = conn.Insert(&resource)
	if err != nil {
		utils.Logger().Error(err)
		return resource, err
	}
	return resource, nil
}

// UpdateResource in the DB
func UpdateResource(resource models.Resource) (models.Resource, error) {
	conn := db.GetConnection()
	defer conn.Close()
	if conn == nil {
		return resource, utils.ErrorNoDB
	}
	count, err := conn.Model(&resource).Count()
	if err != nil || count == 0 {
		return resource, errorUpdateResource
	}
	err = conn.Update(&resource)
	if err != nil {
		return resource, err
	}
	return resource, nil
}

// DeleteResource in the data base.
func DeleteResource(resource models.Resource) (models.Resource, error) {
	conn := db.GetConnection()
	defer conn.Close()
	if conn == nil {
		return resource, utils.ErrorNoDB
	}
	err := conn.Delete(&resource)
	if err != nil {
		return resource, err
	}
	return resource, nil
}
