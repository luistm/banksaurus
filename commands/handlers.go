package commands

import (
	"go-cli-bank/infrastructure"
	"go-cli-bank/lib/categories"
	"go-cli-bank/lib/reports"
)

// CommandCreateCategory handles category creation command
func CreateCategoryHandler(name string) (string, error) {

	dbHandler := infrastructure.DatabaseHandler{}
	cr := categories.CategoryRepository{DBHandler: &dbHandler}
	i := categories.Interactor{Repository: &cr}
	_, err := i.NewCategory(name)
	if err != nil {
		return "", err
	}

	msg := "Created category " + name
	return msg, nil
}

// CommandShowReport handles report commands
func ShowReportHandler(inputFilePath string) (string, error) {
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
