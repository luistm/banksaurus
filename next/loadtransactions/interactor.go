package loadtransactions

import (
	"errors"
)

var (
	// ErrTransactionRepositoryUndefined ...
	ErrTransactionRepositoryUndefined = errors.New("transaction repository is not defined")
)

// NewInteractor creates an interactor instance
func NewInteractor(tr TransactionGateway) (*Interactor, error) {
	if tr == nil {
		return &Interactor{}, ErrTransactionRepositoryUndefined
	}

	return &Interactor{tr}, nil
}

// Interactor for loading transactions
type Interactor struct {
	transactions TransactionGateway
}

// Execute the load transactions interactor
func (i *Interactor) Execute(r Request) error {

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
