package reportgrouped

import (
	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/seller"
	"github.com/luistm/banksaurus/lib/transaction"
	"github.com/luistm/banksaurus/services"
)

// NewFromString creates a service instance
func New(
	transactionsRepository lib.Repository,
	sellersRepository lib.Repository,
	presenter services.Presenter,
) (*Service, error) {

	if transactionsRepository == nil ||
		sellersRepository == nil {
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

// Service produces a report from a collection of transactions
// grouped by seller name.
type Service struct {
	transactions lib.Repository
	sellers      lib.Repository
	presenter    services.Presenter
}

// Execute the service
func (i *Service) Execute() error {
	var ts []lib.Entity

	// Get all transaction. If there are no transaction, return
	allTransactions, err := i.transactions.GetAll()
	if err != nil {
		return &lib.ErrRepository{Msg: err.Error()}
	}
	if len(allTransactions) == 0 {
		return nil
	}

	// Populate the seller with a name if it is available
	for _, t := range allTransactions {
		allSellers, err := i.sellers.GetAll() // FIXME: For each transaction, fetch only the needed seller, not all the seller
		if err != nil {
			return &lib.ErrRepository{Msg: err.Error()}
		}
		for _, s := range allSellers {
			if s.ID() == t.(*transaction.Transaction).Seller.ID() {
				// TODO: This could a method... Transaction.mergeSeller(s)
				t.(*transaction.Transaction).Seller = s.(*seller.Seller)
				break
			}
		}
		ts = append(ts, t.(*transaction.Transaction))
	}

	transactionsMap := map[string]lib.Entity{}
	var returnTransactions []lib.Entity

	// TODO: Transaction.Add(t). Should return err if transaction does not have the same seller and name
	for _, t := range ts {
		sellerID := t.(*transaction.Transaction).Seller.ID()
		value := *t.(*transaction.Transaction).Value()

		tmp := t
		if _, ok := transactionsMap[sellerID]; ok {
			sum := transactionsMap[sellerID].(*transaction.Transaction).Value().Add(value)
			s := transactionsMap[sellerID].(*transaction.Transaction).Seller
			tmp = transaction.NewFromDecimal(s, &sum)
		}
		transactionsMap[sellerID] = tmp
	}
	// ------------------------------- TODO

	for _, v := range transactionsMap {
		returnTransactions = append(returnTransactions, v)
	}
	if err := i.presenter.Present(returnTransactions...); err != nil {
		return &lib.ErrPresenter{Msg: err.Error()}
	}

	return nil
}
