package cgd_csv

import (
	"encoding/csv"
	"errors"
	"github.com/luistm/banksaurus/next/entity/transaction"
	"os"
	"strconv"
	"strings"
	"time"
)

// New opens and returns a file handler for a CSV file
func New(inputFilePath string) (*Repository, error) {
	_, err := os.Stat(inputFilePath)
	if err != nil {
		return nil, err
	}

	f := &Repository{filePath: inputFilePath}

	return f, nil
}

// Repository represents the content of CSV formatted file
type Repository struct {
	file     *os.File
	filePath string
}

// GetAll returns all transactions
func (r *Repository) GetAll() ([]*transaction.Entity, error) {

	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1

	lines, err := reader.ReadAll()
	if err != nil {
		return []*transaction.Entity{}, err
	}

	transactions := []*transaction.Entity{}

	for _, line := range lines[5 : len(lines)-2] {

		// If not a debt, then is a credit
		valueString := line[3]
		if line[4] != "" {
			valueString = line[4]
		}
		if valueString == "" && line[4] == "" {
			return []*transaction.Entity{}, errors.New("invalid input file")
		}

		valueString = strings.Replace(valueString, ",", "", -1)
		valueString = strings.Replace(valueString, ".", "", -1)
		value, err := strconv.ParseInt(valueString, 10, 64)
		if err != nil {
			return []*transaction.Entity{}, err
		}

		date, err := time.Parse("02-01-2006", line[0])
		if err != nil {
			return []*transaction.Entity{}, err
		}

		t, err := transaction.New(date, line[2], value)
		if err != nil {
			return []*transaction.Entity{}, err
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}
