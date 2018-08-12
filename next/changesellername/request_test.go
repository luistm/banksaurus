package changesellername_test

import (
	"github.com/luistm/banksaurus/next/changesellername"
	"github.com/luistm/testkit"
	"testing"
)

func TestUnitNewRequest(t *testing.T) {
	t.Run("Returns no error", func(t *testing.T) {
		_, err := changesellername.NewRequest("sellerID", "sellerName")
		testkit.AssertIsNil(t, err)
	})

	t.Run("Returns error if sellerID is string zero value", func(t *testing.T) {
		_, err := changesellername.NewRequest("", "sellerName")
		testkit.AssertEqual(t, changesellername.ErrInvalidSellerID, err)
	})

	t.Run("Returns error if seller name is string zero value", func(t *testing.T) {
		_, err := changesellername.NewRequest("sellerID", "")
		testkit.AssertEqual(t, changesellername.ErrInvalidSellerName, err)
	})
}
