package load

import (
	"github.com/luistm/banksaurus/next/entity/seller"
	"github.com/luistm/banksaurus/next/entity/transaction"
)

// TransactionRepository is a collection of entities to be used by the report
type TransactionRepository interface {
	GetAll() ([]*transaction.Entity, error)
}

// SellerRepository is a collection of sellers
type SellerRepository interface {
	Save(entity *seller.Entity) error
}
