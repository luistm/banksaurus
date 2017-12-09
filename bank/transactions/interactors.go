package transactions

import (
	"errors"

	"github.com/luistm/go-bank-cli/lib"
	"github.com/luistm/go-bank-cli/lib/customerrors"
	"github.com/luistm/go-bank-cli/lib/sellers"
)

// NewInteractor creates a new transactions Interactor
func NewInteractor(r *repository, sellerRepository lib.Repository) *Interactor {
	return &Interactor{transactionsRepository: r, sellersRepository: sellerRepository}
}

// Interactor for transactions ...
type Interactor struct {
	transactionsRepository Fetcher
	sellersRepository      lib.Repository
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
		err := i.sellersRepository.Save(t.seller)
		if err != nil {
			return &customerrors.ErrInteractor{Msg: err.Error()}
		}
	}

	return nil
}

// ReportFromRecords makes a report from an input file.
// If a Seller has a pretty name, that name will be used.
func (i *Interactor) ReportFromRecords() (*Report, error) {

	if i.transactionsRepository == nil {
		return &Report{}, customerrors.ErrRepositoryUndefined
	}

	r := &Report{}
	transactions, err := i.transactionsRepository.GetAll()
	if err != nil {
		return r, &customerrors.ErrRepository{Msg: err.Error()}
	}

	if i.sellersRepository == nil {
		return &Report{}, customerrors.ErrRepositoryUndefined
	}

	for _, transaction := range transactions {
		// TODO: Fetch only the needed sellers, not all the sellers
		allSellers, err := i.sellersRepository.GetAll()
		if err != nil {
			return r, &customerrors.ErrRepository{Msg: err.Error()}
		}
		for _, s := range allSellers {
			if s.ID() == transaction.seller.ID() {
				transaction.seller = s.(*sellers.Seller)
				break
			}
		}
	}

	r.transactions = transactions

	return r, nil
}

func mergeTransactions(transactions []*Transaction) ([]*Transaction, error) {
	transactionsMap := map[string]*Transaction{}
	returnTransactions := []*Transaction{}

	for _, t := range transactions {
		if t.seller == nil {
			return []*Transaction{}, errors.New("cannot merge transaction whitout seller")
		}

		if _, ok := transactionsMap[t.seller.ID()]; ok {
			tmpValue := transactionsMap[t.seller.ID()].value.Add(*t.Value())
			transactionsMap[t.seller.ID()].value = &tmpValue
		} else {
			transactionsMap[t.seller.ID()] = t
		}
	}

	for _, v := range transactionsMap {
		returnTransactions = append(returnTransactions, v)
	}

	return returnTransactions, nil
}

// ReportFromRecordsGroupedBySeller products a report which ggroups
func (i *Interactor) ReportFromRecordsGroupedBySeller() (*Report, error) {
	report, err := i.ReportFromRecords()
	if err != nil {
		return &Report{}, err
	}

	transactions, err := mergeTransactions(report.transactions)
	if err != nil {
		return &Report{}, err
	}
	report.transactions = transactions

	return report, nil
}
