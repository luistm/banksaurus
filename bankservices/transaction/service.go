package transaction

import (
	"github.com/luistm/banksaurus/banklib"
	"github.com/luistm/banksaurus/banklib/transaction"
	"github.com/luistm/banksaurus/bankservices"
)

// New crates a service instance
func New(storage banklib.SQLInfrastructer, presenter bankservices.Presenter) (bankservices.Servicer, error) {
	r := transaction.NewRepository(nil, storage)
	return &Service{presenter, r}, nil
}

// Service shows transactions
type Service struct {
	presenter    bankservices.Presenter
	transactions banklib.Fetcher
}

// Execute the service
func (ts *Service) Execute() error {

	return nil
}
