package load

import (
	"errors"
	"github.com/luistm/banksaurus/next/entity/seller"
)

var (
	// ErrTransactionRepositoryUndefined ...
	ErrTransactionRepositoryUndefined = errors.New("transaction repository is not defined")

	// ErrSellerRepositoryUndefined ...
	ErrSellerRepositoryUndefined = errors.New("seller repository is not defined")
)

// NewInteractor creates an interactor instance
func NewInteractor(tr TransactionRepository, sr SellerRepository) (*Interactor, error) {

	if tr == nil {
		return &Interactor{}, ErrTransactionRepositoryUndefined
	}

	if sr == nil {
		return &Interactor{}, ErrSellerRepositoryUndefined
	}

	return &Interactor{tr, sr}, nil
}

// Interactor for load
type Interactor struct {
	transactions TransactionRepository
	sellers      SellerRepository
}

// Execute the load interactor
func (i *Interactor) Execute() error {

	ts, _ := i.transactions.GetAll()
	i.transactions.Save(ts)

	for _, t := range ts {
		s, err := seller.NewFromID(t.Seller())
		if err != nil {
			return err
		}

		i.sellers.Save(s)
	}

	return nil
}
