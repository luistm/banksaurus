package listsellers_test

import (
	"github.com/luistm/banksaurus/next/entity/seller"
	"github.com/luistm/banksaurus/next/listsellers"
	"github.com/luistm/testkit"
	"testing"
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

	t.Run("Does not return error if sellers repository is defined", func(t *testing.T) {
		_, err := listsellers.NewInteractor(&sellerRepository{})
		testkit.AssertIsNil(t, err)
	})
}

func TestUnitListSellers(t *testing.T) {

	testCases := []struct {
		name            string
		repository      *sellerRepository
		expectedSellers []*seller.Entity
		expectedError   error
	}{
		{},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			i, err := listsellers.NewInteractor(tc.repository)
			testkit.AssertIsNil(t, err)

			err = i.Execute()

			testkit.AssertEqual(t, tc.expectedError, err)
			testkit.AssertEqual(t, tc.expectedSellers, tc.expectedSellers)
		})
	}
}
