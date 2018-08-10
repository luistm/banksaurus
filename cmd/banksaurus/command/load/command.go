package load

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"github.com/luistm/banksaurus/app"
	"github.com/luistm/banksaurus/next/adapter/CGDcsv"
	"github.com/luistm/banksaurus/next/adapter/sqlite"
	"github.com/luistm/banksaurus/next/load"
	"os"
)

// Command command to load a csv input from a file
type Command struct{}

// Execute the Command command
func (l *Command) Execute(arguments map[string]interface{}) error {

	// TODO: To much code here

	filePath := arguments["<file>"].(string)
	_, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	// Create repository to access csv
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1

	lines, err := reader.ReadAll()
	if err != nil {
		return err
	}

	tr, err := CGDcsv.New(lines)
	if err != nil {
		return err
	}

	// Create database repository
	dbName, dbPath := app.DatabasePath()
	var db *sql.DB

	if err := validatePath(dbPath); err != nil {
		return errors.New("invalid database path")
	}
	db, err = sql.Open("sqlite3", dbPath+"/"+dbName+".db")
	if err != nil {
		return err
	}

	sr, err := sqlite.NewSellerRepository(db)
	if err != nil {
		return err
	}

	// Execute the interactor
	i, err := load.NewInteractor(tr, sr)
	if err != nil {
		return err
	}

	err = i.Execute()
	if err != nil {
		return err
	}

	return nil
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
