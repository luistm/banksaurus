package application

import (
	"database/sql"
	"errors"
	"os"
)

const (
	createSelleTableStatement = `
CREATE TABLE IF NOT EXISTS seller (
  slug TEXT NOT NULL,
  name TEXT,
  UNIQUE (slug)
);`

	createTransactionTableStatement = `
CREATE TABLE IF NOT EXISTS "transaction"(
  id INTEGER PRIMARY KEY,
  seller BIGINT NOT NULL,
  amount BIGINT NOT NULL,
  FOREIGN KEY (seller) REFERENCES seller(slug)
);`

	dbName = "bank"
)

// Database creates a new database instance
func Database() (*sql.DB, error) {

	if err := validatePath(Path()); err != nil {
		return nil, errors.New("invalid database path")
	}

	db, err := sql.Open("sqlite3", Path()+"/"+dbName+".db")
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

// buildSchema runs the create tables statements
func buildSchema(db *sql.DB) error {

	_, err := db.Exec(createSelleTableStatement)
	if err != nil {
		return err
	}

	_, err = db.Exec(createTransactionTableStatement)
	if err != nil {
		return err
	}

	return nil
}
