package transactions

import (
	"github.com/shopspring/decimal"
)

// Transaction is a money movement
type Transaction struct {
	value decimal.Decimal
	// d        *descriptions.Description
	// c        *categories.Category
	// isCredit bool
	// date     time.Time
}
