package createaccount_test

import (
	"github.com/luistm/banksaurus/banksaurus/createaccount"
	"github.com/luistm/banksaurus/money"
	"github.com/luistm/testkit"
	"testing"
)

func TestUnitRequestUnit(t *testing.T) {
	t.Run("Creates new request", func(t *testing.T) {
		_, err := createaccount.NewRequest(1)
		testkit.AssertIsNil(t, err)
	})
}

func TestUnitRequestBalance(t *testing.T) {

	m1, err := money.NewMoney(1)
	testkit.AssertIsNil(t, err)

	testsCases := []struct {
		name          string
		input         int64
		expectedMoney *money.Money
		expectedError error
	}{
		{
			name:          "Balance returns money",
			input:         m1.Value(),
			expectedMoney: m1,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := createaccount.NewRequest(tc.input)
			testkit.AssertIsNil(t, err)

			m, err := r.Balance()

			testkit.AssertEqual(t, tc.expectedError, err)
			testkit.AssertEqual(t, tc.expectedMoney, m)
		})
	}
}
