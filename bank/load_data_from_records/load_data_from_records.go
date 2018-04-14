package load_data_from_records

import (
	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/customerrors"
	"github.com/luistm/banksaurus/lib/transactions"
)

// New creates a new transactions Interactor
func New(
	transactionsRepository lib.Repository,
	sellerRepository lib.Repository,
	presenter lib.Presenter,
) *LoadDataFromRecords {
	return &LoadDataFromRecords{
		transactionsRepository: transactionsRepository,
		sellersRepository:      sellerRepository,
		presenter:              presenter,
	}
}

// Interactor for transactions ...
type LoadDataFromRecords struct {
	transactionsRepository lib.Repository
	sellersRepository      lib.Repository
	presenter              lib.Presenter
	transactions           []*transactions.Transaction
	donUsePresenter        bool
}

// LoadDataFromRecords fetches raw data from a repository and processes it into objects
// to be persisted in storage.
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
		err := i.sellersRepository.Save(t.(*transactions.Transaction).Seller)
		if err != nil {
			return &customerrors.ErrInteractor{Msg: err.Error()}
		}
	}

	return nil
}
