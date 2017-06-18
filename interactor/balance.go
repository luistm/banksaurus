package interactor

import "expensetracker/entities"
import "github.com/shopspring/decimal"

// CurrentBalance returns the current balance for an account
func CurrentBalance() decimal.Decimal {
	account := entities.Account{}
	return account.Balance()
}
