package createaccount_test

import (
	"github.com/luistm/banksaurus/account"
	"github.com/luistm/banksaurus/banksaurus/createaccount"
	"github.com/luistm/testkit"
	"testing"
)

type repositoryStub struct {
	acc *account.Entity
}

func (rs *repositoryStub) New() (*account.Entity, error) {
	return rs.acc, nil
}

func TestUnitInteractorNew(t *testing.T) {

	t.Run("Returns no error", func(t *testing.T) {
		_, err := createaccount.NewInteractor(&repositoryStub{})
		testkit.AssertIsNil(t, err)
	})

	t.Run("Returns error if repository is undefined", func(t *testing.T) {
		_, err := createaccount.NewInteractor(nil)
		testkit.AssertEqual(t, createaccount.ErrRepositoryUndefined, err)
	})

}

func TestUnitInteractorExecute(t *testing.T) {

	ac1, err := account.New()
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name            string
		repository      *repositoryStub
		expectedAccount *account.Entity
		expectedError   error
	}{
		{
			name:            "Creates a new account",
			repository:      &repositoryStub{acc: ac1},
			expectedAccount: ac1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			i, err := createaccount.NewInteractor(tc.repository)
			testkit.AssertIsNil(t, err)

			err = i.Execute()

			testkit.AssertEqual(t, tc.expectedError, err)
			testkit.AssertEqual(t, tc.expectedAccount, tc.repository.acc)
		})
	}
}
