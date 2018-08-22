package listsellers

import "github.com/luistm/banksaurus/seller"

// SellerGateway to access seller entities
type SellerGateway interface {
	GetAll() ([]*seller.Entity, error)
}

// PresenterListSellers to present use case result
type PresenterListSellers interface {
	Present([]string) error
}
