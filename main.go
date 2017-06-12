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

func parseValue(value string) float64 {
	parsedValue := strings.Replace(value, ".", "", -1)
	parsedValue = strings.Replace(parsedValue, ",", ".", -1)
	retValue, _ := strconv.ParseFloat(parsedValue, 64)

	return retValue
}

func parseRecord(record Record) (float64, string) {

	// Columns in the file are:
	// Date of the movement, --, description, expense, credit, --, balance

	for i := 0; i < len(record.record); i++ {
		// Expense
		if i == 3 && record.record[i] != "" {
			value := parseValue(record.record[i])
			expense += value
			return value, record.record[2]
		}

		// Credit
		if i == 4 && record.record[i] != "" {
			value := parseValue(record.record[i])
			credit += value
			return value, record.record[2]
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

		transactionValue, transactionDescription := parseRecord(record)
		report[transactionDescription] += transactionValue
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
