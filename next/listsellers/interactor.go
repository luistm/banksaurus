package listsellers

import "github.com/pkg/errors"

// ErrSellersRepositoryUndefined ...
var ErrSellersRepositoryUndefined = errors.New("sellers repository is undefined")

// Creates a new interactor instance
func NewInteractor(sr SellerRepository) (*Interactor, error) {
	if sr == nil {
		return &Interactor{}, ErrSellersRepositoryUndefined
	}

	return &Interactor{}, nil
}

// Interactor to list sellers
type Interactor struct{}

func (i *Interactor) Execute() error {
	return nil
}
