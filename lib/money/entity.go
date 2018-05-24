package money

import (
	"strings"

	"github.com/shopspring/decimal"
)

// New creates a Money instance from a string
func New(amount string) (*Money, error) {
	sanedAmount := amount
	if sanedAmount == "" {
		sanedAmount = "0"
	}

	parsedField := strings.Replace(sanedAmount, ".", "", -1)
	parsedField = strings.Replace(parsedField, ",", ".", -1)
	d, err := decimal.NewFromString(parsedField)
	if err != nil {
		return &Money{}, err
	}

	return &Money{&d}, nil
}

// Money represents the quantity of money
type Money struct {
	amount *decimal.Decimal
}

// ToDecimal returns the amount of money in Decimal format
func (ma *Money) ToDecimal() *decimal.Decimal {
	return ma.amount
}
