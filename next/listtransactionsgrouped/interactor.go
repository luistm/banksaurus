package listtransactionsgrouped

import (
	"errors"
	"github.com/luistm/banksaurus/next/entity/seller"
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
func NewInteractor(tr TransactionsRepository, p Presenter) (*Interactor, error) {
	if tr == nil {
		return &Interactor{}, ErrTransactionsRepositoryUndefined
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
	transactionsForSeller := []*transaction.Entity{}

	sellersSeen := map[string]bool{}
	for _, t := range ts {
		s, err := seller.NewFromID(t.Seller())
		if err != nil {
			return err
		}

		_, ok := sellersSeen[t.Seller()]
		if ok {
			continue
		}
		sellersSeen[t.Seller()] = true

		transactionsForSeller, _ = i.transactions.GetBySeller(s)

		var sellerTotal int64
		for _, t := range transactionsForSeller {
			sellerTotal += t.Value()
		}

		t, err := transaction.New(time.Now(), s.ID(), sellerTotal)
		if err != nil {
			return err
		}

		presenterData = append(presenterData, map[string]int64{t.Seller(): t.Value()})
	}

	i.presenter.Present(presenterData)

	return nil
}
