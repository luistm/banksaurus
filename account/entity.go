package account

import "github.com/luistm/banksaurus/money"

// New creates a new instance of account
func New(balance *money.Money) (*Entity, error) {
	return &Entity{balance}, nil
}

// Entity account
type Entity struct {
	balance *money.Money
}

// Balance for the account
func (m *Entity) Balance() *money.Money {
	return m.balance
}
