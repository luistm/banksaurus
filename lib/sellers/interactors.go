package sellers

import (
	"github.com/luistm/go-bank-cli/lib"
	"github.com/luistm/go-bank-cli/lib/customerrors"
)

// NewInteractor creates a new interactor object for sellers
func NewInteractor(storage lib.SQLDatabaseHandler) *interactor {
	return &interactor{
		repository: &repository{SQLStorage: storage},
	}
}

// interactor ...
type interactor struct {
	repository lib.Repository
}

// Create adds a new seller and persists it
func (i *interactor) Create(name string) (lib.Entity, error) {

	if name == "" {
		return &Seller{}, customerrors.ErrBadInput
	}

	if i.repository == nil {
		return &Seller{}, customerrors.ErrRepositoryUndefined
	}

	s := &Seller{slug: name}
	if err := i.repository.Save(s); err != nil {
		return &Seller{}, &customerrors.ErrRepository{Msg: err.Error()}
	}

	return s, nil
}

// GetAll returns all the sellers available in the system
func (i *interactor) GetAll() ([]lib.Entity, error) {

	sellers := []lib.Entity{}
	if i.repository == nil {
		return sellers, customerrors.ErrRepositoryUndefined
	}

	s, err := i.repository.GetAll()
	if err != nil {
		return sellers, &customerrors.ErrRepository{Msg: err.Error()}
	}

	return s, nil
}

// Updates a seller given i'ts ID
func (i *interactor) Update(ID string, name string) error {

	if ID == "" || name == "" {
		return customerrors.ErrBadInput
	}
	if i.repository == nil {
		return customerrors.ErrRepositoryUndefined
	}

	return nil
}
