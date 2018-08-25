package transaction

import (
	"errors"
	"fmt"
	"github.com/luistm/banksaurus/money"
	"github.com/luistm/banksaurus/seller"
	"time"
)

var (
	// ErrInvalidTransactionID ...
	ErrInvalidTransactionID = errors.New("invalid transaction ID")

	// ErrInvalidDate ...
	ErrInvalidDate = errors.New("invalid transaction date")

	// ErrInvalidSeller ...
	ErrInvalidSeller = errors.New("invalid transaction seller ID")

	// ErrInvalidValue ...
	ErrInvalidValue = errors.New("invalid transaction value")
)

// New creates a new transaction
func New(id uint64, date time.Time, seller *seller.Entity, value *money.Money) (*Entity, error) {
	if id <= 0 {
		return &Entity{}, ErrInvalidTransactionID
	}

	if date.Equal(time.Time{}) {
		return &Entity{}, ErrInvalidDate
	}

	if seller == nil {
		return &Entity{}, ErrInvalidSeller
	}

	if value == nil {
		return &Entity{}, ErrInvalidValue
	}

	return &Entity{date: date, seller: seller, value: value}, nil
}

// Entity represents a transaction
type Entity struct {
	id     uint64
	date   time.Time
	seller *seller.Entity
	value  *money.Money
}

// ID returns the identification of the transaction
func (t *Entity) ID() uint64 {
	return t.id
}

// Seller of the transaction
func (t *Entity) Seller() *seller.Entity {
	return t.seller
}

// Value of the transaction
func (t *Entity) Value() *money.Money {
	return t.value
}

// GoString to satisfy fmt.GoStringer
func (t *Entity) GoString() string {
	return fmt.Sprintf("%d %s %s %d", t.id, t.date, t.seller, t.value)
}
