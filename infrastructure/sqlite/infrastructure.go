package sqlite

import (
	"database/sql"
	"errors"
	"os"

	"github.com/luistm/banksaurus/infrastructure"
	"github.com/luistm/banksaurus/lib"

	// To init the database driver
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrUndefinedDataBase    = errors.New("database is not defined")
	ErrStatementUndefined   = errors.New("statement is undefined")
	ErrInvalidConfiguration = errors.New("infrastructure configuration parameters are invalid")
)

// ErrFailedToCreatedDB for database error
type ErrFailedToCreatedDB struct {
	Msg string
}

// Error to satisfy the Error interface
func (e *ErrFailedToCreatedDB) Error() string {
	return fmt.Sprintf("failed to create database: %s", e.Msg)
}

// New creates a new instance of Infrastructure
func New(path string, name string, memory bool) (infrastructure.SQLStorage, error) {

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

// Close closes the connection with the Infrastructure database
func (s *Infrastructure) Close() error {
	if s.db == nil {
		return ErrUndefinedDataBase
	}

	return s.db.Close()
}

// Execute is to execute an sql statement
func (s *Infrastructure) Execute(statement string, values ...interface{}) error {
	if s.db == nil {
		return ErrUndefinedDataBase
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
func (s *Infrastructure) Query(statement string, args ...interface{}) (lib.Rows, error) {
	if s.db == nil {
		return nil, ErrUndefinedDataBase
	}

	if statement == "" {
		return nil, ErrStatementUndefined
	}

	// TODO: why i can't fetch results here? Must read!
	//
	// I'm looking into transaction here, because i don't understand why Query does't return results.
	// Must read more stuff about this.

	// tx, err := s.db.Begin()
	// if err != nil {
	// 	return nil, err
	// }

	rows, err := s.db.Query(statement)
	// rows, err := tx.Query(statement)
	if err != nil {
		// if errTx := tx.Rollback(); errTx != nil {
		// 	return nil, errTx
		// }
		return nil, err
	}

	// err = tx.Commit()
	// if err != nil {
	// 	return nil, err
	// }

	return rows, nil
}
