package reportgrouped

import (
	"github.com/luistm/banksaurus/banklib"
	"github.com/luistm/banksaurus/banklib/seller"
	"github.com/luistm/banksaurus/banklib/transaction"
	"github.com/luistm/banksaurus/bankservices"
)

// New creates a service instance
func New(
	transactionsRepository banklib.Repository,
	sellersRepository banklib.Repository,
	presenter bankservices.Presenter,
) (*Service, error) {

	if transactionsRepository == nil ||
		sellersRepository == nil {
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

// Service produces a report from a collection of transactions
// grouped by seller name.
type Service struct {
	transactions banklib.Repository
	sellers      banklib.Repository
	presenter    bankservices.Presenter
}

// Execute the service
func (i *Service) Execute() error {
	var ts []banklib.Entity

	// Get all transaction. If there are no transaction, return
	allTransactions, err := i.transactions.GetAll()
	if err != nil {
		return &banklib.ErrRepository{Msg: err.Error()}
	}
	if len(allTransactions) == 0 {
		return nil
	}

	// Populate the seller with a name if it is available
	for _, t := range allTransactions {
		allSellers, err := i.sellers.GetAll() // FIXME: For each transaction, fetch only the needed seller, not all the seller
		if err != nil {
			return &banklib.ErrRepository{Msg: err.Error()}
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

	transactionsMap := map[string]banklib.Entity{}
	var returnTransactions []banklib.Entity

	for _, t := range ts {
		// FIXME: I'm seeing a lot of type assertion and i don't like it. Code smell??
		if _, ok := transactionsMap[t.(*transaction.Transaction).Seller.String()]; ok {
			tmpValue := transactionsMap[t.(*transaction.Transaction).Seller.String()].(*transaction.Transaction).Value().Add(*t.(*transaction.Transaction).Value())
			s := transactionsMap[t.(*transaction.Transaction).Seller.String()].(*transaction.Transaction).Seller
			newTransaction, err := transaction.New(s, tmpValue.String())
			if err != nil {
				return err
			}
			transactionsMap[t.(*transaction.Transaction).Seller.String()] = newTransaction
		} else {
			transactionsMap[t.(*transaction.Transaction).Seller.String()] = t
		}
	}

	for _, v := range transactionsMap {
		returnTransactions = append(returnTransactions, v)
	}
	if err := i.presenter.Present(returnTransactions...); err != nil {
		return &banklib.ErrPresenter{Msg: err.Error()}
	}

	return nil
}
