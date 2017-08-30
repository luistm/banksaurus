package categories

import (
	"errors"
	"fmt"
)

//IRepository is the interface for repositories which handle categories
type IRepository interface {
	Save(*Category) error
}

// Interactor for categories
type Interactor struct {
	Repository IRepository // TODO: Make repository private
}

// NewCategory allows the creation of a new category
func (i *Interactor) NewCategory(name string) (*Category, error) {

	if name == "" {
		return &Category{}, errors.New("Cannot create category whitout a category name")
	}

	if i.Repository == nil {
		return &Category{}, errors.New("Repository is not defined")
	}

	c := Category{name: name}
	if err := i.Repository.Save(&c); err != nil {
		return &Category{}, fmt.Errorf("Failed to create category: %s", err)
	}

	return &c, nil
}

// GetCategory returns a category by name
func (i *Interactor) GetCategory(name string) (*Category, error) {
	return &Category{}, errors.New("Could not find category")
}
