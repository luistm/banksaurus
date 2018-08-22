package listtransactionsgrouped

import (
	"errors"
	"github.com/luistm/banksaurus/money"
	"github.com/luistm/banksaurus/transaction"
)

// ErrTransactionsRepositoryUndefined ...
var ErrTransactionsRepositoryUndefined = errors.New("transactions repository is not defined")

// ErrPresenterUndefined ...
var ErrPresenterUndefined = errors.New("presenter is not defined")

// NewInteractor creates a new interactor instance
func NewInteractor(tr TransactionGateway, p Presenter) (*Interactor, error) {
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
	transactions TransactionGateway
}

// Execute the report grouped interactor
func (i *Interactor) Execute() error {

	ts, err := i.transactions.GetAll()
	if err != nil {
		return err
	}

	presenterData := []map[string]*money.Money{}
	transactionsForSeller := []*transaction.Entity{}
	sellersSeen := map[string]bool{}

	for _, t := range ts {
		_, ok := sellersSeen[t.Seller().ID()]
		if ok {
			continue
		}
		sellersSeen[t.Seller().ID()] = true

		transactionsForSeller, err = i.transactions.GetBySeller(t.Seller().ID())
		if err != nil {
			return err
		}

		// TODO: If transactionsForSeller len is zero, something wicked happened
		//       beware of it...
		var sellerTotal *money.Money

		for c, tfs := range transactionsForSeller {
			if c == 0 {
				sellerTotal = t.Value()
				continue
			}
			sellerTotal, err = sellerTotal.Add(tfs.Value())
			if err != nil {
				return err
			}
		}

		presenterData = append(presenterData, map[string]*money.Money{t.Seller().ID(): sellerTotal})
	}

	err = i.presenter.Present(presenterData)
	if err != nil {
		return err
	}

	return nil
}
