package loadtransactions_test

import (
	"github.com/luistm/banksaurus/banksaurus/loadtransactions"
	"github.com/luistm/testkit"
	"testing"
)

func TestUnitRequestNew(t *testing.T) {
	t.Run("Creates new request", func(t *testing.T) {
		_, err := loadtransactions.NewRequest([][]string{}, "ThisIsTheAccountID")
		testkit.AssertIsNil(t, err)
	})
}


func TestUnitRequestAccountID(t *testing.T){

	testCases := []struct{
		name string
		inputAccountID string
		expectedID string
		expectedErr error
	}{
		{
			name: "Request has account ID",
		},
	}

	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T) {
			r, err := loadtransactions.NewRequest([][]string{}, tc.inputAccountID)

			id, err := r.AccountID()

			testkit.AssertEqual(t, tc.expectedErr, err)
			testkit.AssertEqual(t, tc.expectedID, id)
		})
	}
}