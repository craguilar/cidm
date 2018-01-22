package db

import "os"

//TODO : Need t design the KM algorithm to securely distribute keys
var m map[string]string

func init() {
	m = make(map[string]string)
	m["JWT"] = os.Getenv("CIDM_JWT_SECRET")
	m["DBPASSWORD"] = os.Getenv("CIDM_DB_PASSWORD")
}

// GetSecret get a secret from KMS
func GetSecret(key string) (string, error) {
	return m[key], nil
}
