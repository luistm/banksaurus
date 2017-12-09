package commands

import (
	"github.com/luistm/go-bank-cli/bank/transactions"
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

	SQLStorage, err := sqlite.New(DatabasePath, DatabaseName, false)
	if err != nil {
		return out, err
	}

	transactionRepository := transactions.NewRepository(CSVStorage)
	sellersRepository := sellers.NewRepository(SQLStorage)
	transactionsInteractor := transactions.NewInteractor(transactionRepository, sellersRepository)
	err = transactionsInteractor.LoadDataFromRecords()
	if err != nil {
		return out, err
	}

	return "", nil
}
