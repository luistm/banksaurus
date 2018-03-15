package categories

import (
	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/customerrors"
)

// NewInteractor creates an Interactor for categories
func NewInteractor(storage lib.SQLInfrastructer, presenter lib.Presenter) *Interactor {
	cr := repository{SQLStorage: storage}
	return &Interactor{repository: &cr, presenter: presenter}
}

// Interactor for categories
type Interactor struct {
	repository lib.Repository
	presenter  lib.Presenter
}

// Create allows the creation of a new category
func (i *Interactor) Create(name string) error {

	if name == "" {
		return customerrors.ErrBadInput
	}
	if i.repository == nil {
		return customerrors.ErrRepositoryUndefined
	}

	c := Category{name: name}
	if err := i.repository.Save(&c); err != nil {
		return &customerrors.ErrRepository{Msg: err.Error()}
	}

	return nil
}

// GetAll fetches all categories
func (i *Interactor) GetAll() error {

	if i.repository == nil {
		return customerrors.ErrRepositoryUndefined
	}

	if i.presenter == nil {
		return customerrors.ErrPresenterUndefined
	}

	categories, err := i.repository.GetAll()
	if err != nil {
		return &customerrors.ErrRepository{Msg: err.Error()}
	}

	err = i.presenter.Present(categories...)
	if err != nil {
		return &customerrors.ErrPresenter{Msg: err.Error()}
	}

	return nil
}

// GetCategory returns a category by name
func (i *Interactor) GetCategory(name string) error {

	if name == "" {
		return customerrors.ErrBadInput
	}

	if i.repository == nil {
		return customerrors.ErrRepositoryUndefined
	}

	_, err := i.repository.Get(name)
	if err != nil {
		return &customerrors.ErrRepository{Msg: err.Error()}
	}

	return nil
}
