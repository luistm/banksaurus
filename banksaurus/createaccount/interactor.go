package createaccount

import "github.com/pkg/errors"

// ErrRepositoryUndefined ...
var ErrRepositoryUndefined = errors.New("account repository is undefined")

// NewInteractor creates a new instance of the create account interactor
func NewInteractor(r AccountRepository) (*Interactor, error) {
	if r == nil {
		return &Interactor{}, ErrRepositoryUndefined
	}
	return &Interactor{r}, nil
}

// Interactor for creating an account
type Interactor struct {
	accounts AccountRepository
}

// Execute the create account interactor
func (i *Interactor) Execute(r RequestCreateAccount) error {

	balance, _ := r.Balance()

	_, err := i.accounts.New(balance)
	if err != nil {
		return err
	}

	return nil
}
