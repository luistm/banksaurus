package moneyamount

import (
	"strings"

	"github.com/shopspring/decimal"
)

// New creates a MoneyAmount instance from a string
func New(amount string) (*MoneyAmount, error) {
	sanedAmount := amount
	if sanedAmount == "" {
		sanedAmount = "0"
	}

	parsedField := strings.Replace(sanedAmount, ".", "", -1)
	parsedField = strings.Replace(parsedField, ",", ".", -1)
	d, err := decimal.NewFromString(parsedField)
	if err != nil {
		return &MoneyAmount{}, err
	}

	return &MoneyAmount{&d}, nil
}

// MoneyAmount represents the quantity of money
type MoneyAmount struct {
	amount *decimal.Decimal
}

// AmountToDecimal returns the amount of money in Decimal format
func (ma *MoneyAmount) AmountToDecimal() *decimal.Decimal {
	return ma.amount
}
