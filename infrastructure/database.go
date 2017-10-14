package infrastructure

import (
	"database/sql"
	"errors"
	"github.com/luistm/go-bank-cli/lib/categories"
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
	Database *sql.DB
}

// Execute is to execute an sql statement
func (dh *DatabaseHandler) Execute(statement string, values ...interface{}) error {
	if dh.Database == nil {
		return ErrDataBaseConnUndefined
	}

	tx, _ := dh.Database.Begin()
	_, err := tx.Exec(statement, values...)
	if err != nil {
		// TODO: tx.Rollback
		return &ErrDataBase{err.Error()}
	}
	tx.Commit()

	return nil
}

// Query fetches data from the database
func (dh *DatabaseHandler) Query(statement string) (categories.IRow, error) {
	if dh.Database == nil {
		return nil, ErrDataBaseConnUndefined
	}

	rows, err := dh.Database.Query(statement)
	if err != nil {
		return nil, &ErrDataBase{err.Error()}
	}

	return rows, nil
}
