package transaction_test

import (
	"testing"

	"github.com/luistm/banksaurus/lib/seller"
	"github.com/luistm/banksaurus/lib/transaction"
	"github.com/luistm/testkit"
	"github.com/shopspring/decimal"
)

func TestUnitTransactionNew(t *testing.T) {

	s := seller.New("TheSellerSlug", "TheSellerName")
	inputValue := "1,1"
	value := "1.1"

	tr, err := transaction.NewFromString(s, inputValue)

	testkit.AssertIsNil(t, err)
	testkit.AssertEqual(t, tr.Seller, s)
	v, err := decimal.NewFromString(value)
	testkit.AssertIsNil(t, err)
	t.Log(tr.Value())
	testkit.AssertEqual(t, v.String(), tr.Value().String())

	tr, err = transaction.NewFromString(s, "")

	testkit.AssertIsNil(t, err)
	testkit.AssertEqual(t, tr.Seller, s)
	testkit.AssertEqual(t, "0", tr.Value().String())
}
