package descriptions

import "github.com/luistm/go-bank-cli/entities"

// Interactor ...
type Interactor struct {
	Repository IRepository
}

// Add adds a new description
func (i *Interactor) Add(name string) (*Description, error) {

	if i.Repository == nil {
		return &Description{}, entities.ErrRepositoryIsNil
	}

	return &Description{}, nil
}
