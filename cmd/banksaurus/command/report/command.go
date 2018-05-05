package report

import (
	"os"

	"github.com/luistm/banksaurus/banklib/seller"
	"github.com/luistm/banksaurus/banklib/transaction"
	"github.com/luistm/banksaurus/bankservices"
	"github.com/luistm/banksaurus/bankservices/report"
	"github.com/luistm/banksaurus/bankservices/reportgrouped"
	"github.com/luistm/banksaurus/cmd/banksaurus/configurations"
	"github.com/luistm/banksaurus/infrastructure/csv"
	"github.com/luistm/banksaurus/infrastructure/sqlite"
)

// Command handles reports
type Command struct{}

// Execute the reportgrouped command
func (rc *Command) Execute(arguments map[string]interface{}) error {
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

	transactionRepository := transaction.NewRepository(CSVStorage, SQLStorage)
	sellersRepository := seller.NewRepository(SQLStorage)
	presenter := NewPresenter(os.Stdout)

	var rfr bankservices.Servicer
	if grouped {
		rfr, err = reportgrouped.New(transactionRepository, sellersRepository, presenter)
	} else {
		rfr, err = report.New(transactionRepository, sellersRepository, presenter)
	}
	if err != nil {
		return err
	}

	if err := rfr.Execute(); err != nil {
		return nil
	}

	return nil
}
