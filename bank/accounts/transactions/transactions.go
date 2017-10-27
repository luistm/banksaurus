package transactions

import (
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

var isCREDIT = "Credit"
var isDEBT = "Debt"

// New creates a record with some parsed data
func (t *Transaction) New(record Record) *Transaction {

	for i := 0; i < len(record.Record); i++ {
		t.value = record.Record[i]
		t.Description = record.Record[2]

		// Expense
		if i == 3 && record.Record[i] != "" {
			t.TransactionType = isDEBT
			return t
		}

		// Credit
		if i == 4 && record.Record[i] != "" {
			t.TransactionType = isCREDIT
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

	retValue, _ := decimal.NewFromString(parsedField)

	return retValue
}

// IsFromThisMonth checks if a transaction is from the current month
func (t *Transaction) IsFromThisMonth() bool {

	// This code sucks so much.... i'm sure there is a better way, i'm just to lazy now...
	splittedDate := strings.Split(t.date, "-")
	year, _ := strconv.Atoi(splittedDate[2])
	month, _ := strconv.Atoi(splittedDate[1])

	if time.Month(month) == time.Now().Month() &&
		year == time.Now().Year() {
		return true
	}
	return false
}

// IsDebt returns true if a transaction is a debt
func (t *Transaction) IsDebt() bool {
	if t.TransactionType == isDEBT {
		return true
	}
	return false
}
