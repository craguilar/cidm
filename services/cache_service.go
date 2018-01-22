package services

import (
	"github.com/craguilar/cidm/persistence/db"
	"github.com/craguilar/cidm/utils"
)

// FindKey allows you to get a cached key in cache service if available
func FindKey(id string) (bool, error) {

	// Use the connection pool's Get() method to fetch a single Redis .
	conn, err := db.GetCacheConnection().Get()
	if err != nil {
		return false, err
	}
	// Before we do anything else, check that an album with the given id
	// exists. The EXISTS command returns 1 if a specific key exists
	// in the database, and 0 if it doesn't.
	exists, err := conn.Cmd("EXISTS", id).Int()

	if err != nil {
		return false, err //log.Fatal(err)
	}
	return exists == 1, nil
}

// SetKey allows you to set a cached key in cache service if available
func SetKey(key string, value string) error {
	conn, err := db.GetCacheConnection().Get()
	if err != nil {
		return err
	}
	resp := conn.Cmd("SET", key, value)
	// Check the Err field of the *Resp object for any errors.
	if resp.Err != nil {
		utils.Logger().Error(resp.Err)
		return resp.Err
	}
	return nil
}
