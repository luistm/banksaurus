package showaccount_test

import (
	"errors"
	"github.com/luistm/banksaurus/account"
	"github.com/luistm/banksaurus/banksaurus/showaccount"
	"github.com/luistm/banksaurus/money"
	"github.com/luistm/testkit"
	"testing"
)

type presenterStub struct {
	accountMoney *money.Money
	err          error
}

func (ps *presenterStub) Present(m *money.Money) error {
	if ps.err != nil {
		return ps.err
	}

	ps.accountMoney = m

	return nil
}

type repositoryStub struct {
	acc *account.Entity
	err error
}

func (rs *repositoryStub) GetByID(string) (*account.Entity, error) {
	if rs.err != nil {
		return &account.Entity{}, rs.err
	}
	return rs.acc, nil
}

type requestStub struct {
	accountID string
	err       error
}

func (rs *requestStub) AccountID() (string, error) {
	if rs.err != nil {
		return "", rs.err
	}
	return rs.accountID, nil
}

func TestUnitNewInteractor(t *testing.T) {
	t.Run("Returns no error", func(t *testing.T) {
		_, err := showaccount.NewInteractor(&presenterStub{}, &repositoryStub{})
		testkit.AssertIsNil(t, err)
	})

	t.Run("Returns error if presenter is undefined", func(t *testing.T) {
		_, err := showaccount.NewInteractor(nil, nil)
		testkit.AssertEqual(t, showaccount.ErrPresenterIsUndefined, err)
	})

	t.Run("Returns error if account repository os undefined", func(t *testing.T) {
		_, err := showaccount.NewInteractor(&presenterStub{}, nil)
		testkit.AssertEqual(t, showaccount.ErrRepositoryIsUndefined, err)
	})
}

func TestUnitShowAccountInteractorExecute(t *testing.T) {

	m1, err := money.NewMoney(1234)
	testkit.AssertIsNil(t, err)

	acc1, err := account.New(m1)
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name                 string
		presenter            *presenterStub
		repository           *repositoryStub
		request              *requestStub
		expectedAccountMoney *money.Money
		expectedError        error
	}{
		{
			name:                 "Presents an account",
			presenter:            &presenterStub{},
			repository:           &repositoryStub{acc: acc1},
			request:              &requestStub{accountID: "AccountID"},
			expectedAccountMoney: m1,
		},
		{
			name:          "Handles request error",
			presenter:     &presenterStub{},
			repository:    &repositoryStub{},
			request:       &requestStub{err: errors.New("test error")},
			expectedError: errors.New("test error"),
		},
		{
			name:          "Handles repository error",
			presenter:     &presenterStub{},
			repository:    &repositoryStub{err: errors.New("test error")},
			request:       &requestStub{accountID: "AccountID"},
			expectedError: errors.New("test error"),
		},
		{
			name:          "Handles presenter error",
			presenter:     &presenterStub{err: errors.New("test error")},
			repository:    &repositoryStub{acc: acc1},
			request:       &requestStub{accountID: "AccountID"},
			expectedError: errors.New("test error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			i, err := showaccount.NewInteractor(tc.presenter, tc.repository)
			testkit.AssertIsNil(t, err)

			err = i.Execute(tc.request)

			testkit.AssertEqual(t, tc.expectedError, err)
			testkit.AssertEqual(t, tc.expectedAccountMoney, tc.presenter.accountMoney)
		})
	}
}
