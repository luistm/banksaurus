package loadcsv

import (
	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/seller"
	"github.com/luistm/banksaurus/lib/transaction"
	"github.com/luistm/banksaurus/services"
)

// NewFromString creates a new service instance
func New(fileStorage lib.CSVHandler, relationalStorage lib.SQLInfrastructer) services.Servicer {
	tr := transaction.NewRepository(fileStorage, relationalStorage)
	sr := seller.NewRepository(relationalStorage)

	return &Service{tr, sr}
}

// Service copies data available in a collection of transactionsSource to
// the entities which belong to it.
type Service struct {
	transactions lib.Repository
	sellers      lib.Repository
}

// Execute the service
func (i *Service) Execute() error {

	if i.transactions == nil {
		return lib.ErrRepositoryUndefined
	}

	ts, err := i.transactions.GetAll()
	if err != nil {
		return &lib.ErrRepository{Msg: err.Error()}
	}

	if i.sellers == nil {
		return lib.ErrInteractorUndefined
	}

	for _, t := range ts {
		err := i.sellers.Save(t.(*transaction.Transaction).Seller)
		if err != nil {
			return &lib.ErrRepository{Msg: err.Error()}
		}

		err = i.transactions.Save(t.(*transaction.Transaction))
		// TODO: Handle error
		//if err != nil {
		//	return &lib.ErrRepository{Msg: err.Error()}
		//}
	}

	return nil
}
