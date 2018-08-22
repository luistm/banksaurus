package changesellername

import (
	"errors"
	"github.com/luistm/banksaurus/seller"
)

// ErrSellerRepositoryUndefined ...
var ErrSellerRepositoryUndefined = errors.New("seller repositoryStub is undefined")

// NewInteractor creates a new instance
// of the changesellername interactor
func NewInteractor(sellers SellerRepository) (*Interactor, error) {
	if sellers == nil {
		return &Interactor{}, ErrSellerRepositoryUndefined
	}
	return &Interactor{sellers}, nil
}

// Interactor for changesellername
type Interactor struct {
	sellers SellerRepository
}

// Execute the changesellername interactor
func (i *Interactor) Execute(r RequestChangeSellerName) error {

	sellerID, err := r.SellerID()
	if err != nil {
		return err
	}

	s, err := i.sellers.GetByID(sellerID)
	if err != nil {
		return err
	}

	if s == nil {
		return nil
	}

	sellerName, err := r.SellerName()
	if err != nil {
		return err
	}

	s, err = seller.New(s.ID(), sellerName)
	if err != nil {
		return err
	}

	err = i.sellers.UpdateSeller(s)
	if err != nil {
		return err
	}

	return nil
}
