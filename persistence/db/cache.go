package db

import (
	"github.com/craguilar/cidm/utils"
	"github.com/mediocregopher/radix.v2/pool"
	"os"
	"strconv"
)

var (
	db           *pool.Pool
	redisServer  string
	redisPort    string
	redisPool    int
	redisNetwork string
)

func init() {

	utils.Logger().Info("Initialize redis server " + serverName(redisServer, redisPort))
	loadCacheParameters()
	var err error
	db, err = pool.New(redisNetwork, serverName(redisServer, redisPort), redisPool)
	if err != nil {
		utils.Logger().Error(err)
	} else {
		utils.Logger().Info("Finalized initializing redis is UP  " + serverName(redisServer, redisPort))
	}

}

// GetCacheConnection get Pool instance of cache connection
func GetCacheConnection() *pool.Pool {
	return db
}

// loadCacheParameters load config from enviromet variables
func loadCacheParameters() {
	redisServer = os.Getenv("CIDM_REDIS_SERVER")
	redisPort = os.Getenv("CIDM_REDIS_PORT")
	redisPool, _ = strconv.Atoi(os.Getenv("CIDM_REDISPOOL"))
	redisNetwork = os.Getenv("CIDM_REDIS_NETWORK")
	utils.Logger().Print("Param CACHE redis_server,redis_port,redispool,redis_network " +
		redisServer + "," + redisPort + "," + string(redisPool) + "," + redisNetwork)
}

// serverName get server Name
func serverName(server, port string) string {
	return server + ":" + port
}
