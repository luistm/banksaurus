package seller

import (
	"github.com/luistm/banksaurus/lib"
)

// NewInteractor creates a new Interactor object for seller
func NewInteractor(storage lib.SQLInfrastructer, presenter lib.Presenter) *Interactor {
	return &Interactor{
		repository: &Sellers{SQLStorage: storage},
		presenter:  presenter,
	}
}

// Interactor ...
type Interactor struct {
	repository lib.Repository
	presenter  lib.Presenter
}

// Create adds a new seller and persists it
func (i *Interactor) Create(name string) error {

	if name == "" {
		return lib.ErrBadInput
	}

	if i.repository == nil {
		return lib.ErrRepositoryUndefined
	}

	s := &Seller{slug: name}
	if err := i.repository.Save(s); err != nil {
		return &lib.ErrRepository{Msg: err.Error()}
	}

	return nil
}

// GetAll returns all the seller available in the system
func (i *Interactor) GetAll() error {

	if i.repository == nil {
		return lib.ErrRepositoryUndefined
	}

	sellers, err := i.repository.GetAll()
	if err != nil {
		return &lib.ErrRepository{Msg: err.Error()}
	}

	if i.presenter == nil {
		return lib.ErrPresenterUndefined
	}

	if err := i.presenter.Present(sellers...); err != nil {
		return &lib.ErrPresenter{Msg: err.Error()}
	}

	return nil
}

// Update a seller given it's slug
func (i *Interactor) Update(slug string, name string) error {

	if slug == "" || name == "" {
		return lib.ErrBadInput
	}
	if i.repository == nil {
		return lib.ErrRepositoryUndefined
	}

	s := &Seller{slug, name}
	err := i.repository.Save(s)
	if err != nil {
		return &lib.ErrRepository{Msg: err.Error()}
	}

	return nil
}
