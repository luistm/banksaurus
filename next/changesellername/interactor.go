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

type Interactor struct{
	sellers SellerRepository
}

func (i *Interactor) Execute(r *Request) error {

	s, _  := i.sellers.GetByID(r.SellerID())
	s, _ := seller.NewFromSeller(s)
	i.sellers.UpdateSeller(s)

	return nil
}
