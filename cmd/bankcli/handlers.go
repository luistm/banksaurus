package main

import (
	"fmt"

	"github.com/luistm/go-bank-cli/bank/reports"
	"github.com/luistm/go-bank-cli/entities/categories"
	"github.com/luistm/go-bank-cli/infrastructure/sqlite"
)

var DatabaseName = "bank.db"
var DatabasePath = "/tmp"

// createCategoryHandler handles category creation command
func createCategoryHandler(name string) (string, error) {

	storage, err := sqlite.New(DatabasePath, DatabaseName, false)
	if err != nil {
		return "", err
	}
	defer storage.Close()

	cr := categories.CategoryRepository{DBHandler: storage}
	i := categories.Interactor{Repository: &cr}

	cats, err := i.NewCategory(name)
	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("Created category '%s'", cats[0].Name)
	return msg, nil
}

func showCategoryHandler() (string, error) {
	storage, err := sqlite.New(DatabasePath, DatabaseName, false)
	if err != nil {
		return "", err
	}
	defer storage.Close()

	cr := categories.CategoryRepository{DBHandler: storage}
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

// showReportHandler handles report commands
func showReportHandler(inputFilePath string) (string, error) {

	err := reports.LoadReport(inputFilePath)
	if err != nil {
		return "", err
	}

	return "", nil
}
