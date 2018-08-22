package changesellername

import "github.com/luistm/banksaurus/seller"

// SellerRepository interface
type SellerRepository interface {
	GetByID(string) (*seller.Entity, error)
	UpdateSeller(*seller.Entity) error
}

// RequestChangeSellerName ...
type RequestChangeSellerName interface {
	SellerID() (string, error)
	SellerName() (string, error)
}
