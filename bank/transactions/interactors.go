package transactions

import (
	"github.com/luistm/go-bank-cli/lib"
)

type interactor struct {
	repository iRepository
}

// Load fetches raw data from a repository and processes it into objects
// to be persisted in storage.
func (i *interactor) Load() error {

	if i.repository == nil {
		return lib.ErrRepositoryUndefined
	}

	_, err := i.repository.GetAll()
	if err != nil {
		return &lib.ErrRepository{Msg: err.Error()}
	}

	return nil
}
