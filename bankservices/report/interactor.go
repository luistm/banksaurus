package report

import (
	"github.com/luistm/banksaurus/banklib"
	"github.com/luistm/banksaurus/banklib/seller"
	"github.com/luistm/banksaurus/banklib/transaction"
	"github.com/luistm/banksaurus/bankservices"
)

// NewFromString creates a new service instance
func New(
	transactionsRepository banklib.Repository,
	sellersRepository banklib.Repository,
	presenter bankservices.Presenter,
) (*Service, error) {

	if transactionsRepository == nil || sellersRepository == nil {
		return &Service{}, banklib.ErrRepositoryUndefined
	}
	if presenter == nil {
		return &Service{}, banklib.ErrPresenterUndefined
	}

	return &Service{
		transactions: transactionsRepository,
		sellers:      sellersRepository,
		presenter:    presenter,
	}, nil
}

// Service produces a report from an collection of transactions.
type Service struct {
	transactions banklib.Repository
	sellers      banklib.Repository
	presenter    bankservices.Presenter
}

// Execute the service
func (i *Service) Execute() error {

	var ts []banklib.Entity

	transactionsList, err := i.transactions.GetAll()
	if err != nil {
		return &banklib.ErrRepository{Msg: err.Error()}
	}
	if len(transactionsList) == 0 {
		return nil
	}

	for _, t := range transactionsList {
		// FIXME: For each transaction, fetch only the needed seller, not all the seller
		allSellers, err := i.sellers.GetAll()
		if err != nil {
			return &banklib.ErrRepository{Msg: err.Error()}
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
		return &banklib.ErrPresenter{Msg: err.Error()}
	}

	return nil
}
