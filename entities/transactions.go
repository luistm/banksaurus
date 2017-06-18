package entities

import (
	"strings"

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
