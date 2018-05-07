package loaddata

import (
	"github.com/luistm/banksaurus/banklib"
	"github.com/luistm/banksaurus/banklib/seller"
	"github.com/luistm/banksaurus/banklib/transaction"
	"github.com/luistm/banksaurus/bankservices"
)

// NewFromString creates a new service instance
func New(fileStorage banklib.CSVHandler, relationalStorage banklib.SQLInfrastructer) bankservices.Servicer {
	tr := transaction.NewRepository(fileStorage, relationalStorage)
	sr := seller.NewRepository(relationalStorage)

	return &Service{tr, sr}
}

// Service copies data available in a collection of transactionsSource to
// the entities which belong to it.
type Service struct {
	transactions banklib.Repository
	sellers      banklib.Repository
}

// Execute the service
func (i *Service) Execute() error {

	if i.transactions == nil {
		return banklib.ErrRepositoryUndefined
	}

	ts, err := i.transactions.GetAll()
	if err != nil {
		return &banklib.ErrRepository{Msg: err.Error()}
	}

	if i.sellers == nil {
		return banklib.ErrInteractorUndefined
	}

	for _, t := range ts {
		err := i.sellers.Save(t.(*transaction.Transaction).Seller)
		if err != nil {
			return &banklib.ErrRepository{Msg: err.Error()}
		}

		err = i.transactions.Save(t.(*transaction.Transaction))
		// TODO: Handle error
		//if err != nil {
		//	return &banklib.ErrRepository{Msg: err.Error()}
		//}
	}

	return nil
}
