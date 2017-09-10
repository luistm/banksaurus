package main

// This software is an expense tracker i made to read the transactions exported from
// my bank account.

import (
	"flag"
	"fmt"
	"go-cli-bank/accounts"
	"go-cli-bank/categories"
	"go-cli-bank/infrastructure"
	"go-cli-bank/reports"
	"os"
	// flag "github.com/ogier/pflag"
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

func errorf(format string, args ...interface{}) {
	fmt.Fprintln(os.Stderr, format, args)
	os.Exit(2)
}

func main() {

	inputFilePath := flag.String("load", "", "Specify the path to the input file")
	showReport := flag.Bool("report", false, "Show report")
	showBalance := flag.Bool("balance", false, "Show current balance")
	createCategory := flag.String("category", "", "Create category")
	flag.Parse()

	if *createCategory != "" {
		if err := CommandCreateCategory(*createCategory); err != nil {
			errorf("Failed to create category: %v\n", err)
		}
	}

	if *showReport {
		if err := CommandShowReport(*inputFilePath); err != nil {
			errorf("Failed to show report: %v\n", err)
		}
	}

	if *showBalance {
		fmt.Println(accounts.CurrentBalance().String())
	}
}

// var DATABASE_NAME string = "./go-cli-bank.db"
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
