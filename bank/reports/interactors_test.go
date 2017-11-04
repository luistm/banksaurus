package reports

import (
	"testing"

	"github.com/luistm/go-bank-cli/bank/transactions"

	"github.com/luistm/go-bank-cli/elib/testkit"
	"github.com/luistm/go-bank-cli/lib/customerrors"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (m *repositoryMock) GetAll() ([]*transactions.Transaction, error) {
	args := m.Called()
	return args.Get(0).([]*transactions.Transaction), args.Error(1)
}

func TestUnitReport(t *testing.T) {

	testCases := []struct {
		name     string
		output   []interface{}
		withMock bool
	}{
		{
			name:   "Returns error if repository is indefined",
			output: []interface{}{&Report{}, customerrors.ErrRepositoryUndefined},
		},
		{
			name:     "Returns report",
			output:   []interface{}{&Report{}, nil},
			withMock: true,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := interactor{}
		var m *repositoryMock
		if tc.withMock {
			m = new(repositoryMock)
			i.repository = m
		}

		r, err := i.Report()

		if tc.withMock {
			m.AssertExpectations(t)
		}
		testkit.AssertEqual(t, tc.output, []interface{}{r, err})
	}
}
