package transactions_show

import "github.com/luistm/banksaurus/bank"

// New creates a transactions show interactor
func New() (bank.Interactor, error){
	return &TransactionsShow{}, nil
}

// TransactionsShow shows transactions
type TransactionsShow struct{}

// Execute ...
func (ts *TransactionsShow) Execute() error{
	return nil
}