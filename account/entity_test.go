package account_test

import (
	"github.com/luistm/banksaurus/account"
	"github.com/luistm/banksaurus/money"
	"github.com/luistm/testkit"
	"testing"
)

func TestUnitAccountBalance(t *testing.T) {

	t.Run("An account has balance", func(t *testing.T) {
		initialBalance, err := money.NewMoney(12345)
		testkit.AssertIsNil(t, err)

		acc, err := account.New(initialBalance)
		testkit.AssertIsNil(t, err)

		testkit.AssertEqual(t, initialBalance, acc.Balance())
	})
}
