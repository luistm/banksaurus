package loadtransactions

import (
	"github.com/luistm/banksaurus/next/entity/seller"
	"github.com/luistm/banksaurus/next/entity/transaction"
)

// TransactionGateway is a collection of entities to be used by the report
type TransactionGateway interface {
	GetAll() ([]*transaction.Entity, error)
}

// SellerGateway is a collection of sellers
type SellerGateway interface {
	Save(entity *seller.Entity) error
}
