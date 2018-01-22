package persistence

import (
	"github.com/craguilar/cidm/models"
	"github.com/craguilar/cidm/persistence/db"
	"github.com/craguilar/cidm/utils"
)

// IsAuthorized backbone function to check authorization.
func IsAuthorized(user models.User, resource models.Resource, operation string) bool {
	_, err := GetPermission(user, resource, operation)
	if err != nil {
		return false
	}
	return true
}

// GetPermission critical function
func GetPermission(user models.User, resource models.Resource, operation string) (models.Permission, error) {
	conn := db.GetConnection()
	defer conn.Close()
	if conn == nil {
		return models.Permission{}, utils.ErrorNoDB
	}
	var permission models.Permission

	//Search for specific mach
	permission = models.Permission{
		Email:     &user.Id,
		Resource:  &resource,
		Operation: &operation,
		Enabled:   true}
	//Search for specific mach
	err := conn.Model(&permission).Select()
	if err == nil {
		return permission, nil
	}

	//Handle *
	err = conn.Model(&permission).
		Where("email = ?", user.Id).
		Where("resource = ?", "*").
		Where("resource = ?", operation).
		Where("enabled = ?", true).
		Limit(1).
		Select()
	//Return immediately
	if err != nil {
		utils.Logger().Error(utils.ErrorNoDB)
		return permission, err
	}
	return permission, nil
}

// GetPermissionsByUser get permission given a particular user.
func GetPermissionsByUser(user models.User) ([]models.Permission, error) {
	conn := db.GetConnection()
	defer conn.Close()

	if conn == nil {
		return []models.Permission{}, utils.ErrorNoDB
	}
	var permissions []models.Permission
	err := conn.Model(&permissions).
		Where("email = ?", user.Id).
		Where("enabled = ?", true).
		Column("permission.*", "Resource").
		Select()
	if err != nil {
		utils.Logger().Error(utils.ErrorNoDB)
		return permissions, err
	}
	return permissions, nil
}

// AddResourceToUser add resource to a particular User.
func AddResourceToUser(user models.User, resource models.Resource, operation string) (models.Permission, error) {
	conn := db.GetConnection()
	defer conn.Close()
	permission := models.Permission{
		Email:     &user.Id,
		Resource:  &resource,
		Operation: &operation,
		Enabled:   true}
	err := conn.Insert(&permission)

	if err != nil {
		utils.Logger().Error(utils.ErrorNoDB)
		return permission, err
	}
	return models.Permission{}, nil
}

// UpdatePermission update an existin permission for a particular user.
func UpdatePermission(permission models.Permission) (models.Permission, error) {
	conn := db.GetConnection()
	defer conn.Close()
	err := conn.Update(&permission)
	if err != nil {
		return models.Permission{}, err
	}
	return permission, nil
}

// DeletePermission  an existing permission for a particular user.
func DeletePermission(permission models.Permission) (models.Permission, error) {
	conn := db.GetConnection()
	defer conn.Close()
	return permission, nil
}
