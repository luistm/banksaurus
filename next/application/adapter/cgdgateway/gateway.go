package cgdgateway

import (
	"errors"
	"github.com/luistm/banksaurus/next/entity/seller"
	"github.com/luistm/banksaurus/next/entity/transaction"
	"strconv"
	"strings"
	"time"
)

// ErrInvalidNumberOfLines ...
var ErrInvalidNumberOfLines = errors.New("number of lines is less than needed")

// New opens and returns a file handler for a CSV file
func New(lines [][]string) (*Repository, error) {
	minLinesInWellFormattedFile := 8
	if len(lines) < minLinesInWellFormattedFile {
		return &Repository{}, ErrInvalidNumberOfLines
	}

	transactionStartLine := 5
	lastTransactionLine := len(lines) - 2
	f := &Repository{lines: lines[transactionStartLine:lastTransactionLine]}

	return f, nil
}

// Repository represents the content of CSV formatted file
type Repository struct {
	lines [][]string
}

// GetBySeller returns transactions for the specified sellers
func (r *Repository) GetBySeller(s *seller.Entity) ([]*transaction.Entity, error) {

	// TODO: The repository should no know that the seller has an ID method.

	transactions := []*transaction.Entity{}

	for _, line := range r.lines {
		sellerID := strings.TrimSpace(line[2])
		if sellerID != s.ID() {
			continue
		}

		// If not a debt, then is a credit
		isDebt := true
		valueString := line[3]
		if line[4] != "" {
			valueString = line[4]
			isDebt = false
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

		if isDebt {
			value = value * -1
		}

		m, err := transaction.NewMoney(value)
		if err != nil {
			return []*transaction.Entity{}, err
		}

		t, err := transaction.New(1, date, sellerID, m)
		if err != nil {
			return []*transaction.Entity{}, err
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}

// GetAll returns all transactions
func (r *Repository) GetAll() ([]*transaction.Entity, error) {

	transactions := []*transaction.Entity{}

	for _, line := range r.lines {
		sellerID := strings.TrimSpace(line[2])

		// If not a debt, then is a credit
		isDebt := true
		valueString := line[3]
		if line[4] != "" {
			valueString = line[4]
			isDebt = false
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

		if isDebt {
			value = value * -1
		}

		m, err := transaction.NewMoney(value)
		if err != nil {
			return []*transaction.Entity{}, err
		}

		t, err := transaction.New(1, date, sellerID, m)
		if err != nil {
			return []*transaction.Entity{}, err
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}
