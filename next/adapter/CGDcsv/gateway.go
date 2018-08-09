package CGDcsv

import (
	"errors"
	"github.com/luistm/banksaurus/next/entity/transaction"
	"strconv"
	"strings"
	"time"
)

// ErrInvalidNumberOfLines ...
var ErrInvalidNumberOfLines = errors.New("number of lines is less than needed")

// New opens and returns a file handler for a CSV file
func New(lines [][]string) (*Repository, error) {
	if len(lines) < 8 {
		return &Repository{}, ErrInvalidNumberOfLines
	}

	f := &Repository{lines: lines[5 : len(lines)-2]}
	return f, nil
}

// Repository represents the content of CSV formatted file
type Repository struct {
	lines [][]string
}

// GetAll returns all transactions
func (r *Repository) GetAll() ([]*transaction.Entity, error) {

	transactions := []*transaction.Entity{}

	for _, line := range r.lines {

		// If not a debt, then is a credit
		valueString := line[3]
		if line[4] != "" {
			valueString = line[4]
			continue
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
