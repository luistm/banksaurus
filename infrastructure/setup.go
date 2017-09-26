package infrastructure

import (
	"database/sql"
	"fmt"
	"os"

	// To init the database driver
	_ "github.com/mattn/go-sqlite3"
)

var errMessage = "Storage initialization failed: "

// ErrInitFailed signals the the init of the database failed
type ErrInitFailed struct {
	arg string
}

func (e *ErrInitFailed) Error() string {
	return fmt.Sprintf(errMessage+"%s", e.arg)
}

// I don't need this stuff.. just return the error....
var messageDBNameEmpty = "Database name is empty"
var messageDBPathEmpty = "Database path is empty"
var messageDBPathInvalid = "Database path is invalid"

func validatePath(path string) error {
	// Validate that database directory or file exists and it is
	// in a proper format to be used
	fileInfo, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return &ErrInitFailed{arg: err.Error()}
		}
		fmt.Println("Directory created")
	} else {
		if err != nil {
			return &ErrInitFailed{arg: err.Error()}
		}
		fileInfo, _ = os.Stat(path)
		if !fileInfo.Mode().IsDir() {
			return &ErrInitFailed{arg: "Path to database is a file."}
		}
	}

	return nil
}

// InitStorage configures the storage to persist data.
// Receives:
// a) The name of the database
// b) The path where the database file should be
//
// Returns the path of the database file
func InitStorage(dbName string, path string) (*sql.DB, error) {

	if dbName == "" {
		return nil, &ErrInitFailed{arg: messageDBNameEmpty}
	}

	if path == "" {
		return nil, &ErrInitFailed{arg: messageDBPathEmpty}
	}

	if err := validatePath(path); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", path+"/"+dbName+".db")
	if err != nil {
		return nil, &ErrInitFailed{arg: err.Error()}
	}

	// Create table in order to create the database file
	sqlStmt := `
	create table categories (id integer not null primary key, name text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, &ErrInitFailed{arg: err.Error()}
	}

	// TODO: Don't forget to close the database connection

	return db, nil
}
