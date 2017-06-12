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

func parseValue(value string) float64 {
	parsedValue := strings.Replace(value, ",", ".", -1)
	retValue, _ := strconv.ParseFloat(parsedValue, 64)

	return retValue
}

func parseRecord(record []string) (float64, string) {

	// Columns in the file are:
	// Date of the movement, --, description, expense, credit, --, balance

	for i := 0; i < len(record); i++ {
		// Expense
		if i == 3 && record[i] != "" {
			value := parseValue(record[i])
			expense += value
			fmt.Println("Expense is ", expense)
			return value, record[2]
		}

		// Credit
		if i == 4 && record[i] != "" {
			value := parseValue(record[i])
			credit += value
			fmt.Println("Credit is ", credit)
			return value, record[2]
		}
	}

	return 0, ""
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

	for {
		fmt.Println("Reading line ", lineCount)
		record, error := reader.Read()
		if error == io.EOF {
			break
		}
		if error != nil {
			fmt.Println("Error:", error)
			lineCount++
			continue
		}

		if lineCount < 7 {
			fmt.Println("Ignoring line")
			lineCount++
			continue
		}

		fmt.Println("Record", lineCount, "is", record, "and has", len(record), "fields")
		transactionValue, transactionDescription := parseRecord(record)
		report[transactionDescription] += transactionValue
		lineCount++
	}

	for transactionDescription, transactionValue := range report {
		fmt.Printf("%24s %8.2f \n", transactionDescription, transactionValue)
	}

	// TODO: Fetch data
	// Here, i want this data
	// Initial balance
	// Final Balance
	// Expense per 'description field'
	// Total expense
	// Total credit

}
