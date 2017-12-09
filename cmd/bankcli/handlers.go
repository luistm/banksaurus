package main

import (
	"fmt"

	"github.com/luistm/go-bank-cli/bank/transactions"
	"github.com/luistm/go-bank-cli/infrastructure/csv"
	"github.com/luistm/go-bank-cli/infrastructure/sqlite"
	"github.com/luistm/go-bank-cli/lib/categories"
	"github.com/luistm/go-bank-cli/lib/sellers"
)

var DatabaseName = "bank.db"
var DatabasePath = "/tmp"

// createCategoryHandler handles category creation command
func createCategoryHandler(name string) (string, error) {

	SQLStorage, err := sqlite.New(DatabasePath, DatabaseName, false)
	if err != nil {
		return "", err
	}
	defer SQLStorage.Close()

	categoriesInteractor := categories.NewInteractor(SQLStorage)
	cats, err := categoriesInteractor.Create(name)
	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("Created category '%s'", cats[0])
	return msg, nil
}

func showCategoriesHandler() (string, error) {
	SQLStorage, err := sqlite.New(DatabasePath, DatabaseName, false)
	if err != nil {
		return "", err
	}
	defer SQLStorage.Close()

	categoriesInteractor := categories.NewInteractor(SQLStorage)
	cats, err := categoriesInteractor.GetAll()
	if err != nil {
		return "", err
	}

	out := ""
	for _, c := range cats {
		out += fmt.Sprintf("%s\n", c)
	}

	return out, nil
}

func createSellerHandler(name string) (string, error) {
	var out string
	SQLStorage, err := sqlite.New(DatabasePath, DatabaseName, false)
	if err != nil {
		return out, err
	}
	defer SQLStorage.Close()

	sellersInteractor := sellers.NewInteractor(SQLStorage)
	s, err := sellersInteractor.Create(name)
	if err != nil {
		return out, err
	}

	return s.String(), nil
}

func showSellersHandler() (string, error) {
	var out string
	SQLStorage, err := sqlite.New(DatabasePath, DatabaseName, false)
	if err != nil {
		return out, err
	}
	defer SQLStorage.Close()

	sellersInteractor := sellers.NewInteractor(SQLStorage)
	sellers, err := sellersInteractor.GetAll()
	if err != nil {
		return out, err
	}

	for _, s := range sellers {
		out += fmt.Sprintf("%s\n", s)
	}

	return out, nil
}

func sellerChangePrettyName(sellerID string, name string) (string, error) {
	var out string
	SQLStorage, err := sqlite.New(DatabasePath, DatabaseName, false)
	if err != nil {
		return out, err
	}
	defer SQLStorage.Close()

	sellersInteractor := sellers.NewInteractor(SQLStorage)
	err = sellersInteractor.Update(sellerID, name)
	if err != nil {
		return out, err
	}

	return out, nil
}

func loadHandler(inputFilePath string) (string, error) {
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
	err = transactionsInteractor.LoadDataFromRecords()
	if err != nil {
		return out, err
	}

	return "", nil
}

// ReportCommand handles reports
type ReportCommand struct {
	commandType string
}

func (rc *ReportCommand) execute(arguments map[string]interface{}) *Response {
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
