package changesellername

import "github.com/luistm/banksaurus/next/entity/seller"

type SellerRepository interface {
	GetByID(string) (*seller.Entity, error)
	UpdateSeller(*seller.Entity) error
}
