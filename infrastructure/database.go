package infrastructure

import (
	"database/sql"
	"errors"
	"go-cli-bank/categories"
)

// DatabaseHandler handles database operations
type DatabaseHandler struct {
	path string
	*sql.DB
}

// Execute an sql statement
func (dh *DatabaseHandler) Execute(statement string) error {
	return errors.New("Failed to run database statement")
}

// Query fetches data from the database
func (dh *DatabaseHandler) Query(statement string) (categories.IRow, error) {
	return nil, errors.New("Failed to query database")
}
