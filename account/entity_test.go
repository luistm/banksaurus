package account_test

import (
	"github.com/luistm/banksaurus/account"
	"github.com/luistm/banksaurus/money"
	"github.com/luistm/testkit"
	"testing"
)

func TestUnitAccountNew(t *testing.T) {
	t.Run("Returns error if invalid ID", func(t *testing.T) {
		_, err := account.New("", nil)
		testkit.AssertEqual(t, account.ErrInvalidID, err)
	})

	t.Run("Returns error if balance undefined", func(t *testing.T) {
		_, err := account.New("AccountID", nil)
		testkit.AssertEqual(t, account.ErrInvalidBalance, err)
	})
}

func TestUnitAccountBalance(t *testing.T) {

	t.Run("An account must have balance", func(t *testing.T) {
		initialBalance, err := money.NewMoney(12345)
		testkit.AssertIsNil(t, err)

		acc, err := account.New("AccountID", initialBalance)
		testkit.AssertIsNil(t, err)

		testkit.AssertEqual(t, initialBalance, acc.Balance())
	})
}
