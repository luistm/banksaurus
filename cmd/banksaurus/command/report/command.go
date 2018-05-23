package report

import (
	"os"

	"github.com/luistm/banksaurus/cmd/banksaurus/configurations"
	"github.com/luistm/banksaurus/infrastructure/csv"
	"github.com/luistm/banksaurus/infrastructure/sqlite"
	"github.com/luistm/banksaurus/lib/seller"
	"github.com/luistm/banksaurus/lib/transaction"
	"github.com/luistm/banksaurus/services"
	"github.com/luistm/banksaurus/services/report"
	"github.com/luistm/banksaurus/services/reportgrouped"
)

// Command handles reports
type Command struct{}

// Execute the report command
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

	var rfr services.Servicer
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
