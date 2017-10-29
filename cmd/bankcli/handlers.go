package main

import (
	"fmt"

	"github.com/luistm/go-bank-cli/bank/reports"
	"github.com/luistm/go-bank-cli/entities/categories"
	"github.com/luistm/go-bank-cli/entities/descriptions"
	"github.com/luistm/go-bank-cli/infrastructure/csv"
	"github.com/luistm/go-bank-cli/infrastructure/sqlite"
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

func createDescriptionHandler(name string) (string, error) {
	var out string
	SQLStorage, err := sqlite.New(DatabasePath, DatabaseName, false)
	if err != nil {
		return out, err
	}
	defer SQLStorage.Close()

	descriptionsInteractor := descriptions.NewInteractor(SQLStorage)
	d, err := descriptionsInteractor.Create(name)
	if err != nil {
		return out, err
	}

	return d.String(), nil
}

func showDescriptionsHandler() (string, error) {
	var out string
	SQLStorage, err := sqlite.New(DatabasePath, DatabaseName, false)
	if err != nil {
		return out, err
	}
	defer SQLStorage.Close()

	descriptionsInteractor := descriptions.NewInteractor(SQLStorage)
	descriptions, err := descriptionsInteractor.GetAll()
	if err != nil {
		return out, err
	}

	for _, d := range descriptions {
		out += fmt.Sprintf("%s\n", d.String())
	}

	return out, nil
}

// showReportHandler handles report commands
func showReportHandler(inputFilePath string) (string, error) {

	var out string
	CSVStorage, err := csv.New(inputFilePath)
	if err != nil {
		return out, err
	}
	defer CSVStorage.Close()

	reportsInteractor := reports.NewInteractor(CSVStorage)
	_, err = reportsInteractor.CurrentMonth()
	if err != nil {
		return out, err
	}

	return out, nil
}
