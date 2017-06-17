package main

// This software is an expense tracker i made to read the transactions exported from
// my bank account.
// TODO: Save a report for the current data in an excel format
// TODO: Save credit/debt to an excel file if --excel option provided
// TODO: Save each transaction to an sql database
// TODO: Use shopspring decimals to represent money
// TODO: Refactor, because this code is crap
// TODO: Add some kind of GUI to see transactions. Maybe i will use eletron.. i dunno
// TODO: Add a good readme: How to Build? How to run? How to contribute?
// TODO: Have fun!!

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"expensetracker/interactor"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shopspring/decimal"
	"github.com/tealeg/xlsx"
)

var DATABASE_NAME string = "./expensetracker.db"
var DATABASE_ENGINE = "sqlite3"

func toExcel(value decimal.Decimal, description string) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = description
	cell = row.AddCell()
	// cell.Value = strconv.FormatFloat(value, 'f', 2, 64)
	cell.Value = value.String()

	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

// Database is the holder for database operations
type Database struct {
	db *sql.DB
}

// NewDBConnection provides a new connection to the database
func (d *Database) NewDBConnection() error {
	log.Println("Creating new database connection")

	if _, err := os.Stat(DATABASE_NAME); os.IsNotExist(err) && d.db == nil {
		// os.Remove(DATABASE_NAME)
		db, err := sql.Open(DATABASE_ENGINE, DATABASE_NAME)
		if err != nil {
			return err
		}

		d.db = db

		return nil
	}

	if _, err := os.Stat(DATABASE_NAME); os.IsNotExist(err) && d.db != nil {
		log.Fatal("Connection exists, but database file does not... wtf??")
	}

	if d.db == nil {
		os.Remove(DATABASE_NAME)
		db, err := sql.Open(DATABASE_ENGINE, DATABASE_NAME)
		if err != nil {
			return err
		}

		d.db = db

		return nil
	}

	return errors.New("Database connection already exists")
}

// CreateExpenseDatabase creates the expense database
func (d *Database) CreateExpenseDatabase() error {
	log.Println("Creating new database")

	sqlStmt := `
	create table expenses (id integer not null primary key, description text, value float);
	create table credits (id integer not null primary key, description text, value float);
	delete from expenses;
	delete from expenses;
	`
	d.db.Ping()
	_, err := d.db.Exec(sqlStmt)
	if err != nil {
		// log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}

	return nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) SaveExpense(value decimal.Decimal, description string) {
	_, err := d.db.Exec("insert into expenses(value, description) values(?, ?)", value, description)
	if err != nil {
		log.Fatal("Error while saving expense: ", err)
	}
}

// func (d *Database) SaveCredit(value float64, description string) {
// 	tx, err := d.db.Begin()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	stmt, err := tx.Prepare("insert into credits(id, name) values(?, ?)")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()
// 	_, err = stmt.Exec()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	tx.Commit()
// }

// documentation for csv is at http://golang.org/pkg/encoding/csv/
func main() {

	file, error := os.Open("comprovativo.csv")
	if error != nil {
		fmt.Println("Error:", error)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1 // If FieldsPerRecord is negative, no check is made and records may have a variable number of fields.

	database := Database{}
	if err := database.NewDBConnection(); err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	if err := database.CreateExpenseDatabase(); err != nil {
		log.Fatal(err)
	}

	database.SaveExpense(decimal.NewFromFloat(1), "descricao")

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Failed to read CSV file ", err)
	}
	interactor.MonthlyReport(records)

}
