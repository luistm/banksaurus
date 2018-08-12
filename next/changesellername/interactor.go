package changesellername

import (
	"errors"
	"github.com/luistm/banksaurus/next/entity/seller"
)

var ErrSellerRepositoryUndefined = errors.New("seller repository is undefined")

func NewInteractor(sellers SellerRepository) (*Interactor, error) {
	if sellers == nil {
		return &Interactor{}, ErrSellerRepositoryUndefined
	}
	return &Interactor{sellers}, nil
}

type Interactor struct {
	sellers SellerRepository
}

func (i *Interactor) Execute(r *Request) error {

	sellerID, err := r.SellerID()
	if err != nil {
		return err
	}

	s, err := i.sellers.GetByID(sellerID)
	if err != nil {
		return err
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
