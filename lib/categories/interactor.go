package categories

import (
	"errors"
	"fmt"
)

//IRepository is the interface for repositories which handle categories
type IRepository interface {
	Save(*Category) error
	Get(string) (*Category, error)
}

// Interactor for categories
type Interactor struct {
	Repository IRepository // TODO: Make repository private
}

// NewCategory allows the creation of a new category
func (i *Interactor) NewCategory(name string) ([]*Category, error) {

	cs := []*Category{}

	if name == "" {
		return cs, errors.New("Cannot create category whitout a category name")
	}

	if i.Repository == nil {
		return cs, errors.New("Repository is not defined")
	}

	c := Category{Name: name}
	if err := i.Repository.Save(&c); err != nil {
		return cs, fmt.Errorf("Failed to create category: %s", err)
	}

	cs = append(cs, &c)
	return cs, nil
}

// GetCategory returns a category by name
func (i *Interactor) GetCategory(name string) ([]*Category, error) {

	cs := []*Category{}

	if name == "" {
		return cs, errors.New("Cannot get category whitout a category name")
	}

	if i.Repository == nil {
		return cs, errors.New("Repository is not defined")
	}

	c, err := i.Repository.Get(name)
	if err != nil {
		return cs, fmt.Errorf("Failed to get category: %s", err)
	}

	cs = append(cs, c)
	return cs, nil
}
