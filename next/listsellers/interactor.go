package listsellers

import "github.com/pkg/errors"

var (
	ErrSellersRepositoryUndefined   = errors.New("seller repository is undefined")
	ErrPresenterRepositoryUndefined = errors.New("seller presenter is undefined")
)

// Creates a new interactor instance
func NewInteractor(sr SellerRepository, sp SellerPresenter) (*Interactor, error) {
	if sr == nil {
		return &Interactor{}, ErrSellersRepositoryUndefined
	}

	if sp == nil {
		return &Interactor{}, ErrPresenterRepositoryUndefined
	}

	return &Interactor{sr, sp}, nil
}

// Interactor to list sellers
type Interactor struct {
	sellers   SellerRepository
	presenter SellerPresenter
}

func (i *Interactor) Execute() error {

	sellers, err := i.sellers.GetAll()
	if err != nil {
		return err
	}

	sellersToPresenter := []string{}
	for _, s := range sellers {
		sellersToPresenter = append(sellersToPresenter, s.ID())
	}

	err = i.presenter.Present(sellersToPresenter)
	if err != nil {
		return err
	}

	return nil
}
