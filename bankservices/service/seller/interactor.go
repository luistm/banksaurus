package seller

import (
	"github.com/luistm/banksaurus/banklib"
	"github.com/luistm/banksaurus/banklib/seller"
	"github.com/luistm/banksaurus/bankservices"
)

// NewInteractor creates a new Service object for seller
func NewInteractor(storage banklib.SQLInfrastructer, presenter bankservices.Presenter) *Service {
	return &Service{
		repository: &seller.Sellers{SQLStorage: storage},
		presenter:  presenter,
	}
}

// Service ...
type Service struct {
	repository banklib.Repository
	presenter  bankservices.Presenter
}

// Create adds a new seller and persists it
func (i *Service) Create(name string) error {

	if name == "" {
		return banklib.ErrBadInput
	}

	if i.repository == nil {
		return banklib.ErrRepositoryUndefined
	}

	s := seller.New(name, "")
	if err := i.repository.Save(s); err != nil {
		return &banklib.ErrRepository{Msg: err.Error()}
	}

	return nil
}

// GetAll returns all the seller available in the system
func (i *Service) GetAll() error {

	if i.repository == nil {
		return banklib.ErrRepositoryUndefined
	}

	sellers, err := i.repository.GetAll()
	if err != nil {
		return &banklib.ErrRepository{Msg: err.Error()}
	}

	if i.presenter == nil {
		return banklib.ErrPresenterUndefined
	}

	if err := i.presenter.Present(sellers...); err != nil {
		return &banklib.ErrPresenter{Msg: err.Error()}
	}

	return nil
}

// Update a seller given it's slug
func (i *Service) Update(slug string, name string) error {

	if slug == "" || name == "" {
		return banklib.ErrBadInput
	}
	if i.repository == nil {
		return banklib.ErrRepositoryUndefined
	}

	s := seller.New(slug, name)
	err := i.repository.Save(s)
	if err != nil {
		return &banklib.ErrRepository{Msg: err.Error()}
	}

	return nil
}
