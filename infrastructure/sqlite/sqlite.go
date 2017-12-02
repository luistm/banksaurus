package sqlite

import (
	"database/sql"
	"errors"
	"os"

	"github.com/luistm/go-bank-cli/infrastructure"
	"github.com/luistm/go-bank-cli/lib"

	// To init the database driver
	_ "github.com/mattn/go-sqlite3"
)

var errUndefinedDataBase = errors.New("database is not defined")

// ErrStatementUndefined ...
var ErrStatementUndefined = errors.New("statement is undefined")
var errInvalidConfiguration = errors.New("sqlite configuration parameters are invalid")
var errFailedToCreatedDB = errors.New("failed to create database")

// New creates a new instance of sqlite
func New(path string, name string, memory bool) (infrastructure.SQLStorage, error) {

	if name == "" || path == "" {
		return &sqlite{}, errInvalidConfiguration
	}

	var db *sql.DB
	var err error

	if memory {
		db, err = sql.Open("sqlite3", ":memory:")
	} else {
		if err := validatePath(path); err != nil {
			return &sqlite{}, errInvalidConfiguration
		}
		db, err = sql.Open("sqlite3", path+"/"+name+".db")
	}
	if err != nil {
		return nil, err
	}

	// Create table in order to create the database file
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS categories
	(id INTEGER NOT NULL PRIMARY KEY, name TEXT);
	CREATE TABLE IF NOT EXISTS sellers
	(slug TEXT NOT NULL PRIMARY KEY, name TEXT);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return &sqlite{}, errFailedToCreatedDB
	}

	s := &sqlite{db}
	return s, nil
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

// sqlite contains a sqlite database
type sqlite struct {
	db *sql.DB
}

// Close closes the connection with the sqlite database
func (s *sqlite) Close() error {
	if s.db == nil {
		return errUndefinedDataBase
	}

	return s.db.Close()
}

// Execute is to execute an sql statement
func (s *sqlite) Execute(statement string, values ...interface{}) error {
	if s.db == nil {
		return errUndefinedDataBase
	}

	if statement == "" {
		return ErrStatementUndefined
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(statement, values...)
	// TODO: this section is missing unit tests ----------------------------------
	if err != nil {
		if errTx := tx.Rollback(); errTx != nil {
			return errTx
		}
		return err
	}
	// ---------------------------------------------------------------------------
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// TODO: Make sure rows are being closed across the code

// Query fetches data from the database
func (s *sqlite) Query(statement string, args ...interface{}) (lib.Row, error) {
	if s.db == nil {
		return nil, errUndefinedDataBase
	}

	if statement == "" {
		return nil, ErrStatementUndefined
	}

	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(statement)
	// rows, err := tx.Query(statement) // TODO: why i can't fetch results here? Must read!
	if err != nil {
		if errTx := tx.Rollback(); errTx != nil {
			return nil, errTx
		}
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return rows, nil
}
