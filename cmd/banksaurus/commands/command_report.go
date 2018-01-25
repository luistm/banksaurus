package commands

import (
	"os"

	"github.com/luistm/banksaurus/bank/transactions"
	"github.com/luistm/banksaurus/cmd/banksaurus/configurations"
	"github.com/luistm/banksaurus/infrastructure/csv"
	"github.com/luistm/banksaurus/infrastructure/sqlite"
	"github.com/luistm/banksaurus/lib/sellers"
)

// Report handles reports
type Report struct{}

// Execute the report command
func (rc *Report) Execute(arguments map[string]interface{}) *Response {
	var grouped bool

	if arguments["--grouped"].(bool) {
		grouped = true
	}

	CSVStorage, err := csv.New(arguments["<file>"].(string))
	if err != nil {
		return &Response{err: err}
	}
	defer CSVStorage.Close()

	dbName, dbPath := configurations.DatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return &Response{err: err}
	}

	transactionRepository := transactions.NewRepository(CSVStorage)
	sellersRepository := sellers.NewRepository(SQLStorage)
	transactionsInteractor := transactions.NewInteractor(
		transactionRepository,
		sellersRepository,
		NewPresenter(os.Stdout),
	)

	if grouped {
		err = transactionsInteractor.ReportFromRecordsGroupedBySeller()
	} else {
		err = transactionsInteractor.ReportFromRecords()
	}

	return &Response{err: err}
}
