package db

import (
	"github.com/craguilar/cidm/utils"
	"github.com/go-pg/pg"
	"os"
)

var (
	rdbIP     string
	rdbUser   string
	rdbDBname string
	rdbSchema string
	password  string
)

func init() {
	utils.Logger().Info("Initialize DB server ")
	password, _ = GetSecret("DBPASSWORD")
	loadParameters()
	utils.Logger().Info("Finalized initializing  DB configured " + rdbIP)
}

// GetConnection get a pooled connection for Persisted storage
func GetConnection() *pg.DB {
	params := make(map[string]interface{})
	params["search_path"] = rdbSchema

	db := pg.Connect(&pg.Options{
		Addr:     rdbIP,
		User:     rdbUser,
		Password: password,
		Database: rdbDBname,
	})
	_, err := db.Exec("SET search_path = " + rdbSchema)
	if err != nil {
		utils.Logger().Error(err)
		return nil
	}
	return db
}

func loadParameters() {
	rdbIP = os.Getenv("CIDM_DB_URL")
	rdbUser = os.Getenv("CIDM_DB_USER")
	rdbDBname = os.Getenv("CIDM_DB_NAME")
	rdbSchema = os.Getenv("CIDM_SCHEMA")
	utils.Logger().Print("Param DB rdb_ip,rdb_user,rdbDB name " + rdbIP + "," + rdbUser + "," + rdbDBname)
}
