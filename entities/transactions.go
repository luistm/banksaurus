package entities

import (
	"strconv"
	"strings"
)

// CREDIT ...
var CREDIT = "Credit"

// DEBT ..
var DEBT = "Debt"

// TODO: Use shopspring/decimal package for money

// Transaction is a money movement
type Transaction struct {
	// TODO: Make all these propreties private
	field           string
	Description     string
	TransactionType string
}

// New creates a record with some parsed data
func (t *Transaction) New(record Record) *Transaction {

	for i := 0; i < len(record.Record); i++ {
		t.field = record.Record[i]
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
func (t *Transaction) Value() float64 {
	parsedField := strings.Replace(t.field, ".", "", -1)
	parsedField = strings.Replace(parsedField, ",", ".", -1)
	retValue, _ := strconv.ParseFloat(parsedField, 64)

	return retValue
}
