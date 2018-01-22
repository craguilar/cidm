package persistence

import (
	"errors"
	"github.com/craguilar/cidm/models"
	"github.com/craguilar/cidm/persistence/db"
	"github.com/craguilar/cidm/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"time"
)

const (
	//OauthType constant definition
	OauthType = "OAUTH2"
	//OauthGoogle constant definition for Oauth 2 google service
	OauthGoogle = "OAUTH2-GOOGLE"
)

var (
	errorExistLogin   = errors.New("login details already exist in the system")
	errorExistDef     = errors.New("default login already exist")
	errorOauth        = errors.New("could not get the default oauth2 token handler ")
	googleOauthConfig = &oauth2.Config{}
	oauthStateString  = GetEnvVariable("OAUTH2STRING")
)

func init() {
	initOauthTypes()
}

// GetDefaultLogin Returns the system default login for an specific running instance
func GetDefaultLogin() (models.Login, error) {
	var login models.Login
	conn := db.GetConnection()
	defer conn.Close()
	err := conn.Model(&login).
		Where("is_default = ?", true).
		Limit(1).
		Select()
	if err != nil {
		utils.Logger().Error(err)
		return login, err
	}
	login.CanonicalURL = GetServerAddress() + login.CanonicalURL
	return login, nil
}

// GetLoginByID get login details by Id
func GetLoginByID(id int64) (models.Login, error) {
	login := models.Login{ID: id}
	conn := db.GetConnection()
	err := conn.Select(&login)
	if err != nil {
		utils.Logger().Error(err)
		return login, err
	}
	login.CanonicalURL = GetServerAddress() + login.CanonicalURL
	return login, nil
}

// Insert a new login as Login model into the persisted storage
func Insert(loginID models.Login) (models.Login, error) {
	conn := db.GetConnection()

	// Simple params.
	count, err := conn.Model(&loginID).
		Where("id=?", loginID.ID).
		//WhereOr("is_default =  ?", true).
		Count()
	if err != nil {
		return models.Login{}, err
	}
	if count > 0 {
		return models.Login{}, errorExistLogin
	}
	if loginID.IsDefault && existDefault() {
		return models.Login{}, errorExistDef
	}
	err = conn.Insert(&loginID)
	if err != nil {
		return models.Login{}, err
	}
	insertLoginControlAttrs(loginID)
	//Insert
	return loginID, nil
}

// Update an existing login
func Update(login models.Login) (models.Login, error) {
	conn := db.GetConnection()

	// Simple params.
	err := conn.Update(&login)
	if err != nil {
		return login, err
	}
	//Insert
	updateLoginControlAttrs(login)
	return login, nil

}

// OauthAuthenticationURLByType returns the login authentication for an specific OAuth 2 service
func OauthAuthenticationURLByType(auth string) string {

	if auth == OauthGoogle {
		return googleOauthConfig.AuthCodeURL(GetOauthStateString())
	}
	return ""
}

// OauthAuthenticationTokenByType get a token based on the type of authentication defined , this will connect with
// external parties.
func OauthAuthenticationTokenByType(auth string, state string, code string) (*oauth2.Token, error) {
	var token *oauth2.Token
	if auth == OauthGoogle && state == GetOauthStateString() {
		return googleOauthConfig.Exchange(oauth2.NoContext, code)
	}
	return token, errorOauth
}

// GetOauthStateString get state string for oauth connection
func GetOauthStateString() string {
	return oauthStateString
}

func existDefault() bool {
	conn := db.GetConnection()
	count, err := conn.Model(&models.Login{}).
		Where("is_default = ?", true).
		//WhereOr("is_default =  ?", true).
		Count()
	if err != nil {
		return true
	}
	return count > 0
}

func updateLoginControlAttrs(login models.Login) bool {
	conn := db.GetConnection()
	_, err := conn.Model(&login).
		Set("modified_on = ?", time.Now()).
		Set("modified_by = 0").
		Where("id = ?id").
		Update()
	if err != nil {
		return false
	}
	return true

}

func insertLoginControlAttrs(login models.Login) bool {
	conn := db.GetConnection()
	_, err := conn.Model(&login).
		Set("created_on = ?", time.Now()).
		Set("created_by = 0").
		Where("id = ?id").
		Update()
	if err != nil {
		return false
	}
	return true

}

func initOauthTypes() {
	conn := db.GetConnection()
	gOauth := models.Authentication{}
	conn.Model(&gOauth).
		Where("authentication_type=?", OauthType).Select()
	googleOauthConfig.RedirectURL = GetServerAddress() + gOauth.RedirectUrl
	googleOauthConfig.ClientID = gOauth.OauthClientId
	googleOauthConfig.ClientSecret = gOauth.OauthClientSecret
	//TODO needs to be sent as persisted storage
	googleOauthConfig.Scopes = []string{"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email"}
	googleOauthConfig.Endpoint = google.Endpoint

}
