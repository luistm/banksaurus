package report

import (
	"os"

	"github.com/luistm/banksaurus/app"
	"github.com/luistm/banksaurus/infrastructure/csv"
	"github.com/luistm/banksaurus/infrastructure/sqlite"
	"github.com/luistm/banksaurus/lib/seller"
	"github.com/luistm/banksaurus/lib/transaction"
	"github.com/luistm/banksaurus/services"
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

	if grouped {
		CSVStorage, err := csv.New(arguments["<file>"].(string))
		if err != nil {
			return err
		}
		defer CSVStorage.Close()

		dbName, dbPath := app.DatabasePath()
		SQLStorage, err := sqlite.New(dbPath, dbName, false)
		if err != nil {
			return err
		}
		defer SQLStorage.Close()

		transactionRepository := transaction.NewRepository(CSVStorage, SQLStorage)
		sellersRepository := seller.NewRepository(SQLStorage)
		presenter := NewPresenter(os.Stdout)

		var rfr services.Servicer

		rfr, err = reportgrouped.New(transactionRepository, sellersRepository, presenter)
		if err != nil {
			return err
		}

		if err := rfr.Execute(); err != nil {
			return nil
		}

	} else {
		// rfr, err = report.New(transactionRepository, sellersRepository, presenter)
		filePath := arguments["<file>"].(string)
		_, err := os.Stat(filePath)
		if err != nil {
			return err
		}

		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

	}

	return nil
}
