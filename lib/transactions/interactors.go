package transactions

import (
	"errors"

	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/customerrors"
)

// NewInteractor creates a new transactions Interactor
func NewInteractor(
	transactionsRepository Fetcher,
	sellerRepository lib.Repository,
	presenter lib.Presenter,
) *Interactor {
	return &Interactor{
		transactionsRepository: transactionsRepository,
		sellersRepository:      sellerRepository,
		presenter:              presenter,
	}
}

// Interactor for transactions ...
type Interactor struct {
	transactionsRepository Fetcher
	sellersRepository      lib.Repository
	presenter              lib.Presenter
	transactions           []*Transaction
	donUsePresenter        bool
}

// LoadDataFromRecords fetches raw data from a repository and processes it into objects
// to be persisted in storage.
func (i *Interactor) LoadDataFromRecords() error {

	if i.transactionsRepository == nil {
		return customerrors.ErrRepositoryUndefined
	}

	transactions, err := i.transactionsRepository.GetAll()
	if err != nil {
		return &customerrors.ErrRepository{Msg: err.Error()}
	}

	if i.sellersRepository == nil {
		return customerrors.ErrInteractorUndefined
	}

	for _, t := range transactions {
		err := i.sellersRepository.Save(t.(*Transaction).Seller)
		if err != nil {
			return &customerrors.ErrInteractor{Msg: err.Error()}
		}
	}

	return nil
}

func mergeTransactions(transactions []*Transaction) ([]lib.Entity, error) {
	transactionsMap := map[string]*Transaction{}
	returnTransactions := []lib.Entity{}

	for _, t := range transactions {
		if t.Seller == nil {
			return []lib.Entity{}, errors.New("cannot merge transaction whitout seller")
		}

		if _, ok := transactionsMap[t.Seller.String()]; ok {
			tmpValue := transactionsMap[t.Seller.String()].value.Add(*t.Value())
			transactionsMap[t.Seller.String()].value = &tmpValue
		} else {
			transactionsMap[t.Seller.String()] = t
		}
	}

	for _, v := range transactionsMap {
		returnTransactions = append(returnTransactions, v)
	}

	return returnTransactions, nil
}

// ReportFromRecordsGroupedBySeller products a report which ggroups
func (i *Interactor) ReportFromRecordsGroupedBySeller() error {
	// TODO: This should have some unit tests

	//i.donUsePresenter = true
	//err := i.ReportFromRecords()
	//if err != nil {
	//	return err
	//}

	transactions, err := mergeTransactions(i.transactions)
	if err != nil {
		return err
	}

	if i.presenter == nil {
		return customerrors.ErrPresenterUndefined
	}

	if err := i.presenter.Present(transactions...); err != nil {
		return &customerrors.ErrPresenter{Msg: err.Error()}
	}

	return nil
}
