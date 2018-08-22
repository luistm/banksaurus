package listtransactions

import (
	"errors"
	"github.com/luistm/banksaurus/transaction"
)

var (
	// ErrPresenterUndefined ...
	ErrPresenterUndefined = errors.New("presenter is not defined")
	// ErrRepositoryUndefined ...
	ErrRepositoryUndefined = errors.New("presenter is not defined")
	// ErrPresenter ...
	ErrPresenter = &customError{msg: "error in presenter"}
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

	returnData := []map[string]*transaction.Money{}
	for _, t := range ts {
		transactionData := map[string]*transaction.Money{t.Seller().ID(): t.Value()}
		returnData = append(returnData, transactionData)
	}

	err = i.presenter.Present(returnData)
	if err != nil {
		return ErrPresenter.AppendError(err)
	}

	return nil
}
