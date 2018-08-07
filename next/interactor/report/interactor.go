package report

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

// NewInteractor creates a new report interactor instance
func NewInteractor(p Presenter, r Repository) (*Interactor, error) {

	if p == nil {
		return &Interactor{}, ErrPresenterUndefined
	}

	if r == nil {
		return &Interactor{}, ErrRepositoryUndefined
	}

	return &Interactor{presenter: p, transactions: r}, nil
}

// Interactor for report
type Interactor struct {
	presenter    Presenter
	transactions Repository
}

// Execute the interactor
func (i *Interactor) Execute(r *Request) error {

	ts, err := i.transactions.GetAll()
	if err != nil {
		return ErrRepository.AppendError(err)
	}

	err = i.presenter.Present(ts)
	if err != nil {
		return ErrPresenter.AppendError(err)
	}

	return nil
}
