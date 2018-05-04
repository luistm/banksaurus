package loaddata

import (
	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/transaction"
)

// New creates a new  Interactor to loaddata data from records
func New(transactionsRepository lib.Repository, sellerRepository lib.Repository) *LoadDataFromRecords {
	return &LoadDataFromRecords{
		transactions: transactionsRepository,
		sellers:      sellerRepository,
	}
}

// LoadDataFromRecords saves records into transactions
type LoadDataFromRecords struct {
	transactions lib.Repository
	sellers      lib.Repository
}

// Execute the LoadDataFromRecords interactor
func (i *LoadDataFromRecords) Execute() error {

	if i.transactions == nil {
		return lib.ErrRepositoryUndefined
	}

	ts, err := i.transactions.GetAll()
	if err != nil {
		return &lib.ErrRepository{Msg: err.Error()}
	}

	if i.sellers == nil {
		return lib.ErrInteractorUndefined
	}

	for _, t := range ts {
		err := i.sellers.Save(t.(*transaction.Transaction).Seller)
		if err != nil {
			return &lib.ErrInteractor{Msg: err.Error()}
		}
	}

	return nil
}
