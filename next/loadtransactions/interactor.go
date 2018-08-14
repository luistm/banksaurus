package loadtransactions

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
func NewInteractor(tr TransactionGateway, sr SellerGateway) (*Interactor, error) {

	if tr == nil {
		return &Interactor{}, ErrTransactionRepositoryUndefined
	}

	if sr == nil {
		return &Interactor{}, ErrSellerRepositoryUndefined
	}

	return &Interactor{tr, sr}, nil
}

// Interactor for loadtransactions
type Interactor struct {
	transactions TransactionGateway
	sellers      SellerGateway
}

// Execute the loadtransactions interactor
func (i *Interactor) Execute() error {

	ts, err := i.transactions.GetAll()
	if err != nil {
		return err
	}

	for _, t := range ts {
		s, err := seller.New(t.Seller(), "")
		if err != nil {
			return err
		}

		err = i.sellers.Save(s)
		if err != nil {
			return err
		}
	}

	return nil
}
