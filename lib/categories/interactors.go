package categories

import (
	"github.com/luistm/go-bank-cli/lib"
	"github.com/luistm/go-bank-cli/lib/customerrors"
)

// NewInteractor creates an Interactor for categories
func NewInteractor(storage lib.SQLInfrastructer) *Interactor {
	cr := repository{SQLStorage: storage}

	return &Interactor{repository: &cr}
}

// Interactor for categories
type Interactor struct {
	repository lib.Repository
}

// Create allows the creation of a new category
func (i *Interactor) Create(name string) ([]lib.Entity, error) {

	cs := []lib.Entity{}

	if name == "" {
		return cs, nil
	}

	if i.repository == nil {
		return cs, customerrors.ErrRepositoryUndefined
	}

	c := Category{name: name}
	if err := i.repository.Save(&c); err != nil {
		return cs, &customerrors.ErrRepository{Msg: err.Error()}
	}

	cs = append(cs, &c)
	return cs, nil
}

// GetAll fetches all categories
func (i *Interactor) GetAll() ([]lib.Entity, error) {

	cs := []lib.Entity{}
	if i.repository == nil {
		return cs, customerrors.ErrRepositoryUndefined
	}

	cs, err := i.repository.GetAll()
	if err != nil {
		return cs, &customerrors.ErrRepository{Msg: err.Error()}
	}

	return cs, nil
}

// GetCategory returns a category by name
func (i *Interactor) GetCategory(name string) ([]lib.Entity, error) {

	cs := []lib.Entity{}

	if name == "" {
		return cs, nil
	}

	if i.repository == nil {
		return cs, customerrors.ErrRepositoryUndefined
	}

	c, err := i.repository.Get(name)
	if err != nil {
		return cs, &customerrors.ErrRepository{Msg: err.Error()}
	}

	cs = append(cs, c)
	return cs, nil
}
