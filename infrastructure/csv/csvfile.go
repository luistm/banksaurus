package csv

import (
	"encoding/csv"
	"os"
)

// OpenFile opens and returns a file handler for a CSV file
func OpenFile(inputFilePath string) ([][]string, error) {

	// TODO: Check if path exists and if it is a file
	_, err := os.Stat(inputFilePath)
	if err != nil {
		return [][]string{}, err
	}

	file, err := os.Open(inputFilePath)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, err
}
