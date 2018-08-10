package reportgrouped

import (
	"github.com/luistm/banksaurus/next/entity/seller"
	"github.com/luistm/banksaurus/next/entity/transaction"
)

// Presenter interface for formatting data for presentation
type Presenter interface {
	Present([]map[string]int64) error
}

// TransactionsRepository is a collection of transactions to be used by the report
type TransactionsRepository interface {
	GetAll() ([]*transaction.Entity, error)
}

// SellersRepository is a collection of sellers to be used by the report
type SellersRepository interface {
	GetByID() ([]*seller.Entity, error)
}
