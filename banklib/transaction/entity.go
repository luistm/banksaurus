package transaction

import (
	"fmt"
	"strconv"
	"time"

	"github.com/luistm/banksaurus/banklib/moneyamount"
	"github.com/luistm/banksaurus/banklib/seller"
	"github.com/shopspring/decimal"
)

// New creates a record with some parsed data
func New(s *seller.Seller, value string) (*Transaction, error) {
	moneyAmount, err := moneyamount.New(value)
	if err != nil {
		return &Transaction{}, err
	}

	return &Transaction{Seller: s, isCredit: false, value: moneyAmount.AmountToDecimal()}, nil
}

// Transaction is a money movement
type Transaction struct {
	id       uint64
	value    *decimal.Decimal
	Seller   *seller.Seller
	isCredit bool
	date     time.Time
}

// ID of a transaction
func (t *Transaction) ID() string {
	return strconv.FormatUint(t.id, 10)
}

// String to satisfy the fmt.Stringer interface
func (t *Transaction) String() string {
	return fmt.Sprintf("%s %s", t.Value(), t.Seller)
}

// IsDebt returns true if a transaction is a debt
func (t *Transaction) IsDebt() bool {
	if t.isCredit {
		return false
	}
	return true
}

// Value returns the field moneyamount parsed for money
func (t *Transaction) Value() *decimal.Decimal {
	return t.value
}
