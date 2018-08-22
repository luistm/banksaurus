package showaccount

import "errors"

var (
	// ErrPresenterIsUndefined ...
	ErrPresenterIsUndefined = errors.New("presenter is undefined")

	// ErrRepositoryIsUndefined ...
	ErrRepositoryIsUndefined = errors.New("repository is undefined")
)

// NewInteractor creates an instance of the interactor
func NewInteractor(p PresenterShowAccount, r AccountRepository) (*Interactor, error) {
	if p == nil {
		return &Interactor{}, ErrPresenterIsUndefined
	}

	if r == nil {
		return &Interactor{}, ErrRepositoryIsUndefined
	}

	return &Interactor{p, r}, nil
}

// Interactor to show account
type Interactor struct {
	presenter PresenterShowAccount
	accounts  AccountRepository
}

// Execute the show account interactor
func (i *Interactor) Execute(r RequestShowAccount) error {

	accountID, err := r.AccountID()
	if err != nil {
		return err
	}

	acc, err := i.accounts.GetByID(accountID)
	if err != nil {
		return err
	}

	err = i.presenter.Present(acc.Balance())
	if err != nil {
		return err
	}

	return nil
}
