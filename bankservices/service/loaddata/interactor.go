package loaddata

import (
	"github.com/luistm/banksaurus/banklib"
	"github.com/luistm/banksaurus/banklib/transaction"
)

// New creates a new  Interactor to loaddata data from records
func New(transactionsRepository banklib.Repository, sellerRepository banklib.Repository) *Service {
	return &Service{
		transactions: transactionsRepository,
		sellers:      sellerRepository,
	}
}

// Service saves records into transactions
type Service struct {
	transactions banklib.Repository
	sellers      banklib.Repository
}

// Execute the Service interactor
func (i *Service) Execute() error {

	if i.transactions == nil {
		return banklib.ErrRepositoryUndefined
	}

	ts, err := i.transactions.GetAll()
	if err != nil {
		return &banklib.ErrRepository{Msg: err.Error()}
	}

	if i.sellers == nil {
		return banklib.ErrInteractorUndefined
	}

	for _, t := range ts {
		err := i.sellers.Save(t.(*transaction.Transaction).Seller)
		if err != nil {
			return &banklib.ErrInteractor{Msg: err.Error()}
		}
	}

	return nil
}
