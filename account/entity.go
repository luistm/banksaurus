package account

import (
	"errors"
	"github.com/luistm/banksaurus/money"
)

var (
	// ErrInvalidID ...
	ErrInvalidID = errors.New("invalid error")

	// ErrInvalidBalance ...
	ErrInvalidBalance = errors.New("balance is undefined")
)

// New creates a new instance of account
func New(id string, balance *money.Money) (*Entity, error) {
	if id == "" {
		return &Entity{}, ErrInvalidID
	}

	if balance == nil {
		return &Entity{}, ErrInvalidBalance
	}
	return &Entity{id, balance}, nil
}

// Entity account
type Entity struct {
	id      string
	balance *money.Money
}

// Balance for the account
func (m *Entity) Balance() *money.Money {
	return m.balance
}

// ID ...
func (m *Entity) ID() string {
	return m.id
}
