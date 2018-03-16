package commands

import (
	"os"

	"github.com/luistm/banksaurus/bank/reportfromrecords"
	"github.com/luistm/banksaurus/cmd/banksaurus/configurations"
	"github.com/luistm/banksaurus/infrastructure/csv"
	"github.com/luistm/banksaurus/infrastructure/sqlite"
	"github.com/luistm/banksaurus/lib/sellers"
	"github.com/luistm/banksaurus/lib/transactions"
)

// Report handles reports
type Report struct{}

// Execute the report command
func (rc *Report) Execute(arguments map[string]interface{}) error {
	var grouped bool

	if arguments["--grouped"].(bool) {
		grouped = true
	}

	CSVStorage, err := csv.New(arguments["<file>"].(string))
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
	if grouped {
		transactionsInteractor := transactions.NewInteractor(
			transactionRepository,
			sellersRepository,
			NewPresenter(os.Stdout),
		)
		err = transactionsInteractor.ReportFromRecordsGroupedBySeller()
	} else {
		rfr, err := reportfromrecords.New(transactionRepository, sellersRepository, NewPresenter(os.Stdout))
		if err != nil {
			return err
		}
		rfr.Execute()
	}

	if err != nil {
		return nil
	}

	return nil
}
