package main

import (
	"fmt"
	"go-bank-cli/infrastructure"
	"go-bank-cli/lib/categories"
	"go-bank-cli/lib/reports"
)

var DatabaseName = "bank.db"
var DatabasePath = "/tmp"

// createCategoryHandler handles category creation command
func createCategoryHandler(name string) (string, error) {

	// TODO: Get database values from a configuration
	db, err := infrastructure.ConnectDB(DatabaseName, DatabasePath)
	if err != nil {
		return "", err
	}
	defer db.Close()

	dbHandler := infrastructure.DatabaseHandler{Database: db}
	cr := categories.CategoryRepository{DBHandler: &dbHandler}
	i := categories.Interactor{Repository: &cr}
	_, err = i.NewCategory(name)
	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("Created category '%s'", name)
	return msg, nil
}

// showReportHandler handles report commands
func showReportHandler(inputFilePath string) (string, error) {
	file, err := infrastructure.OpenFile(inputFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	err = reports.MonthlyReport(file)
	if err != nil {
		return "", err
	}

	return "", nil
}
