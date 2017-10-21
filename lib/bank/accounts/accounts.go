package accounts

import "github.com/shopspring/decimal"

// Accounts contains the status of an account
type Account struct {
	balance decimal.Decimal
}

// Balance returns the current account balance
func (a *Account) Balance() decimal.Decimal {
	return a.balance
}
