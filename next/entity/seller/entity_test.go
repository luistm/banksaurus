package seller

import (
	"github.com/luistm/testkit"
	"testing"
)

func TestUnitSellerString(t *testing.T) {

	t.Run("Seller string shows id", func(t *testing.T) {
		s, err := New("SellerID", "")
		testkit.AssertIsNil(t, err)

		testkit.AssertEqual(t, s.String(), "SellerID")
	})

	t.Run("Seller string shows name", func(t *testing.T) {
		s, err := New("SellerID", "SellerName")
		testkit.AssertIsNil(t, err)

		testkit.AssertEqual(t, s.String(), "SellerName")
	})

}
