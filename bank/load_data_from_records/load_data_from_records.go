package load_data_from_records

import (
	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/customerrors"
	"github.com/luistm/banksaurus/lib/transaction"
)

// New creates a new transaction Interactor
func New(transactionsRepository lib.Repository, sellerRepository lib.Repository, presenter lib.Presenter) *LoadDataFromRecords {
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
	presenter              lib.Presenter
	transactions           []*transaction.Transaction
	donUsePresenter        bool
}

// Execute the LoadDataFromRecords interactor
func (i *LoadDataFromRecords) Execute() error {

	if i.transactionsRepository == nil {
		return customerrors.ErrRepositoryUndefined
	}

	ts, err := i.transactionsRepository.GetAll()
	if err != nil {
		return &customerrors.ErrRepository{Msg: err.Error()}
	}

	if i.sellersRepository == nil {
		return customerrors.ErrInteractorUndefined
	}

	for _, t := range ts {
		err := i.sellersRepository.Save(t.(*transaction.Transaction).Seller)
		if err != nil {
			return &customerrors.ErrInteractor{Msg: err.Error()}
		}
	}

	return nil
}
