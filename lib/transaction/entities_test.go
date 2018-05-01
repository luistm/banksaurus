package transaction_test

import (
	"testing"

	"github.com/luistm/banksaurus/elib/testkit"
	"github.com/luistm/banksaurus/lib/sellers"
	"github.com/luistm/banksaurus/lib/transaction"
	"github.com/shopspring/decimal"
)

func TestUnitTransactionNew(t *testing.T) {

	s := sellers.New("TheSellerSlug", "TheSellerName")
	value := "1.1"

	tr, err := transaction.New(s, value)

	testkit.AssertIsNil(t, err)
	testkit.AssertEqual(t, tr.Seller, s)
	v, err := decimal.NewFromString(value)
	testkit.AssertIsNil(t, err)
	testkit.AssertEqual(t, tr.Value().String(), v.String())

	tr, err = transaction.New(s, "")

	testkit.AssertIsNil(t, err)
	testkit.AssertEqual(t, tr.Seller, s)
	testkit.AssertIsNil(t, err)
	testkit.AssertEqual(t, tr.Value().String(), "0")
}
