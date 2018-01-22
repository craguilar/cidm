package persistence

import (
	"encoding/json"

	"github.com/craguilar/cidm/models"
	"github.com/craguilar/cidm/persistence/db"
	"github.com/craguilar/cidm/utils"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

// GoogleUserByToken this procedure will take a valid token Id and return a user from Google  userinfo API
func GoogleUserByToken(token oauth2.Token) (models.User, error) {
	var user models.User

	response, err := http.Get(GetGoogleUserAPI() + token.AccessToken)

	if err != nil {
		utils.Logger().Error(err)
		return user, err
	}

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		utils.Logger().Error(err)
		return user, err
	}
	err = json.Unmarshal(contents, &user)

	if err != nil {
		utils.Logger().Error(err)
		return user, err
	}
	return user, nil
}

// CreateUser in persisted storage
func CreateUser(user models.User) error {
	conn := db.GetConnection()
	if conn == nil {
		return utils.ErrorNoDB
	}

	count, err := conn.Model(&user).Where("id = ?", user.Id).Count()

	utils.Logger().Printf(" user : %s , count= %d", user.Id, count)
	if err != nil {
		utils.Logger().Error(err)
		return err
	}
	if count > 0 {
		utils.Logger().Printf("User already exist no need to Insert")
		return nil
	}
	//If user does not exist
	err = conn.Insert(&user)
	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}

// GetToken from table
func GetToken(token models.UserToken) (models.UserToken, error) {
	conn := db.GetConnection()
	// Simple params.
	err := conn.Select(&token)
	if err != nil {
		utils.Logger().Error(err)
		return token, err
	}
	return token, nil
}

// UpdateToken existing token associated with a particular user
func UpdateToken(token models.UserToken) error {
	conn := db.GetConnection()
	// Simple params.
	count, err := conn.Model(&token).
		Where("user_id=?", token.UserId).
		Count()
	if err != nil {
		return err
	}
	utils.Logger().Printf(" token : %s , count=%d", token.UserId, count)
	//Insert
	if count == 0 {
		return conn.Insert(&token)
	}
	return nil
}

//GetGoogleUserAPI get API Uri end point
func GetGoogleUserAPI() string {
	return UserInfoAPI
}
