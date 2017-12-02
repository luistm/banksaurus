package sellers

import (
	"github.com/luistm/go-bank-cli/lib"
	"github.com/luistm/go-bank-cli/lib/customerrors"
)

// NewInteractor creates a new Interactor object for sellers
func NewInteractor(storage lib.SQLDatabaseHandler) *Interactor {
	return &Interactor{
		repository: &repository{SQLStorage: storage},
	}
}

// Interactor ...
type Interactor struct {
	repository lib.Repository
}

// Create adds a new seller and persists it
func (i *Interactor) Create(name string) (lib.Entity, error) {

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
func (i *Interactor) GetAll() ([]lib.Entity, error) {

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

// Update a seller given it's slug
func (i *Interactor) Update(slug string, name string) error {

	if slug == "" || name == "" {
		return customerrors.ErrBadInput
	}
	if i.repository == nil {
		return customerrors.ErrRepositoryUndefined
	}

	s := &Seller{slug, name}
	err := i.repository.Save(s)
	if err != nil {
		return &customerrors.ErrRepository{Msg: err.Error()}
	}

	return nil
}
