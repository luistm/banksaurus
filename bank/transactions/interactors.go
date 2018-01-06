package transactions

import (
	"errors"

	"github.com/luistm/go-bank-cli/lib"
	"github.com/luistm/go-bank-cli/lib/customerrors"
	"github.com/luistm/go-bank-cli/lib/sellers"
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
	report                 *Report
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
		err := i.sellersRepository.Save(t.(*Transaction).seller)
		if err != nil {
			return &customerrors.ErrInteractor{Msg: err.Error()}
		}
	}

	return nil
}

// ReportFromRecords makes a report from an input file.
// If a Seller has a pretty name, that name will be used.
func (i *Interactor) ReportFromRecords() error {

	if i.transactionsRepository == nil {
		return customerrors.ErrRepositoryUndefined
	}

	r := &Report{}
	transactionsList, err := i.transactionsRepository.GetAll()
	if err != nil {
		return &customerrors.ErrRepository{Msg: err.Error()}
	}

	if i.sellersRepository == nil {
		return customerrors.ErrRepositoryUndefined
	}

	for _, t := range transactionsList {
		// FIXME: For each transaction,
		//        fetch only the needed sellers,
		//        not all the sellers
		allSellers, err := i.sellersRepository.GetAll()
		if err != nil {
			return &customerrors.ErrRepository{Msg: err.Error()}
		}

		for _, s := range allSellers {
			if s.ID() == t.(*Transaction).seller.ID() {
				t.(*Transaction).seller = s.(*sellers.Seller)
				break
			}
		}
		r.transactions = append(r.transactions, t.(*Transaction))
	}

	if i.presenter == nil {
		return customerrors.ErrPresenterUndefined
	}

	return nil
}

func mergeTransactions(transactions []*Transaction) ([]*Transaction, error) {
	transactionsMap := map[string]*Transaction{}
	returnTransactions := []*Transaction{}

	for _, t := range transactions {
		if t.seller == nil {
			return []*Transaction{}, errors.New("cannot merge transaction whitout seller")
		}

		if _, ok := transactionsMap[t.seller.String()]; ok {
			tmpValue := transactionsMap[t.seller.String()].value.Add(*t.Value())
			transactionsMap[t.seller.String()].value = &tmpValue
		} else {
			transactionsMap[t.seller.String()] = t
		}
	}

	for _, v := range transactionsMap {
		returnTransactions = append(returnTransactions, v)
	}

	return returnTransactions, nil
}

// ReportFromRecordsGroupedBySeller products a report which ggroups
func (i *Interactor) ReportFromRecordsGroupedBySeller() error {
	err := i.ReportFromRecords()
	if err != nil {
		return err
	}

	transactions, err := mergeTransactions(i.report.transactions)
	if err != nil {
		return err
	}
	i.report.transactions = transactions

	return nil
}
