package report

import (
	"github.com/luistm/banksaurus/banklib"
	"github.com/luistm/banksaurus/banklib/seller"
	"github.com/luistm/banksaurus/banklib/transaction"
	"github.com/luistm/banksaurus/bankservices"
)

// New creates a new ReportFromRecords use case
func New(
	transactionsRepository banklib.Repository, sellersRepository banklib.Repository, presenter bankservices.Presenter,
) (*ReportFromRecords, error) {

	if transactionsRepository == nil ||
		sellersRepository == nil {
		return &ReportFromRecords{}, banklib.ErrRepositoryUndefined
	}
	if presenter == nil {
		return &ReportFromRecords{}, banklib.ErrPresenterUndefined
	}

	return &ReportFromRecords{
		transactionsRepository: transactionsRepository,
		sellersRepository:      sellersRepository,
		presenter:              presenter,
	}, nil
}

// ReportFromRecords makes a reportgrouped from an input file.
// If a Command has a pretty name, that name will be used.
type ReportFromRecords struct {
	transactionsRepository banklib.Repository
	sellersRepository      banklib.Repository
	presenter              bankservices.Presenter
}

// Execute ...
func (i *ReportFromRecords) Execute() error {

	var ts []banklib.Entity

	transactionsList, err := i.transactionsRepository.GetAll()
	if err != nil {
		return &banklib.ErrRepository{Msg: err.Error()}
	}
	if len(transactionsList) == 0 {
		return nil
	}

	for _, t := range transactionsList {
		// FIXME: For each transaction, fetch only the needed seller, not all the seller
		allSellers, err := i.sellersRepository.GetAll()
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
