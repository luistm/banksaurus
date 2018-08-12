package changesellername_test

import (
	"testing"
	"github.com/luistm/testkit"
	"github.com/luistm/banksaurus/next/changesellername"
)

func TestUnitNewRequest(t *testing.T){
	t.Run("Returns no error", func(t *testing.T) {
		_, err := changesellername.NewRequest("someInput")
		testkit.AssertIsNil(t, err)
	})

	t.Run("Returns error if sellerID is string zero value", func(t *testing.T) {
		_, err := changesellername.NewRequest("")
		testkit.AssertEqual(t, changesellername.ErrInvalidInput, err)
	})
}
