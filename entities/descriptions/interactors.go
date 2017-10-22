package descriptions

import "github.com/luistm/go-bank-cli/entities"
import "github.com/luistm/go-bank-cli/infrastructure"

// NewInteractor creates a new interactor object for descriptions
func NewInteractor(storage infrastructure.SQLStorage) *Interactor {
	return &Interactor{
		repository: &repository{SQLStorage: storage},
	}
}

// Interactor ...
type Interactor struct {
	repository IRepository
}

// Add adds a new description
func (i *Interactor) Add(name string) (*Description, error) {

	if name == "" {
		return &Description{}, entities.ErrBadInput
	}

	if i.repository == nil {
		return &Description{}, entities.ErrRepositoryIsNil
	}

	d := &Description{rawName: name}
	if err := i.repository.Save(d); err != nil {
		return &Description{}, &entities.ErrRepository{Msg: err.Error()}
	}

	return d, nil
}
