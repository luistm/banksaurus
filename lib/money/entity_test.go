package money_test

import (
	"testing"

	"github.com/luistm/banksaurus/lib/money"
	"github.com/luistm/testkit"
)

func TestUnitMoneyAmountNew(t *testing.T) {

	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{"Empty string is zero", "", "0"},
		{"Zero is zero", "0", "0"},
		{"One is one", "1", "1"},
		{"1,1 is 1.1", "1,1", "1.1"},
		{"0,0001 is 0.0001", "0,0001", "0.0001"},
	}

	for _, tc := range testCases {
		t.Log(tc.name)

		ma, err := money.New(tc.input)

		testkit.AssertIsNil(t, err)
		testkit.AssertEqual(t, tc.output, ma.ToDecimal().String())
	}
}
