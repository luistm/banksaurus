package commands

import (
	"github.com/luistm/go-bank-cli/bank/transactions"
	"github.com/luistm/go-bank-cli/cmd/bankcli/configurations"
	"github.com/luistm/go-bank-cli/infrastructure/csv"
	"github.com/luistm/go-bank-cli/infrastructure/sqlite"
	"github.com/luistm/go-bank-cli/lib/sellers"
)

// Report handles reports
type Report struct{}

// Execute the report command
func (rc *Report) Execute(arguments map[string]interface{}) *Response {
	var out string
	var grouped bool

	if arguments["--grouped"].(bool) {
		grouped = true
	}

	CSVStorage, err := csv.New(arguments["<file>"].(string))
	if err != nil {
		return &Response{err: err, output: out}
	}
	defer CSVStorage.Close()

	dbName, dbPath := configurations.GetDatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return &Response{err: err, output: out}
	}

	transactionRepository := transactions.NewRepository(CSVStorage)
	sellersRepository := sellers.NewRepository(SQLStorage)

	transactionsInteractor := transactions.NewInteractor(transactionRepository, sellersRepository, &CLIPresenter{})
	if grouped {
		err = transactionsInteractor.ReportFromRecordsGroupedBySeller()
	} else {
		err = transactionsInteractor.ReportFromRecords()
	}

	return &Response{err: err, output: out}
}
