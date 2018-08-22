package createaccount_test

import (
	"errors"
	"github.com/luistm/banksaurus/account"
	"github.com/luistm/banksaurus/banksaurus/createaccount"
	"github.com/luistm/banksaurus/money"
	"github.com/luistm/testkit"
	"testing"
)

type repositoryStub struct {
	err error
	acc *account.Entity
}

func (rs *repositoryStub) New(balance *money.Money) (*account.Entity, error) {
	if rs.err != nil {
		return &account.Entity{}, rs.err
	}

	acc, err := account.New(balance)
	if err != nil {
		return &account.Entity{}, err
	}

	rs.acc = acc

	return rs.acc, nil
}

type requestStub struct {
	money *money.Money
	err   error
}

func (rs *requestStub) Balance() (*money.Money, error) {
	if rs.err != nil {
		return &money.Money{}, rs.err
	}
	return rs.money, nil
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

	m1, err := money.NewMoney(124)
	testkit.AssertIsNil(t, err)
	ac1, err := account.New(m1)
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name            string
		repository      *repositoryStub
		request         *requestStub
		expectedAccount *account.Entity
		expectedError   error
	}{
		{
			name:            "Creates a new account",
			repository:      &repositoryStub{},
			request:         &requestStub{money: m1},
			expectedAccount: ac1,
		},
		{
			name:          "Handles repository error",
			repository:    &repositoryStub{err: errors.New("test error")},
			request:       &requestStub{money: m1},
			expectedError: errors.New("test error"),
		},
		{
			name:          "Handles request error",
			repository:    &repositoryStub{},
			request:       &requestStub{err: errors.New("test error")},
			expectedError: errors.New("test error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			i, err := createaccount.NewInteractor(tc.repository)
			testkit.AssertIsNil(t, err)

			err = i.Execute(tc.request)

			testkit.AssertEqual(t, tc.expectedError, err)
			testkit.AssertEqual(t, tc.expectedAccount, tc.repository.acc)
		})
	}
}
