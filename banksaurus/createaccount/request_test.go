package createaccount_test

import (
	"github.com/luistm/banksaurus/banksaurus/createaccount"
	"github.com/luistm/banksaurus/money"
	"github.com/luistm/testkit"
	"testing"
)

func TestUnitRequestUnit(t *testing.T) {
	t.Run("Creates new request", func(t *testing.T) {
		_, err := createaccount.NewRequest("123,123.123")
		testkit.AssertIsNil(t, err)
	})

	t.Run("Returns error if input is empty", func(t *testing.T) {
		_, err := createaccount.NewRequest("")
		testkit.AssertEqual(t, createaccount.ErrInvalidData, err)
	})

	t.Run("Request handles unrecognized string", func(t *testing.T) {
		_, err := createaccount.NewRequest("#%SDF")
		testkit.AssertEqual(t, "strconv.ParseInt: parsing \"#%SDF\": invalid syntax", err.Error())
	})
}

func TestUnitRequestBalance(t *testing.T) {

	m1, err := money.NewMoney(1123)
	testkit.AssertIsNil(t, err)

	testsCases := []struct {
		name          string
		input         string
		expectedMoney *money.Money
		expectedError error
	}{
		{
			name:          "Request handles string with dot",
			input:         "11.23",
			expectedMoney: m1,
		},
		{
			name:          "Request handles string with comma",
			input:         "11,23",
			expectedMoney: m1,
		},
		{
			name:          "Request handles string",
			input:         "11,23",
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
