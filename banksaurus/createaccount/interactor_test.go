package createaccount_test

import (
	"testing"
	"github.com/luistm/testkit"
	"github.com/luistm/banksaurus/banksaurus/createaccount"
)

func TestUnitInteractorNew(t *testing.T){

	t.Run("Returns no error", func(t *testing.T) {
		_, err := createaccount.NewInteractor()
		testkit.AssertIsNil(t, err)
	})
}