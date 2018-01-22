package services

import (
	"github.com/craguilar/cidm/models"
	"github.com/craguilar/cidm/persistence"
	"golang.org/x/oauth2"

	"github.com/craguilar/cidm/utils"
)

func LoginUser(token oauth2.Token) (models.User, error) {
	utils.Logger().Print("Get User from Persistence Layer.")

	//Get User from Persistence Layer.
	user, err := persistence.GoogleUserByToken(token)

	if err != nil {
		utils.Logger().Error(err)
		return user, err
	}
	utils.Logger().Print("Update and create a record in the DB.")

	//Update and create a record in the DB.
	err = persistence.CreateUser(user)

	if err != nil {
		utils.Logger().Error(err)
		return user, err
	}
	utils.Logger().Print("Token Service Creation , maybe part of a token repository?")

	//Token Service Creation , maybe part of a token repository?
	userToken := models.UserToken{}
	userToken.Id = token.AccessToken
	userToken.UserId = user.Id
	userToken.TokenType = "oauth2.google"
	userToken.Expiration = token.Expiry
	//
	utils.Logger().Print("UpdateToken")
	err = persistence.UpdateToken(userToken)
	if err != nil {
		utils.Logger().Error(err)
		return user, err
	}
	return user, nil

}
