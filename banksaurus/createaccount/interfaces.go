package createaccount

import "github.com/luistm/banksaurus/account"

// AccountRepository ...
type AccountRepository interface {
	New() (*account.Entity, error)
}
