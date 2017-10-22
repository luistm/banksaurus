package main

import (
	"fmt"

	"github.com/luistm/go-bank-cli/bank/reports"
	"github.com/luistm/go-bank-cli/entities/categories"
	"github.com/luistm/go-bank-cli/entities/descriptions"
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

	cr := categories.Repository{SQLStorage: SQLStorage}
	i := categories.Interactor{Repository: &cr}

	cats, err := i.NewCategory(name)
	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("Created category '%s'", cats[0].Name)
	return msg, nil
}

func showCategoryHandler() (string, error) {
	SQLStorage, err := sqlite.New(DatabasePath, DatabaseName, false)
	if err != nil {
		return "", err
	}
	defer SQLStorage.Close()

	cr := categories.Repository{SQLStorage: SQLStorage}
	i := categories.Interactor{Repository: &cr}

	cats, err := i.GetCategories()
	if err != nil {
		return "", err
	}

	out := ""
	for _, c := range cats {
		out += fmt.Sprintf("%s\n", c.Name)
	}

	return out, nil
}

func showDescriptionHandler(name string) (string, error) {
	var out string
	SQLStorage, err := sqlite.New(DatabasePath, DatabaseName, false)
	if err != nil {
		return out, err
	}
	defer SQLStorage.Close()

	descriptionsInteractor := descriptions.NewInteractor(SQLStorage)
	d, err := descriptionsInteractor.Add(name)
	if err != nil {
		return out, err
	}

	return d.String(), nil
}

// showReportHandler handles report commands
func showReportHandler(inputFilePath string) (string, error) {

	err := reports.LoadReport(inputFilePath)
	if err != nil {
		return "", err
	}

	return "", nil
}
