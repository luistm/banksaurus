package reportfromrecordsgrouped

import "github.com/luistm/banksaurus/lib/customerrors"
import (
	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/sellers"
	"github.com/luistm/banksaurus/lib/transaction"
)

// New creates a new ReportFromRecords use case
func New(
	transactionsRepository lib.Repository, sellersRepository lib.Repository, presenter lib.Presenter,
) (*ReportFromRecordsGrouped, error) {

	if transactionsRepository == nil ||
		sellersRepository == nil {
		return &ReportFromRecordsGrouped{}, customerrors.ErrRepositoryUndefined
	}
	if presenter == nil {
		return &ReportFromRecordsGrouped{}, customerrors.ErrPresenterUndefined
	}

	return &ReportFromRecordsGrouped{
		transactionsRepository: transactionsRepository,
		sellersRepository:      sellersRepository,
		presenter:              presenter,
	}, nil
}

// ReportFromRecordsGrouped makes a report from an input file.
// If a Seller has a pretty name, that name will be used.
type ReportFromRecordsGrouped struct {
	transactionsRepository lib.Repository
	sellersRepository      lib.Repository
	presenter              lib.Presenter
}

// Execute an instance of ReportFromRecordsGrouped
func (i *ReportFromRecordsGrouped) Execute() error {
	var ts []lib.Entity

	// Get all transaction. If there are no transaction, return
	allTransactions, err := i.transactionsRepository.GetAll()
	if err != nil {
		return &customerrors.ErrRepository{Msg: err.Error()}
	}
	if len(allTransactions) == 0 {
		return nil
	}

	// Populate the sellers with a name if it is available
	for _, t := range allTransactions {
		allSellers, err := i.sellersRepository.GetAll() // FIXME: For each transaction, fetch only the needed sellers, not all the sellers
		if err != nil {
			return &customerrors.ErrRepository{Msg: err.Error()}
		}
		for _, s := range allSellers {
			if s.ID() == t.(*transaction.Transaction).Seller.ID() {
				// TODO: This could a method... Transaction.mergeSeller(s)
				t.(*transaction.Transaction).Seller = s.(*sellers.Seller)
				break
			}
		}
		ts = append(ts, t.(*transaction.Transaction))
	}

	transactionsMap := map[string]lib.Entity{}
	var returnTransactions []lib.Entity

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
		return &customerrors.ErrPresenter{Msg: err.Error()}
	}

	return nil
}
