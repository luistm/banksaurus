package load

import (
	"encoding/csv"
	"errors"
	"github.com/luistm/banksaurus/banksauruslib/usecases/loadtransactions"
	"github.com/luistm/banksaurus/cmd/bscli/adapter/transactiongateway"
	"github.com/luistm/banksaurus/cmd/bscli/application"
	"os"
)

// Command command to loadtransactions a csv input from a file
type Command struct{}

// Execute the Command command
func (l *Command) Execute(arguments map[string]interface{}) error {

	// TODO: To much code here, we need a refactor \0/

	filePath := arguments["<file>"].(string)
	_, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	db, err := application.Database()
	if err != nil {
		return err
	}
	defer db.Close()

	sr, err := transactiongateway.NewTransactionRepository(db)
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

	// Execute the interactor
	i, err := loadtransactions.NewInteractor(sr)
	if err != nil {
		return err
	}

	minLinesInWellFormattedFile := 8
	if len(lines) < minLinesInWellFormattedFile {
		return errors.New("Bad formatted file")
	}

	transactionStartLine := 5
	lastTransactionLine := len(lines) - 2
	lines = lines[transactionStartLine:lastTransactionLine]

	r, err := loadtransactions.NewRequest(lines)
	if err != nil {
		return err
	}

	err = i.Execute(r)
	if err != nil {
		return err
	}

	return nil
}
