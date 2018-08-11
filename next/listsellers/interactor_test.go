package listsellers_test

import (
	"github.com/luistm/banksaurus/next/listsellers"
	"github.com/luistm/testkit"
	"testing"
)

func TestUnitNewInteractor(t *testing.T) {

	t.Run("Returns error if sellers repository undefined", func(t *testing.T) {
		_, err := listsellers.NewInteractor(nil)
		testkit.AssertEqual(t, listsellers.ErrSellersRepositoryUndefined, err)
	})
}
