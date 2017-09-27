package accounts

import "github.com/shopspring/decimal"

// CurrentBalance returns the current balance for an account
func CurrentBalance() decimal.Decimal {
	account := Account{}
	return account.Balance()
}
