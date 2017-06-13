package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var credit float64
var expense float64

// Record represents a single line the the file
type Record struct {
	record []string
}

// CREDIT ...
var CREDIT = "Credit"

// DEBT ..
var DEBT = "Debt"

// Transaction is a money movement
type Transaction struct {
	field           string
	description     string
	transactionType string
}

func (t *Transaction) new(record Record) *Transaction {

	for i := 0; i < len(record.record); i++ {
		t.field = record.record[i]
		t.description = record.record[2]

		// Expense
		if i == 3 && record.record[i] != "" {
			value := t.value()
			t.transactionType = DEBT
			expense += value
			return t
		}

		// Credit
		if i == 4 && record.record[i] != "" {
			t.transactionType = CREDIT
			value := t.value()
			credit += value
			return t
		}
	}

	return t
}

func (t *Transaction) value() float64 {
	parsedField := strings.Replace(t.field, ".", "", -1)
	parsedField = strings.Replace(parsedField, ",", ".", -1)
	retValue, _ := strconv.ParseFloat(parsedField, 64)

	return retValue
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
	// TODO: Check if the initial balance matches the one comming in the file
	// TODO: Read the records and save each one to the SQlite

	for {
		r, error := reader.Read()
		record := Record{r}
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

		if len(record.record) != 8 {
			lineCount++
			continue
		}

		t := Transaction{}
		transaction := t.new(record)
		report[transaction.description] += transaction.value()
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
