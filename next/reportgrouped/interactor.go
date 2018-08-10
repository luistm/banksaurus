package reportgrouped

import (
	"errors"
	"github.com/luistm/banksaurus/next/entity/transaction"
	"time"
)

// ErrTransactionsRepositoryUndefined ...
var ErrTransactionsRepositoryUndefined = errors.New("transactions repository is not defined")

// ErrSellersRepositoryUndefined ...
var ErrSellersRepositoryUndefined = errors.New("sellers repository is not defined")

// ErrPresenterUndefined ...
var ErrPresenterUndefined = errors.New("presenter is not defined")

// NewInteractor creates a new interactor instance
func NewInteractor(tr TransactionsRepository, sr SellersRepository, p Presenter) (*Interactor, error) {
	if tr == nil {
		return &Interactor{}, ErrTransactionsRepositoryUndefined
	}

	if sr == nil {
		return &Interactor{}, ErrSellersRepositoryUndefined
	}

	if p == nil {
		return &Interactor{}, ErrPresenterUndefined
	}

	return &Interactor{presenter: p, transactions: tr}, nil
}

// Interactor for report grouped
type Interactor struct {
	presenter    Presenter
	transactions TransactionsRepository
}

// Execute the report grouped interactor
func (i *Interactor) Execute() error {

	ts, _ := i.transactions.GetAll()

	presenterData := []map[string]int64{}
	transactionsForSeller := map[string][]*transaction.Entity{}

	groupTransactionsBySeller(ts, transactionsForSeller)

	presenterData, _ = sumValueBySeller(transactionsForSeller, presenterData)

	i.presenter.Present(presenterData)

	return nil
}

// TODO: transform this into Transaction.Add(*Transaction)
func sumValueBySeller(transactionsForSeller map[string][]*transaction.Entity, presenterData []map[string]int64) ([]map[string]int64, error) {
	for s, ts := range transactionsForSeller {
		var sellerTotal int64
		for _, t := range ts {
			sellerTotal += t.Value()
		}

		t, err := transaction.New(time.Now(), s, sellerTotal)
		if err != nil {
			return []map[string]int64{}, err
		}

		presenterData = append(presenterData, map[string]int64{t.Seller(): t.Value()})
	}

	return presenterData, nil
}

// TODO: Transform this into Repository.GetAllBySeller()
func groupTransactionsBySeller(ts []*transaction.Entity, transactionsForSeller map[string][]*transaction.Entity) {
	for _, t := range ts {
		listOfTransactionForSeller, ok := transactionsForSeller[t.Seller()]
		if !ok {
			transactionsForSeller[t.Seller()] = []*transaction.Entity{t}
		} else {
			transactionsForSeller[t.Seller()] = append(listOfTransactionForSeller, t)
		}
	}
}
