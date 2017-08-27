package categories

import "errors"

// NewCategory allows the creation of a new category
func NewCategory(name string) (Category, error) {

	c := Category{name: name}

	return c, errors.New("Could not create a new category")
}
