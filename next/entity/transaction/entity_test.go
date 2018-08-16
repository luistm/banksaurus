package transaction_test

import (
	"github.com/luistm/banksaurus/next/entity/transaction"
	"github.com/luistm/testkit"
	"testing"
	"time"
)

func TestUnitTransactionNew(t *testing.T) {

	m, err := transaction.NewMoney(1)
	testkit.AssertIsNil(t, err)

	t.Run("Returns error if id is not defined", func(t *testing.T) {
		_, err := transaction.New(0, time.Now(), "", m)
		testkit.AssertEqual(t, transaction.ErrInvalidTransactionID, err)
	})

	t.Run("Returns error if date is invalid", func(t *testing.T) {
		_, err := transaction.New(10, time.Time{}, "", m)
		testkit.AssertEqual(t, transaction.ErrInvalidDate, err)
	})

	t.Run("Returns error if seller ID is invalid", func(t *testing.T) {
		_, err := transaction.New(10, time.Now(), "", m)
		testkit.AssertEqual(t, transaction.ErrInvalidSeller, err)
	})

	t.Run("Returns error if value is zero", func(t *testing.T) {
		_, err := transaction.New(10, time.Now(), "SellerID", nil)
		testkit.AssertEqual(t, transaction.ErrInvalidValue, err)
	})
}
