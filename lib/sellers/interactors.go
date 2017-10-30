package sellers

import "github.com/luistm/go-bank-cli/lib"

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
func (i *interactor) Create(name string) (*Seller, error) {

	if name == "" {
		return &Seller{}, lib.ErrBadInput
	}

	if i.repository == nil {
		return &Seller{}, lib.ErrRepositoryUndefined
	}

	s := &Seller{slug: name}
	if err := i.repository.Save(s); err != nil {
		return &Seller{}, &lib.ErrRepository{Msg: err.Error()}
	}

	return s, nil
}

// GetAll returns all the sellers available in the system
func (i *interactor) GetAll() ([]lib.Entity, error) {

	sellers := []lib.Entity{}
	if i.repository == nil {
		return sellers, lib.ErrRepositoryUndefined
	}

	s, err := i.repository.GetAll()
	if err != nil {
		return sellers, &lib.ErrRepository{Msg: err.Error()}
	}

	sellers = append(sellers, s...)

	return sellers, nil
}
