package listtransactionsgrouped

import (
	"github.com/luistm/banksaurus/banksauruslib/entities/seller"
	"github.com/luistm/banksaurus/banksauruslib/entities/transaction"
)

// Presenter interface for formatting data for presentation
type Presenter interface {
	Present([]map[string]*transaction.Money) error
}

// TransactionGateway is a collection of transactions to be used by the report
type TransactionGateway interface {
	GetAll() ([]*transaction.Entity, error)
	GetBySeller(string) ([]*transaction.Entity, error)
}

// SellerGateway is a collection of sellers to be used by the report
type SellerGateway interface {
	GetByID() ([]*seller.Entity, error)
}
