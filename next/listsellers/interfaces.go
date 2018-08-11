package listsellers

import "github.com/luistm/banksaurus/next/entity/seller"

type SellerRepository interface {
	GetAll() ([]*seller.Entity, error)
}

type SellerPresenter interface {
	Present([]map[string]string) error
}
