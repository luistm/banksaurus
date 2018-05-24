package transaction

import (
	"fmt"
	"strconv"
	"time"

	"github.com/luistm/banksaurus/lib/money"
	"github.com/luistm/banksaurus/lib/seller"
	"github.com/shopspring/decimal"
)

// NewFromDecimal creates a transaction using a money amount from a decimal
func NewFromDecimal(s *seller.Seller, moneyAmount *decimal.Decimal) *Transaction {
	return &Transaction{Seller: s, value: moneyAmount}
}

// NewFromString creates a transaction using a money amount from a string
func NewFromString(s *seller.Seller, value string) (*Transaction, error) {
	m, err := money.New(value)
	if err != nil {
		return &Transaction{}, err
	}

	return &Transaction{Seller: s, value: m.ToDecimal()}, nil
}

// Transaction is a money movement
type Transaction struct {
	id     uint64
	value  *decimal.Decimal
	Seller *seller.Seller
	date   time.Time
}

// ID of a transaction
func (t *Transaction) ID() string {
	return strconv.FormatUint(t.id, 10)
}

// String to satisfy the fmt.Stringer interface
func (t *Transaction) String() string {
	return fmt.Sprintf("%s %s", t.Value(), t.Seller)
}

// Value returns the field money parsed for money
func (t *Transaction) Value() *decimal.Decimal {
	return t.value
}
