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

	grouped := false
	if arguments["--grouped"].(bool) {
		grouped = true
	}
	out, err := rc.showReport(arguments["<file>"].(string), grouped)
	return &Response{err: err, output: out}
}

func (rc *Report) showReport(inputFilePath string, grouped bool) (string, error) {

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
	transactionsInteractor := transactions.NewInteractor(transactionRepository, sellersRepository)
	var report *transactions.Report
	if grouped {
		report, err = transactionsInteractor.ReportFromRecordsGroupedBySeller()
	} else {
		report, err = transactionsInteractor.ReportFromRecords()
	}
	if err != nil {
		return out, err
	}

	return report.String(), nil
}
