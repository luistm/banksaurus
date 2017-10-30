package sellers

import "github.com/luistm/go-bank-cli/entities"

// NewInteractor creates a new interactor object for sellers
func NewInteractor(storage entities.SQLDatabaseHandler) *interactor {
	return &interactor{
		repository: &repository{SQLStorage: storage},
	}
}

// interactor ...
type interactor struct {
	repository entities.Repository
}

// Create adds a new seller and persists it
func (i *interactor) Create(name string) (*Seller, error) {

	if name == "" {
		return &Seller{}, entities.ErrBadInput
	}

	if i.repository == nil {
		return &Seller{}, entities.ErrRepositoryUndefined
	}

	s := &Seller{slug: name}
	if err := i.repository.Save(s); err != nil {
		return &Seller{}, &entities.ErrRepository{Msg: err.Error()}
	}

	return s, nil
}

// GetAll returns all the sellers available in the system
func (i *interactor) GetAll() ([]entities.Entity, error) {

	sellers := []entities.Entity{}
	if i.repository == nil {
		return sellers, entities.ErrRepositoryUndefined
	}

	s, err := i.repository.GetAll()
	if err != nil {
		return sellers, &entities.ErrRepository{Msg: err.Error()}
	}

	sellers = append(sellers, s...)

	return sellers, nil
}
