package loadtransactions_test

import (
	"errors"
	"github.com/luistm/banksaurus/banksauruslib/usecases/loadtransactions"
	"github.com/luistm/testkit"
	"testing"
)

type repository struct {
	lines [][]string
	err   error
}

func (r *repository) NewFromLine(line []string) error {
	if r.err != nil {
		return r.err
	}
	r.lines = append(r.lines, line)

	return nil
}

type request struct {
	lines [][]string
	err   error
}

func (r *request) Lines() ([][]string, error) {
	if r.err != nil {
		return [][]string{}, r.err
	}
	return r.lines, nil
}

func TestUnitNewInteractor(t *testing.T) {

	t.Run("Returns error if transaction repository is undefined", func(t *testing.T) {
		_, err := loadtransactions.NewInteractor(nil)
		testkit.AssertEqual(t, loadtransactions.ErrTransactionRepositoryUndefined, err)
	})

	t.Run("New interactor receives repository", func(t *testing.T) {
		_, err := loadtransactions.NewInteractor(&repository{})
		testkit.AssertIsNil(t, err)
	})
}

func TestUnitLoad(t *testing.T) {

	testCases := []struct {
		name            string
		transactionRepo *repository
		request         *request
		expectedErr     error
		expectedLines   [][]string
	}{
		{
			name:            "Loads zero transactions",
			transactionRepo: &repository{},
			request:         &request{},
		},
		{
			name:            "Loads one transaction",
			transactionRepo: &repository{},
			request:         &request{lines: [][]string{{"item1", "item2"}}},
			expectedLines:   [][]string{{"item1", "item2"}},
		},
		{
			name:            "Loads multiple transactions",
			transactionRepo: &repository{},
			request:         &request{lines: [][]string{{"item1", "item2"}, {"item1", "item2"}}},
			expectedLines:   [][]string{{"item1", "item2"}, {"item1", "item2"}},
		},
		{
			name:            "Handles transaction repository errors",
			transactionRepo: &repository{err: errors.New("test error"), lines: [][]string{}},
			request:         &request{lines: [][]string{{"item1", "item2"}, {"item1", "item2"}}},
			expectedErr:     errors.New("test error"),
			expectedLines:   [][]string{},
		},
		{
			name:            "Handles request error",
			transactionRepo: &repository{},
			request:         &request{lines: [][]string{}, err: errors.New("test error")},
			expectedErr:     errors.New("test error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := loadtransactions.NewInteractor(tc.transactionRepo)
			testkit.AssertIsNil(t, err)

			err = r.Execute(tc.request)

			testkit.AssertEqual(t, tc.expectedErr, err)
			testkit.AssertEqual(t, tc.expectedLines, tc.transactionRepo.lines)
		})
	}
}
