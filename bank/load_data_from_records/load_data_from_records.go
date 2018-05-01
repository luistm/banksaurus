package load_data_from_records

import (
	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/transaction"
	"github.com/luistm/banksaurus/bank"
)

// New creates a new transaction Interactor
func New(transactionsRepository lib.Repository, sellerRepository lib.Repository, presenter bank.Presenter) *LoadDataFromRecords {
	return &LoadDataFromRecords{
		transactionsRepository: transactionsRepository,
		sellersRepository:      sellerRepository,
		presenter:              presenter,
	}
}

// LoadDataFromRecords saves records into transaction
type LoadDataFromRecords struct {
	transactionsRepository lib.Repository
	sellersRepository      lib.Repository
	presenter              bank.Presenter
	transactions           []*transaction.Transaction
	donUsePresenter        bool
}

// Execute the LoadDataFromRecords interactor
func (i *LoadDataFromRecords) Execute() error {

	if i.transactionsRepository == nil {
		return lib.ErrRepositoryUndefined
	}

	ts, err := i.transactionsRepository.GetAll()
	if err != nil {
		return &lib.ErrRepository{Msg: err.Error()}
	}

	if i.sellersRepository == nil {
		return lib.ErrInteractorUndefined
	}

	for _, t := range ts {
		err := i.sellersRepository.Save(t.(*transaction.Transaction).Seller)
		if err != nil {
			return &lib.ErrInteractor{Msg: err.Error()}
		}
	}

	return nil
}
