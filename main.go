package main

import (
	"encoding/csv"
	"expensetracker/entities"
	"fmt"
	"io"
	"os"
)

var credit float64
var expense float64

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
	// TODO: Check if the initial balance matches the one comming in the file
	// TODO: Read the records and save each one to the SQlite

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
