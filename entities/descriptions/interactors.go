package descriptions

import "github.com/luistm/go-bank-cli/entities"

// NewInteractor creates a new interactor object for descriptions
func NewInteractor(storage entities.SQLDatabaseHandler) *interactor {
	return &interactor{
		repository: &repository{SQLStorage: storage},
	}
}

// interactor ...
type interactor struct {
	repository entities.Repository
}

// Create adds a new description and persists it
func (i *interactor) Create(name string) (*Description, error) {

	if name == "" {
		return &Description{}, entities.ErrBadInput
	}

	if i.repository == nil {
		return &Description{}, entities.ErrRepositoryUndefined
	}

	d := &Description{slug: name}
	if err := i.repository.Save(d); err != nil {
		return &Description{}, &entities.ErrRepository{Msg: err.Error()}
	}

	return d, nil
}

// GetAll returns all the descriptions available in the system
func (i *interactor) GetAll() ([]entities.Entity, error) {

	descriptions := []entities.Entity{}
	if i.repository == nil {
		return descriptions, entities.ErrRepositoryUndefined
	}

	d, err := i.repository.GetAll()
	if err != nil {
		return descriptions, &entities.ErrRepository{Msg: err.Error()}
	}

	descriptions = append(descriptions, d...)

	return descriptions, nil
}
