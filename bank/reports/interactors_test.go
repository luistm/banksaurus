package reports

import (
	"errors"
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
		name       string
		output     []interface{}
		withMock   bool
		mockOutput []interface{}
	}{
		{
			name:   "Returns error if repository is indefined",
			output: []interface{}{&Report{}, customerrors.ErrRepositoryUndefined},
		},
		{
			name:       "Returns error if repository returns error",
			output:     []interface{}{&Report{}, &customerrors.ErrRepository{Msg: "Test Error"}},
			withMock:   true,
			mockOutput: []interface{}{[]*transactions.Transaction{}, errors.New("Test Error")},
		},
		{
			name: "Report has transactions",
			output: []interface{}{
				&Report{transactions: []*transactions.Transaction{&transactions.Transaction{}}},
				nil,
			},
			withMock:   true,
			mockOutput: []interface{}{[]*transactions.Transaction{&transactions.Transaction{}}, nil},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := interactor{}
		var m *repositoryMock
		if tc.withMock {
			m = new(repositoryMock)
			m.On("GetAll").Return(tc.mockOutput...)
			i.repository = m
		}

		r, err := i.ReportFromRecords()

		if tc.withMock {
			m.AssertExpectations(t)
		}
		testkit.AssertEqual(t, tc.output, []interface{}{r, err})
	}
}
