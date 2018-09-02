package loadtransactions_test

import (
	"errors"
	"github.com/luistm/banksaurus/account"
	"github.com/luistm/banksaurus/money"

	"github.com/luistm/banksaurus/banksaurus/loadtransactions"

	"github.com/luistm/testkit"
	"testing"
)

type repository struct {
	lines [][]string
	err   error
}

func (r *repository) Exists(id string) error {
	if r.err != nil {
		return r.err
	}

	return nil
}

func (r *repository) NewFromLine(line []string) error {
	if r.err != nil {
		return r.err
	}
	r.lines = append(r.lines, line)

	return nil
}

type request struct {
	lines        [][]string
	accountID    string
	errLines     error
	errAccountID error
}

func (r *request) AccountID() (string, error) {
	if r.errAccountID != nil {
		return "", r.errAccountID
	}
	return r.accountID, nil
}

func (r *request) Lines() ([][]string, error) {
	if r.errLines != nil {
		return [][]string{}, r.errLines
	}
	return r.lines, nil
}

func TestUnitNewInteractor(t *testing.T) {

	t.Run("Returns error if transaction repository is undefined", func(t *testing.T) {
		_, err := loadtransactions.NewInteractor(nil, nil)
		testkit.AssertEqual(t, loadtransactions.ErrTransactionRepositoryUndefined, err)
	})

	t.Run("Returns error if account repository is undefined", func(t *testing.T) {
		_, err := loadtransactions.NewInteractor(&repository{}, nil)
		testkit.AssertEqual(t, loadtransactions.ErrAccountRepositoryUndefined, err)
	})

	t.Run("New interactor receives repository", func(t *testing.T) {
		_, err := loadtransactions.NewInteractor(&repository{}, &repository{})
		testkit.AssertIsNil(t, err)
	})
}

func TestUnitLoad(t *testing.T) {

	m, err := money.NewMoney(123)
	testkit.AssertIsNil(t, err)
	a, err := account.New("AccountID", m)
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name              string
		transactionRepo   *repository
		accountRepository *repository
		request           *request
		expectedErr       error
		expectedLines     [][]string
	}{
		{
			name:              "Loads zero transactions",
			accountRepository: &repository{},
			transactionRepo:   &repository{},
			request:           &request{},
		},
		{
			name:            "Handles request error on AccountID method",
			transactionRepo: &repository{},
			request:         &request{lines: [][]string{}, errAccountID: errors.New("test error")},
			expectedErr:     errors.New("test error"),
		},
		{
			name:              "Handles account repository error",
			accountRepository: &repository{err: errors.New("test error")},
			transactionRepo:   &repository{},
			request:           &request{},
			expectedErr:       errors.New("test error"),
		},
		{
			name:              "Loads one transaction",
			accountRepository: &repository{},
			transactionRepo:   &repository{},
			request: &request{
				lines:     [][]string{{"item1", "item2"}},
				accountID: a.ID(),
			},
			expectedLines: [][]string{{"item1", "item2"}},
		},
		{
			name:              "Loads multiple transactions",
			accountRepository: &repository{},
			transactionRepo:   &repository{},
			request: &request{
				lines:     [][]string{{"item1", "item2"}, {"item1", "item2"}},
				accountID: a.ID(),
			},
			expectedLines: [][]string{{"item1", "item2"}, {"item1", "item2"}},
		},
		{
			name:              "Handles transaction repository errors",
			accountRepository: &repository{},
			transactionRepo:   &repository{err: errors.New("test error"), lines: [][]string{}},
			request: &request{
				lines:     [][]string{{"item1", "item2"}, {"item1", "item2"}},
				accountID: a.ID(),
			},
			expectedErr:   errors.New("test error"),
			expectedLines: [][]string{},
		},
		{
			name:              "Handles request error on Lines method",
			accountRepository: &repository{},
			transactionRepo:   &repository{},
			request: &request{
				lines: [][]string{}, errLines: errors.New("test error"),
				accountID: a.ID(),
			},
			expectedErr: errors.New("test error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := loadtransactions.NewInteractor(tc.transactionRepo, tc.accountRepository)
			testkit.AssertIsNil(t, err)

			err = r.Execute(tc.request)

			testkit.AssertEqual(t, tc.expectedErr, err)
			testkit.AssertEqual(t, tc.expectedLines, tc.transactionRepo.lines)
		})
	}
}
