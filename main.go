package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// TODO: Persist the expense types
// TODO: Parse transaction description in order to categorize expenses
// TODO: Want to know the expenses per category in this month
// TODO: produce a report in an excel file
// TODO: Send the excel file by mail
// TODO: Use decimals to represent money
// TODO: The final report should have the initial balance and the end balance

// struct Category{}
// struct Expense{}
// struct Credit{}

var credit float64
var expense float64

func parseValue(value string) float64 {
	parsedValue := strings.Replace(value, ",", ".", -1)
	retValue, _ := strconv.ParseFloat(parsedValue, 64)

	return retValue
}

func parseRecord(record []string) {
	// for i := 0; i < len(record); i++ {
	// 	fmt.Println(" ", record[i])
	// }

	for i := 0; i < len(record); i++ {
		// Expense
		if i == 3 {
			value := parseValue(record[i])
			expense += value
			fmt.Println("Expense is ", expense)
		}

		// Credit
		if i == 4 {
			value := parseValue(record[i])
			credit += value
			fmt.Println("Credit is ", credit)
		}
	}

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
	lineCount := 0

	for {
		record, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Println("Error:", error)
			return
		}

		fmt.Println("Record", lineCount, "is", record, "and has", len(record), "fields")
		parseRecord(record)

		lineCount++
	}
}
