package report

import (
	"encoding/csv"
	"os"

	"github.com/luistm/banksaurus/app"
	infraCSV "github.com/luistm/banksaurus/infrastructure/csv"
	"github.com/luistm/banksaurus/infrastructure/sqlite"
	"github.com/luistm/banksaurus/lib/seller"
	"github.com/luistm/banksaurus/lib/transaction"
	"github.com/luistm/banksaurus/next/adapter/CGDcsv"
	"github.com/luistm/banksaurus/next/adapter/transactionpresenter"
	"github.com/luistm/banksaurus/next/report"
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
		CSVStorage, err := infraCSV.New(arguments["<file>"].(string))
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

		reader := csv.NewReader(file)
		reader.Comma = ';'
		reader.FieldsPerRecord = -1

		lines, err := reader.ReadAll()
		if err != nil {
			return err
		}

		inputGateway, err := CGDcsv.New(lines)
		if err != nil {
			return err
		}

		p, err := transactionpresenter.NewPresenter()
		if err != nil {
			return err
		}

		i, err := report.NewInteractor(p, inputGateway)
		if err != nil {
			return err
		}

		r, _ := report.NewRequest()
		err = i.Execute(r)
		if err != nil {
			return err
		}

		vm, err := p.ViewModel()
		if err != nil {
			return err
		}

		vm.Write(os.Stdout)
	}

	return nil
}
