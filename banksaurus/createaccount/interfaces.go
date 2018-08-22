package createaccount

import (
	"github.com/luistm/banksaurus/account"
	"github.com/luistm/banksaurus/money"
)

// AccountRepository ...
type AccountRepository interface {
	New(*money.Money) (*account.Entity, error)
}

// RequestCreateAccount interface
type RequestCreateAccount interface {
	Balance() (*money.Money, error)
}
