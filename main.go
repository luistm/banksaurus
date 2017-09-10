package main

// This software is an expense tracker i made to read the transactions exported from
// my bank account.

import (
	"expensetracker/accounts"
	"expensetracker/categories"
	"expensetracker/infrastructure"
	"expensetracker/reports"
	"fmt"
	"log"

	flag "github.com/ogier/pflag"
)

// CommandCreateCategory handles category creation command
func CommandCreateCategory(name string) error {

	dbHandler := infrastructure.DatabaseHandler{}
	cr := categories.CategoryRepository{DBHandler: &dbHandler}
	i := categories.Interactor{Repository: &cr}
	_, err := i.NewCategory(name)
	if err != nil {
		return err
	}

	return nil
}

// CommandShowReport handles report commands
func CommandShowReport(inputFilePath string) error {
	file, err := infrastructure.OpenFile(inputFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = reports.MonthlyReport(file)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	var inputFilePath string
	flag.StringVarP(&inputFilePath, "load", "l", "", "Specify the path to the input file")

	var showReport bool
	flag.BoolVarP(&showReport, "report", "r", false, "Show report")

	var showBalance bool
	flag.BoolVarP(&showBalance, "balance", "b", false, "Show current balance")

	var createCategory string
	flag.StringVarP(&createCategory, "category", "c", "", "Create category")

	flag.Parse()

	if createCategory != "" {
		if err := CommandCreateCategory(createCategory); err != nil {
			log.Fatalf("Failed to create category: %s", err)
		}
	}

	if showReport {
		if err := CommandShowReport(inputFilePath); err != nil {
			log.Fatalf("Failed to show report %s", err)
		}
	}

	if showBalance {
		fmt.Println(accounts.CurrentBalance().String())
	}
}

// var DATABASE_NAME string = "./expensetracker.db"
// var DATABASE_ENGINE = "sqlite3"

// func toExcel(value decimal.Decimal, description string) {
// 	var file *xlsx.File
// 	var sheet *xlsx.Sheet
// 	var row *xlsx.Row
// 	var cell *xlsx.Cell
// 	var err error

// 	file = xlsx.NewFile()
// 	sheet, err = file.AddSheet("Sheet1")
// 	if err != nil {
// 		fmt.Printf(err.Error())
// 	}
// 	row = sheet.AddRow()
// 	cell = row.AddCell()
// 	cell.Value = description
// 	cell = row.AddCell()
// 	// cell.Value = strconv.FormatFloat(value, 'f', 2, 64)
// 	cell.Value = value.String()

// 	err = file.Save("MyXLSXFile.xlsx")
// 	if err != nil {
// 		fmt.Printf(err.Error())
// 	}
// }
