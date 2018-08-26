package showaccount_test

import (
	"github.com/luistm/banksaurus/banksaurus/showaccount"
	"github.com/luistm/testkit"
	"testing"
)

func TestUnitRequestNew(t *testing.T) {
	t.Run("Returns no error", func(t *testing.T) {
		_, err := showaccount.NewRequest("accountID")
		testkit.AssertIsNil(t, err)
	})

	t.Run("Returns error if id is empty", func(t *testing.T) {
		_, err := showaccount.NewRequest("")
		testkit.AssertEqual(t, showaccount.ErrInvalidAccountID, err)
	})
}

func TestUnitRequestAccountID(t *testing.T) {

	testCases := []struct {
		name              string
		expectedAccountID string
		expectedError     error
	}{
		{
			name:              "Request has account ID",
			expectedAccountID: "AccountID",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := showaccount.NewRequest(tc.expectedAccountID)
			testkit.AssertIsNil(t, err)

			accountID, err := r.AccountID()

			testkit.AssertEqual(t, tc.expectedError, err)
			testkit.AssertEqual(t, tc.expectedAccountID, accountID)
		})
	}
}
