package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

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

		for i := 0; i < len(record); i++ {
			fmt.Println(" ", record[i])
		}
		fmt.Println()
		lineCount += 1
	}
}
