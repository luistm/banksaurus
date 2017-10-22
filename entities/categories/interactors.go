package categories

import (
	"errors"
	"fmt"

	"github.com/luistm/go-bank-cli/infrastructure"

	"github.com/luistm/go-bank-cli/entities"
)

// NewInteractor creates an interactor for categories
func NewInteractor(storage infrastructure.SQLStorage) *Interactor {
	cr := repository{SQLStorage: storage}

	return &Interactor{repository: &cr}
}

// Interactor for categories
type Interactor struct {
	repository IRepository
}

// Add allows the creation of a new category
func (i *Interactor) Add(name string) ([]*Category, error) {

	cs := []*Category{}

	if name == "" {
		return cs, errors.New("Cannot create category whitout a category name")
	}

	if i.repository == nil {
		return cs, entities.ErrRepositoryIsNil
	}

	c := Category{Name: name}
	if err := i.repository.Save(&c); err != nil {
		return cs, fmt.Errorf("Failed to create category: %s", err)
	}

	cs = append(cs, &c)
	return cs, nil
}

// Get fetches all categories
func (i *Interactor) Get() ([]*Category, error) {

	cs := []*Category{}
	if i.repository == nil {
		return cs, entities.ErrRepositoryIsNil
	}

	cs, err := i.repository.GetAll()
	if err != nil {
		return cs, &entities.ErrRepository{Msg: err.Error()}
	}

	return cs, nil
}

// GetCategory returns a category by name
func (i *Interactor) GetCategory(name string) ([]*Category, error) {

	cs := []*Category{}

	if name == "" {
		return cs, errors.New("Cannot get category whitout a category name")
	}

	if i.repository == nil {
		return cs, entities.ErrRepositoryIsNil
	}

	c, err := i.repository.Get(name)
	if err != nil {
		return cs, &entities.ErrRepository{Msg: err.Error()}
	}

	cs = append(cs, c)
	return cs, nil
}
