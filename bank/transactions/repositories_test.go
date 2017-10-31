package transactions

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luistm/go-bank-cli/lib/customerrors"
	"github.com/luistm/go-bank-cli/lib/sellers"
	"github.com/stretchr/testify/mock"
)

type storageMock struct {
	mock.Mock
}

func (s *storageMock) Lines() ([][]string, error) {
	args := s.Called()
	return args.Get(0).([][]string), args.Error(1)
}

func TestUnitTransactionRepositoryGetAll(t *testing.T) {

	testCases := []struct {
		name       string
		output     []interface{}
		withMock   bool
		mockOutput []interface{}
	}{
		{
			name:       "Returns error if storage is not defined",
			output:     []interface{}{[]*Transaction{}, customerrors.ErrInfrastructureUndefined},
			withMock:   false,
			mockOutput: nil,
		},
		{
			name:       "Returns error if infrastructure fails",
			output:     []interface{}{[]*Transaction{}, &customerrors.ErrInfrastructure{Msg: "Test Error"}},
			withMock:   true,
			mockOutput: []interface{}{[][]string{}, errors.New("Test Error")},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		r := &repository{}
		var m *storageMock
		if tc.withMock {
			m = new(storageMock)
			m.On("Lines").Return(tc.mockOutput...)
			r.storage = m
		}

		transactions, err := r.GetAll()

		got := []interface{}{transactions, err}
		if tc.withMock {
			m.AssertExpectations(t)
		}
		if !reflect.DeepEqual(tc.output, got) {
			t.Errorf("Expected '%v', got '%v'", tc.output, got)
		}
	}
}

func TestUnitTransactionRepositoryBuildTransactions(t *testing.T) {

	testCases := []struct {
		name                 string
		input                [][]string
		output               error
		expectedTransactions []*Transaction
	}{
		{
			name:   "Parses a single line",
			input:  [][]string{[]string{"25-10-2017", "25-10-2017", "COMPRA CAFETARIA HEAR", "4,30", "", "233,86", "233,86"}},
			output: nil,
			expectedTransactions: []*Transaction{
				&Transaction{
					s: sellers.New("COMPRA CAFETARIA HEAR", "COMPRA CAFETARIA HEAR"),
				},
			},
		},
	}

	for _, tc := range testCases {
		r := &repository{}

		err := r.buildTransactions(tc.input)

		if !reflect.DeepEqual(tc.output, err) {
			t.Errorf("Expected '%v', got '%v'", tc.output, err)
		}
		if !reflect.DeepEqual(tc.expectedTransactions, r.transactions) {
			t.Errorf("Expected '%v', got '%v'", tc.expectedTransactions, r.transactions)
		}
	}
}
