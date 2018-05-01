package commands

import (
	"os"

	"github.com/luistm/banksaurus/bank"
	"github.com/luistm/banksaurus/bank/reportfromrecords"
	"github.com/luistm/banksaurus/bank/reportfromrecordsgrouped"
	"github.com/luistm/banksaurus/cmd/banksaurus/configurations"
	"github.com/luistm/banksaurus/infrastructure/csv"
	"github.com/luistm/banksaurus/infrastructure/sqlite"
	"github.com/luistm/banksaurus/lib/sellers"
	"github.com/luistm/banksaurus/lib/transaction"
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
	defer SQLStorage.Close()

	transactionRepository := transaction.NewRepository(CSVStorage)
	sellersRepository := sellers.NewRepository(SQLStorage)
	presenter := NewPresenter(os.Stdout)

	var rfr bank.Interactor
	if grouped {
		rfr, err = reportfromrecordsgrouped.New(transactionRepository, sellersRepository, presenter)
	} else {
		rfr, err = reportfromrecords.New(transactionRepository, sellersRepository, presenter)
	}
	if err != nil {
		return err
	}

	if err := rfr.Execute(); err != nil {
		return nil
	}

	return nil
}
