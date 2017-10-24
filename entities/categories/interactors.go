package categories

import (
	"github.com/luistm/go-bank-cli/infrastructure"

	"github.com/luistm/go-bank-cli/entities"
)

// NewInteractor creates an interactor for categories
func NewInteractor(storage infrastructure.SQLStorage) *interactor {
	cr := repository{SQLStorage: storage}

	return &interactor{repository: &cr}
}

// interactor for categories
type interactor struct {
	repository IRepository
}

// Add allows the creation of a new category
func (i *interactor) Add(name string) ([]*Category, error) {

	cs := []*Category{}

	if name == "" {
		return cs, nil
	}

	if i.repository == nil {
		return cs, entities.ErrRepositoryUndefined
	}

	c := Category{Name: name}
	if err := i.repository.Save(&c); err != nil {
		return cs, &entities.ErrRepository{Msg: err.Error()}
	}

	cs = append(cs, &c)
	return cs, nil
}

// GetAll fetches all categories
func (i *interactor) GetAll() ([]*Category, error) {

	cs := []*Category{}
	if i.repository == nil {
		return cs, entities.ErrRepositoryUndefined
	}

	cs, err := i.repository.GetAll()
	if err != nil {
		return cs, &entities.ErrRepository{Msg: err.Error()}
	}

	return cs, nil
}

// GetCategory returns a category by name
func (i *interactor) GetCategory(name string) ([]*Category, error) {

	cs := []*Category{}

	if name == "" {
		return cs, nil
	}

	if i.repository == nil {
		return cs, entities.ErrRepositoryUndefined
	}

	c, err := i.repository.Get(name)
	if err != nil {
		return cs, &entities.ErrRepository{Msg: err.Error()}
	}

	cs = append(cs, c)
	return cs, nil
}
