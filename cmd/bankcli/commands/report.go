package commands

import (
	"github.com/luistm/go-bank-cli/bank/transactions"
	"github.com/luistm/go-bank-cli/infrastructure/csv"
	"github.com/luistm/go-bank-cli/infrastructure/sqlite"
	"github.com/luistm/go-bank-cli/lib/sellers"
)

// ReportCommand handles reports
type ReportCommand struct {
	commandType string
}

// Execute the report command
func (rc *ReportCommand) Execute(arguments map[string]interface{}) *Response {
	out, err := rc.showReportHandler(arguments["<file>"].(string))
	return &Response{err: err, output: out}
}

func (rc *ReportCommand) showReportHandler(inputFilePath string) (string, error) {

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
	r, err := transactionsInteractor.ReportFromRecords()
	if err != nil {
		return out, err
	}

	return r.String(), nil
}
