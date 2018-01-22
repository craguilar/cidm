package services

import (
	"github.com/craguilar/cidm/persistence/db"
)

// GetSecret get a secret from KMS
func GetSecret(key string) (string, error) {
	return db.GetSecret(key)
}
