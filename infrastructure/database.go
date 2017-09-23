package infrastructure

import (
	"database/sql"
	"errors"
	"go-cli-bank/lib/categories"
)

// ErrDataBaseConnUndefined is to be returns when the
// database connection is not available
var ErrDataBaseConnUndefined = errors.New("Database connection is undefined")

// ErrDataBase to be used when the infrastructure
// database returns error
type ErrDataBase struct {
	s string
}

func (e *ErrDataBase) Error() string {
	return e.s
}

// DatabaseHandler handles database operations
type DatabaseHandler struct {
	dbConn *sql.DB
}

// Execute is to execute an sql statement
func (dh *DatabaseHandler) Execute(statement string) error {
	if dh.dbConn == nil {
		return ErrDataBaseConnUndefined
	}

	tx, _ := dh.dbConn.Begin()
	_, err := tx.Exec(statement)
	if err != nil {
		return &ErrDataBase{err.Error()}
	}
	tx.Commit()

	return nil
}

// Query fetches data from the database
func (dh *DatabaseHandler) Query(statement string) (categories.IRow, error) {
	return nil, errors.New("Failed to query database")
}
