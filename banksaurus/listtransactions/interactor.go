package listtransactions

import (
	"errors"
	"github.com/luistm/banksaurus/money"
)

var (
	// ErrPresenterUndefined ...
	ErrPresenterUndefined = errors.New("presenterStub is not defined")
	// ErrRepositoryUndefined ...
	ErrRepositoryUndefined = errors.New("presenterStub is not defined")
)

// NewInteractor creates a new listtransactions interactor instance
func NewInteractor(p Presenter, r TransactionGateway) (*Interactor, error) {

	if p == nil {
		return &Interactor{}, ErrPresenterUndefined
	}

	if r == nil {
		return &Interactor{}, ErrRepositoryUndefined
	}

	return &Interactor{presenter: p, transactions: r}, nil
}

// Interactor for listtransactions
type Interactor struct {
	presenter    Presenter
	transactions TransactionGateway
}

// Execute the interactor
func (i *Interactor) Execute() error {

	ts, err := i.transactions.GetAll()
	if err != nil {
		return err
	}

	returnData := []map[string]*money.Money{}
	for _, t := range ts {
		transactionData := map[string]*money.Money{t.Seller().String(): t.Value()}
		returnData = append(returnData, transactionData)
	}

	err = i.presenter.Present(returnData)
	if err != nil {
		return err
	}

	return nil
}
