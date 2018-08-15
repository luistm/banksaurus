package transaction_test

import (
	"testing"
	"github.com/luistm/banksaurus/next/entity/transaction"
	"github.com/luistm/testkit"
)

func TestUnitNewMoney(t *testing.T){
	t.Run("Returns error if value is zero", func(t *testing.T) {
		_, err := transaction.NewMoney(0)
		testkit.AssertEqual(t, transaction.ErrInvalidMoneyValue, err)
	})
	
	t.Run("Does not return error", func(t *testing.T) {
		_, err := transaction.NewMoney(1)
		testkit.AssertIsNil(t, err)
	})
}