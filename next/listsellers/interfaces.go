package listsellers

import "github.com/luistm/banksaurus/next/entity/seller"

type SellerGateway interface {
	GetAll() ([]*seller.Entity, error)
}

type SellerPresenter interface {
	Present([]string) error
}
