package showaccount

import (
	"github.com/luistm/banksaurus/account"
	"github.com/luistm/banksaurus/money"
)

// PresenterShowAccount ...
type PresenterShowAccount interface {
	Present(*money.Money) error
}

// AccountRepository ...
type AccountRepository interface {
	GetByID(string) (*account.Entity, error)
}

// RequestShowAccount ...
type RequestShowAccount interface {
	AccountID() (string, error)
}
