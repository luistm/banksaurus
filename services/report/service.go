package report

import (
	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/seller"
	"github.com/luistm/banksaurus/lib/transaction"
	"github.com/luistm/banksaurus/services"
)

// NewFromString creates a new service instance
func New(
	transactionsRepository lib.Repository,
	sellersRepository lib.Repository,
	presenter services.Presenter,
) (*Service, error) {

	if transactionsRepository == nil || sellersRepository == nil {
		return &Service{}, lib.ErrRepositoryUndefined
	}
	if presenter == nil {
		return &Service{}, lib.ErrPresenterUndefined
	}

	return &Service{
		transactions: transactionsRepository,
		sellers:      sellersRepository,
		presenter:    presenter,
	}, nil
}

// Service produces a report from an collection of transactions.
type Service struct {
	transactions lib.Repository
	sellers      lib.Repository
	presenter    services.Presenter
}

// Execute the service
func (i *Service) Execute() error {

	var ts []lib.Entity

	transactionsList, err := i.transactions.GetAll()
	if err != nil {
		return &lib.ErrRepository{Msg: err.Error()}
	}
	if len(transactionsList) == 0 {
		return nil
	}

	for _, t := range transactionsList {
		// FIXME: For each transaction, fetch only the needed seller, not all the seller
		allSellers, err := i.sellers.GetAll()
		if err != nil {
			return &lib.ErrRepository{Msg: err.Error()}
		}

		for _, s := range allSellers {
			if s.ID() == t.(*transaction.Transaction).Seller.ID() {
				t.(*transaction.Transaction).Seller = s.(*seller.Seller)
				break
			}
		}
		ts = append(ts, t.(*transaction.Transaction))
	}

	if err := i.presenter.Present(ts...); err != nil {
		return &lib.ErrPresenter{Msg: err.Error()}
	}

	return nil
}
