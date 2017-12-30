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
	out, err := l.loadFile(arguments["<file>"].(string))
	return &Response{err: err, output: out}
}

func (l *Load) loadFile(inputFilePath string) (string, error) {
	var out string
	CSVStorage, err := csv.New(inputFilePath)
	if err != nil {
		return out, err
	}
	defer CSVStorage.Close()

	dbName, dbPath := configurations.GetDatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return out, err
	}

	transactionRepository := transactions.NewRepository(CSVStorage)
	sellersRepository := sellers.NewRepository(SQLStorage)
	transactionsInteractor := transactions.NewInteractor(transactionRepository, sellersRepository, nil)
	err = transactionsInteractor.LoadDataFromRecords()
	if err != nil {
		return out, err
	}

	return "", nil
}
