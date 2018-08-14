package seller_test

import (
	"github.com/luistm/banksaurus/next/entity/seller"
	"github.com/luistm/testkit"
	"testing"
)

func TestUnitSellerString(t *testing.T) {

	t.Run("Seller string shows id", func(t *testing.T) {
		s, err := seller.New("SellerID", "")
		testkit.AssertIsNil(t, err)

		testkit.AssertEqual(t, s.String(), "SellerID")
	})

	t.Run("Seller string shows name", func(t *testing.T) {
		s, err := seller.New("SellerID", "SellerName")
		testkit.AssertIsNil(t, err)

		testkit.AssertEqual(t, s.String(), "SellerName")
	})

	t.Run("Seller cannot be created without an ID", func(t *testing.T) {
		_, err := seller.New("", "")
		testkit.AssertEqual(t, seller.ErrInvalidSellerID, err)
	})

}
