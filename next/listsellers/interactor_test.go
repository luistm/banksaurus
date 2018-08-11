package listsellers_test

import (
	"github.com/luistm/banksaurus/next/listsellers"
	"github.com/luistm/testkit"
	"testing"
	"github.com/luistm/banksaurus/next/entity/seller"
)

type sellerRepository struct{}

func (*sellerRepository) GetAll() []*seller.Entity {
	panic("implement me")
}

func TestUnitNewInteractor(t *testing.T) {

	t.Run("Returns error if sellers repository undefined", func(t *testing.T) {
		_, err := listsellers.NewInteractor(nil)
		testkit.AssertEqual(t, listsellers.ErrSellersRepositoryUndefined, err)
	})

	t.Run("Does not return error if sellers repository is defined", func(t *testing.T){
		_, err := listsellers.NewInteractor(&sellerRepository{})
		testkit.AssertIsNil(t, err)
	})
}