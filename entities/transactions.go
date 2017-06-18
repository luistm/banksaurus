package entities

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

// CREDIT ...
var CREDIT = "Credit"

// DEBT ..
var DEBT = "Debt"

// Transaction is a money movement
type Transaction struct {
	value           string
	Description     string
	TransactionType string
	date            string
}

// New creates a record with some parsed data
func (t *Transaction) New(record Record) *Transaction {

	for i := 0; i < len(record.Record); i++ {
		t.value = record.Record[i]
		t.Description = record.Record[2]

		// Expense
		if i == 3 && record.Record[i] != "" {
			t.TransactionType = DEBT
			return t
		}

		// Credit
		if i == 4 && record.Record[i] != "" {
			t.TransactionType = CREDIT
			return t
		}

		// Transaction Date
		if i == 0 {
			t.date = record.Record[i]
		}
	}

	return t
}

// Value returns the field value parsed for money
func (t *Transaction) Value() decimal.Decimal {
	parsedField := strings.Replace(t.value, ".", "", -1)
	parsedField = strings.Replace(parsedField, ",", ".", -1)
	// retValue, _ := strconv.ParseFloat(parsedField, 64)

	retValue, _ := decimal.NewFromString(parsedField)

	return retValue
}

// IsFromThisMonth checks if a transaction is from the current month
func (t *Transaction) IsFromThisMonth() bool {

	// This code sucks so much.... i'm sure there is a better way, i'm just to lazy now...
	splittedDate := strings.Split(t.date, "-")
	fmt.Println("Date is:", t.date)
	year, _ := strconv.Atoi(splittedDate[2])
	month, _ := strconv.Atoi(splittedDate[1])

	if time.Month(month) == time.Now().Month() &&
		year == time.Now().Year() {
		return true
	}
	return false
}
