package commands

import (
	"github.com/luistm/go-bank-cli/bank/transactions"
	"github.com/luistm/go-bank-cli/cmd/bankcli/configurations"
	"github.com/luistm/go-bank-cli/infrastructure/csv"
	"github.com/luistm/go-bank-cli/infrastructure/sqlite"
	"github.com/luistm/go-bank-cli/lib/sellers"
)

// Load command to load input from a file
type Load struct{}

// Execute the Load command
func (l *Load) Execute(arguments map[string]interface{}) *Response {
	err := l.loadFile(arguments["<file>"].(string))
	return &Response{err: err}
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
