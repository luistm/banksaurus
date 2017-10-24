package descriptions

import "github.com/luistm/go-bank-cli/entities"
import "github.com/luistm/go-bank-cli/infrastructure"

// NewInteractor creates a new interactor object for descriptions
func NewInteractor(storage infrastructure.SQLStorage) *interactor {
	return &interactor{
		repository: &repository{SQLStorage: storage},
	}
}

// interactor ...
type interactor struct {
	repository entities.IRepository
}

// Add adds a new description
func (i *interactor) Add(name string) (*Description, error) {

	if name == "" {
		return &Description{}, entities.ErrBadInput
	}

	if i.repository == nil {
		return &Description{}, entities.ErrRepositoryUndefined
	}

	d := &Description{rawName: name}
	if err := i.repository.Save(d); err != nil {
		return &Description{}, &entities.ErrRepository{Msg: err.Error()}
	}

	return d, nil
}

func (i *interactor) GetAll() ([]*Description, error) {

	descriptions := []*Description{}
	if i.repository == nil {
		return descriptions, entities.ErrRepositoryUndefined
	}

	return []*Description{}, nil
}
