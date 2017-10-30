package categories

import (
	"github.com/luistm/go-bank-cli/lib"
)

// NewInteractor creates an interactor for categories
func NewInteractor(storage lib.SQLDatabaseHandler) *interactor {
	cr := repository{SQLStorage: storage}

	return &interactor{repository: &cr}
}

// interactor for categories
type interactor struct {
	repository lib.Repository
}

// Create allows the creation of a new category
func (i *interactor) Create(name string) ([]lib.Entity, error) {

	cs := []lib.Entity{}

	if name == "" {
		return cs, nil
	}

	if i.repository == nil {
		return cs, lib.ErrRepositoryUndefined
	}

	c := Category{name: name}
	if err := i.repository.Save(&c); err != nil {
		return cs, &lib.ErrRepository{Msg: err.Error()}
	}

	cs = append(cs, &c)
	return cs, nil
}

// GetAll fetches all categories
func (i *interactor) GetAll() ([]lib.Entity, error) {

	cs := []lib.Entity{}
	if i.repository == nil {
		return cs, lib.ErrRepositoryUndefined
	}

	cs, err := i.repository.GetAll()
	if err != nil {
		return cs, &lib.ErrRepository{Msg: err.Error()}
	}

	return cs, nil
}

// GetCategory returns a category by name
func (i *interactor) GetCategory(name string) ([]lib.Entity, error) {

	cs := []lib.Entity{}

	if name == "" {
		return cs, nil
	}

	if i.repository == nil {
		return cs, lib.ErrRepositoryUndefined
	}

	c, err := i.repository.Get(name)
	if err != nil {
		return cs, &lib.ErrRepository{Msg: err.Error()}
	}

	cs = append(cs, c)
	return cs, nil
}
