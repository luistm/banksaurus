package loadtransactions

import (
	"errors"
)

var (
	// ErrTransactionRepositoryUndefined ...
	ErrTransactionRepositoryUndefined = errors.New("transaction repository is not defined")

	// ErrAccountRepositoryUndefined ...
	ErrAccountRepositoryUndefined = errors.New("account repository is not defined")

	// ErrAccountNotFound ...
	ErrAccountNotFound = errors.New("account not found")
)

// NewInteractor creates an interactor instance
func NewInteractor(tg TransactionGateway, ag AccountGateway) (*Interactor, error) {
	if tg == nil {
		return &Interactor{}, ErrTransactionRepositoryUndefined
	}

	if ag == nil {
		return &Interactor{}, ErrAccountRepositoryUndefined
	}

	return &Interactor{tg, ag}, nil
}

// Interactor for loading transactions
type Interactor struct {
	transactions TransactionGateway
	accounts     AccountGateway
}

// Execute the load transactions interactor
func (i *Interactor) Execute(r RequestLoadTransactions) error {

	accountID, err := r.AccountID()
	if err != nil {
		return err
	}

	err = i.accounts.Exists(accountID)
	if err != nil {
		return err
	}

	lines, err := r.Lines()
	if err != nil {
		return err
	}

	for _, line := range lines {
		err := i.transactions.NewFromLine(line)
		if err != nil {
			return err
		}
	}

	return nil
}
