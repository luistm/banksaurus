package main

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"expensetracker/entities"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var credit float64
var expense float64

var DATABASE_NAME string = "./expensetracker.db"
var DATABASE_ENGINE = "sqlite3"

// Database is the holder for database operations
type Database struct {
	db *sql.DB
}

// NewDBConnection provides a new connection to the database
func (d *Database) NewDBConnection() error {
	log.Println("Creating new database connection")

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
	create table foo (id integer not null primary key, name text);
	delete from foo;
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
	lineCount := 0

	var report map[string]float64
	report = make(map[string]float64)

	// TODO: Open SQlite, read the initial balance.
	database := Database{}
	if err := database.NewDBConnection(); err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	if err := database.CreateExpenseDatabase(); err != nil {
		log.Fatal(err)
	}

	for {
		r, error := reader.Read()
		record := entities.Record{Record: r}
		if error == io.EOF {
			break
		}
		if error != nil {
			fmt.Println("Error:", error)
			lineCount++
			continue
		}

		if lineCount < 4 {
			lineCount++
			continue
		}

		if len(record.Record) != 8 {
			lineCount++
			continue
		}

		t := entities.Transaction{}
		transaction := t.New(record)
		report[transaction.Description] += transaction.Value()
		if transaction.TransactionType == entities.DEBT {
			expense += transaction.Value()
		} else {
			credit += transaction.Value()
		}
		lineCount++
	}

	for transactionDescription, transactionValue := range report {
		fmt.Printf("%24s %8.2f \n", transactionDescription, transactionValue)
	}

	fmt.Println("Expense is ", expense)
	fmt.Println("Credit is ", credit)

	// TODO: Fetch data
	// Here, i want this data
	// Initial balance
	// Final Balance
	// Expense per 'description field'
	// Total expense
	// Total credit

}
