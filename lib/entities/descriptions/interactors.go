package descriptions

import "errors"

// Interactor ...
type Interactor struct {
	Repository IRepository
}

// Add adds a new description
func (i *Interactor) Add(name string) (*Description, error) {

	if i.Repository == nil {
		return &Description{}, errors.New("Repository is not defined")
	}

	return &Description{}, nil
}
