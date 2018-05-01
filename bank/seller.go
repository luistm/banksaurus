package bank

import (
	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/seller"
)

// NewInteractor creates a new SellerInteractor object for seller
func NewInteractor(storage lib.SQLInfrastructer, presenter Presenter) *SellerInteractor {
	return &SellerInteractor{
		repository: &seller.Sellers{SQLStorage: storage},
		presenter:  presenter,
	}
}

// SellerInteractor ...
type SellerInteractor struct {
	repository lib.Repository
	presenter  Presenter
}

// Create adds a new seller and persists it
func (i *SellerInteractor) Create(name string) error {

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
func (i *SellerInteractor) GetAll() error {

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
func (i *SellerInteractor) Update(slug string, name string) error {

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
