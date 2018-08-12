package load

import (
	"encoding/csv"
	"github.com/luistm/banksaurus/next/application/adapter/cgdcsv"
	"github.com/luistm/banksaurus/next/application/adapter/sqlite"
	"github.com/luistm/banksaurus/next/application/infrastructure/relational"
	"github.com/luistm/banksaurus/next/loadtransactions"
	"os"
)

// Command command to loadtransactions a csv input from a file
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

	tr, err := cgdcsv.New(lines)
	if err != nil {
		return err
	}

	db, err := relational.NewDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	sr, err := sqlite.NewSellerRepository(db)
	if err != nil {
		return err
	}

	// Execute the interactor
	i, err := loadtransactions.NewInteractor(tr, sr)
	if err != nil {
		return err
	}

	err = i.Execute()
	if err != nil {
		return err
	}

	return nil
}
