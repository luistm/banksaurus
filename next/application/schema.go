package application

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var ErrInvalidConfiguration = errors.New("infrastructure configuration parameters are invalid")

// ErrFailedToCreatedDB for database error
type ErrFailedToCreatedDB struct {
	Msg string
}

// Error to satisfy the Error interface
func (e *ErrFailedToCreatedDB) Error() string {
	return fmt.Sprintf("failed to create database: %s", e.Msg)
}

// New creates a new instance of Infrastructure
func NewSchema(path string, name string, memory bool) (*Infrastructure, error) {

	if name == "" || path == "" {
		return &Infrastructure{}, ErrInvalidConfiguration
	}

	var db *sql.DB
	var err error

	if memory {
		db, err = sql.Open("sqlite3", ":memory:")
	} else {
		if err := validatePath(path); err != nil {
			return &Infrastructure{}, ErrInvalidConfiguration
		}
		db, err = sql.Open("sqlite3", path+"/"+name+".db")
	}
	if err != nil {
		return nil, err
	}

	// TODO: Move table creation to a proper place ---------------------------------------------------------------------
	// Create table in order to create the database file
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS seller
	(
		slug TEXT NOT NULL PRIMARY KEY,
		name TEXT
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return &Infrastructure{}, &ErrFailedToCreatedDB{Msg: err.Error()}
	}

	sqlStmt = `
	CREATE TABLE IF NOT EXISTS transaction_types
	(
		TYPE CHAR NOT NULL
	)
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return &Infrastructure{}, &ErrFailedToCreatedDB{Msg: err.Error()}
	}

	// Create transactions table in order to create the database file
	// 25-10-2017;25-10-2017;COMPRA CONTINENTE MAI ;77,52;;61,25;61,25;
	// id, seller_id, debt amount, credit amount, contabilistic, real
	sqlStmt = `
	CREATE TABLE IF NOT EXISTS transactions
	(
		ID int NOT NULL PRIMARY KEY,
		SELLER_ID int NOT NULL,
		AMOUNT int DEFAULT 0,
		TYPE 
		BALANCE int NOT NULL
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return &Infrastructure{}, &ErrFailedToCreatedDB{Msg: err.Error()}
	}

	// -----------------------------------------------------------------------------------------------------------------

	s := &Infrastructure{db}
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

// Infrastructure contains a Infrastructure database
type Infrastructure struct {
	db *sql.DB
}
