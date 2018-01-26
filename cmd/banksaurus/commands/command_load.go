package commands

import (
	"github.com/luistm/banksaurus/bank/transactions"
	"github.com/luistm/banksaurus/cmd/banksaurus/configurations"
	"github.com/luistm/banksaurus/infrastructure/csv"
	"github.com/luistm/banksaurus/infrastructure/sqlite"
	"github.com/luistm/banksaurus/lib/sellers"
)

// Load command to load input from a file
type Load struct{}

// Execute the Load command
func (l *Load) Execute(arguments map[string]interface{}) error {
	err := l.loadFile(arguments["<file>"].(string))
	if err != nil {
		return nil
	}

	return nil
}

func (l *Load) loadFile(inputFilePath string) error {
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

	transactionRepository := transactions.NewRepository(CSVStorage)
	sellersRepository := sellers.NewRepository(SQLStorage)
	transactionsInteractor := transactions.NewInteractor(transactionRepository, sellersRepository, nil)
	err = transactionsInteractor.LoadDataFromRecords()
	if err != nil {
		return err
	}

	return nil
}
