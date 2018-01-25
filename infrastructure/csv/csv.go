package csv

import (
	"encoding/csv"
	"os"

	"github.com/luistm/banksaurus/infrastructure"
)

// New opens and returns a file handler for a CSV file
func New(inputFilePath string) (infrastructure.CSVStorage, error) {

	// TODO: Check if path exists and if it is a file
	_, err := os.Stat(inputFilePath)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}

	f := &csvfile{file: file}

	return f, nil
}

// File represents the content of CSV formated file
type csvfile struct {
	file *os.File
}

// Close to disconnect with associated file
func (c *csvfile) Close() error {
	return c.file.Close()
}

// GetAll returns all lines in the file
func (c *csvfile) Lines() ([][]string, error) {
	reader := csv.NewReader(c.file)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1

	lines, err := reader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
