package persistence

import (
	"github.com/craguilar/cidm/models"
	"github.com/craguilar/cidm/persistence/db"
	"github.com/craguilar/cidm/utils"
	"os"
	"strings"
	"time"
)

const (
	// Prefix Env variable used by micro service should all be prefixed
	Prefix = "CIDM"
	//UserInfoAPI Google API
	UserInfoAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
)

var (
	m map[string]string
	//TODO remove dependency for this server base
	serverBase = "http://localhost"
	portBase   = "30030"
)

// GetEnvVariable get enviroment variable from OS/ host system only variables prefixed with persitence . Prefix will be
// returned
func GetEnvVariable(key string) string {
	if m == nil {
		loadVariables()
	}
	return m[Prefix+"_"+key]
}

//GetServerBase gives the current server ip address.
func GetServerBase() string {
	return serverBase
}

//GetPortBase  Returns the current port base for the particular server.
func GetPortBase() string {
	return portBase
}

// GetServerAddress get global server Base.
func GetServerAddress() string {
	server := GetServerBase() + ":" + GetPortBase()
	return server
}

// GetPersistedConfig Returns a persisted configuration in the data base from configurations
func GetPersistedConfig(name string) (models.Configuration, error) {
	var config models.Configuration
	conn := db.GetConnection()
	defer conn.Close()
	err := conn.Model(&config).
		Where("id = ?", name).
		Where("enabled = ?", true).
		Limit(1).
		Select()
	if err != nil {
		utils.Logger().Error(err)
		return config, err
	}
	return config, nil
}

// UpdatePersistedConfig updates a persisted configuration value
func UpdatePersistedConfig(config models.Configuration) error {
	conn := db.GetConnection()
	defer conn.Close()
	err := conn.Update(&config)
	if err != nil {
		utils.Logger().Error(err)
		return err
	}
	_, err = conn.Model(&config).
		Set("modified_on = ?", time.Now()).
		Set("modified_by = 1").
		Where("id = ?id").
		Update()

	if err != nil {
		utils.Logger().Error(err)
		return err
	}
	return nil
}

func loadVariables() {
	m = make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if strings.Contains(pair[0], Prefix) {
			m[pair[0]] = pair[1]
		}
	}
}
