package seller

import (
	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/seller"
	"github.com/luistm/banksaurus/services"
)

// New creates a service instance
func New(storage lib.SQLInfrastructer, presenter services.Presenter) *Service {
	return &Service{
		repository: &seller.Sellers{SQLStorage: storage},
		presenter:  presenter,
	}
}

// Service ...
type Service struct {
	repository lib.Repository
	presenter  services.Presenter
}

// Create adds a new seller and persists it
func (i *Service) Create(name string) error {

	if name == "" {
		return lib.ErrBadInput
	}

	if i.repository == nil {
		return lib.ErrRepositoryUndefined
	}

	s := seller.New(name, "")
	if err := i.repository.Save(s); err != nil {
		return &lib.ErrRepository{Msg: err.Error()}
	}

	return nil
}

// GetAll returns all the seller available in the system
func (i *Service) GetAll() error {

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
func (i *Service) Update(slug string, name string) error {

	if slug == "" || name == "" {
		return lib.ErrBadInput
	}
	if i.repository == nil {
		return lib.ErrRepositoryUndefined
	}

	s := seller.New(slug, name)
	err := i.repository.Save(s)
	if err != nil {
		return &lib.ErrRepository{Msg: err.Error()}
	}

	return nil
}
