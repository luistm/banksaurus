package transactions

import "github.com/luistm/go-bank-cli/lib/customerrors"

type interactor struct {
	repository iRepository
}

// Load fetches raw data from a repository and processes it into objects
// to be persisted in storage.
func (i *interactor) Load() error {

	if i.repository == nil {
		return customerrors.ErrRepositoryUndefined
	}

	_, err := i.repository.GetAll()
	if err != nil {
		return &customerrors.ErrRepository{Msg: err.Error()}
	}

	return nil
}
