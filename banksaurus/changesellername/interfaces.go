package changesellername

import "github.com/luistm/banksaurus/seller"

// SellerRepository interface
type SellerRepository interface {
	GetByID(string) (*seller.Entity, error)
	UpdateSeller(*seller.Entity) error
}
