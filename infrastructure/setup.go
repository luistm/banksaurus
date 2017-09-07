package infrastructure

import (
	"database/sql"
	"errors"
)

var errSetupFailed = errors.New("Failed to setup storage")

const DATABASE_NAME = "~/.expensetracker.db"
const DATABASE_ENGINE = "sqlite3"

// SetupStorage sets up the specified storage mechanism
func SetupStorage(storageType string) (*sql.DB, error) {

	if storageType == "" {
		return nil, errSetupFailed
	}

	return nil, errSetupFailed
}
