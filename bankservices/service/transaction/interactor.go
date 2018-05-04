package transaction

import "github.com/luistm/banksaurus/bankservices"

// New creates a transaction show interactor
func New() (bankservices.Interactor, error) {
	return &TransactionsShow{}, nil
}

// TransactionsShow shows transaction
type TransactionsShow struct{}

// Execute ...
func (ts *TransactionsShow) Execute() error {
	return nil
}
