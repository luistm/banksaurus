package categories

import "errors"

// NewCategory allows the creation of a new category
func NewCategory(name string) (Category, error) {

	// c := Category{name: name}

	return Category{}, errors.New("Could not create category")
}

// GetCategory returns a category by name
func GetCategory(name string) (Category, error) {
	return Category{}, errors.New("Could not find category")
}
