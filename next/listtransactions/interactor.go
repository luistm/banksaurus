package listtransactions

import (
	"errors"
)

var (
	// ErrPresenterUndefined ...
	ErrPresenterUndefined = errors.New("presenter is not defined")
	// ErrRepositoryUndefined ...
	ErrRepositoryUndefined = errors.New("presenter is not defined")
	// ErrRepository ...
	ErrRepository = &customError{msg: "error in repository"}
	// ErrPresenter ...
	ErrPresenter = &customError{msg: "error in presenter"}
)

// NewInteractor creates a new listtransactions interactor instance
func NewInteractor(p Presenter, r Repository) (*Interactor, error) {

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
	transactions Repository
}

// Execute the interactor
func (i *Interactor) Execute() error {

	ts, err := i.transactions.GetAll()
	if err != nil {
		return ErrRepository.AppendError(err)
	}

	returnData := []map[string]int64{}
	for _, t := range ts {
		transactionData := map[string]int64{t.Seller(): t.Value()}
		returnData = append(returnData, transactionData)
	}

	err = i.presenter.Present(returnData)
	if err != nil {
		return ErrPresenter.AppendError(err)
	}

	return nil
}
