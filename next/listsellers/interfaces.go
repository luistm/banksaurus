package listsellers

import "github.com/luistm/banksaurus/next/entity/seller"

type SellerRepository interface {
	GetAll() []*seller.Entity
}
