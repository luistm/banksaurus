package relational

import (
	"database/sql"
	"errors"

	"github.com/luistm/banksaurus/cmd/bscli/application"
	"os"
)

// NewDatabase creates a new database instance
func NewDatabase() (*sql.DB, error) {

	dbName, dbPath := application.DatabasePath()
	var db *sql.DB

	if err := validatePath(dbPath); err != nil {
		return nil, errors.New("invalid database path")
	}
	db, err := sql.Open("sqlite3", dbPath+"/"+dbName+".db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

// validatePath validates that the database directory or file exists and it is
// in a proper format to be used
func validatePath(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	} else {
		if err != nil {
			return err
		}
		fileInfo, _ = os.Stat(path)
		if !fileInfo.Mode().IsDir() {
			return err
		}
	}

	return nil
}
