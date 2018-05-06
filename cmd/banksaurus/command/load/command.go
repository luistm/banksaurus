package load

import (
	"github.com/luistm/banksaurus/bankservices/loaddata"
	"github.com/luistm/banksaurus/cmd/banksaurus/configurations"
	"github.com/luistm/banksaurus/infrastructure/csv"
	"github.com/luistm/banksaurus/infrastructure/sqlite"
)

// Command command to loaddata input from a file
type Command struct{}

// Execute the Command command
func (l *Command) Execute(arguments map[string]interface{}) error {
	err := l.loadFile(arguments["<file>"].(string))
	if err != nil {
		return nil
	}

	return nil
}

func (l *Command) loadFile(inputFilePath string) error {
	CSVStorage, err := csv.New(inputFilePath)
	if err != nil {
		return err
	}
	defer CSVStorage.Close()

	dbName, dbPath := configurations.DatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return err
	}
	defer SQLStorage.Close()

	transactionsInteractor := loaddata.New(CSVStorage, SQLStorage)
	err = transactionsInteractor.Execute()
	if err != nil {
		return err
	}

	return nil
}