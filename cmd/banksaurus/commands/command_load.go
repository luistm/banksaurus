package commands

import (
	"github.com/luistm/banksaurus/bank/usecase/loaddata"
	"github.com/luistm/banksaurus/cmd/banksaurus/configurations"
	"github.com/luistm/banksaurus/infrastructure/csv"
	"github.com/luistm/banksaurus/infrastructure/sqlite"
	"github.com/luistm/banksaurus/lib/seller"
	"github.com/luistm/banksaurus/lib/transaction"
)

// Load command to loaddata input from a file
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
	defer SQLStorage.Close()

	transactionRepository := transaction.NewRepository(CSVStorage)
	sellersRepository := seller.NewRepository(SQLStorage)
	transactionsInteractor := loaddata.New(transactionRepository, sellersRepository)
	err = transactionsInteractor.Execute()
	if err != nil {
		return err
	}

	return nil
}
