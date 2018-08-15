package transaction

import (
	"fmt"
	"github.com/pkg/errors"
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
func New(id uint64, date time.Time, sellerID string, value *Money) (*Entity, error) {
	if id <= 0 {
		return &Entity{}, ErrInvalidTransactionID
	}

	if date.Equal(time.Time{}) {
		return &Entity{}, ErrInvalidDate
	}

	if sellerID == "" {
		return &Entity{}, ErrInvalidSeller
	}

	if value == nil {
		return &Entity{}, ErrInvalidValue
	}

	return &Entity{date: date, sellerID: sellerID, value: value}, nil
}

// Entity represents a transaction
type Entity struct {
	id       uint64
	date     time.Time
	sellerID string
	value    *Money
}

// ID returns the identification of the transaction
func (t *Entity) ID() uint64 {
	return t.id
}

// Seller of the transaction
func (t *Entity) Seller() string {
	return t.sellerID
}

// Value of the transaction
func (t *Entity) Value() *Money {
	return t.value
}

// GoString to satisfy fmt.GoStringer
func (t *Entity) GoString() string {
	return fmt.Sprintf("%d %s %s %d", t.id, t.date, t.sellerID, t.value)
}
