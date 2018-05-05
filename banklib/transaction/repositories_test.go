package transaction

import (
	"errors"
	"testing"

	"github.com/luistm/testkit"

	"github.com/luistm/banksaurus/banklib"

	"github.com/luistm/banksaurus/banklib/seller"
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
			output:     []interface{}{[]banklib.Entity{}, banklib.ErrInfrastructureUndefined},
			withMock:   false,
			mockOutput: nil,
		},
		{
			name:       "Returns error if infrastructure fails",
			output:     []interface{}{[]banklib.Entity{}, &banklib.ErrInfrastructure{Msg: "test error"}},
			withMock:   true,
			mockOutput: []interface{}{[][]string{}, errors.New("test error")},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		r := &Transactions{}
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
		testkit.AssertEqual(t, tc.output, got)
	}
}

func TestUnitTransactionRepositoryBuildTransactions(t *testing.T) {

	value, _ := decimalFromStringWithComma("4,30")

	testCases := []struct {
		name                 string
		input                [][]string
		output               error
		expectedTransactions []banklib.Entity
	}{
		{
			name:   "Parses a single line",
			input:  [][]string{{"25-10-2017", "25-10-2017", "COMPRA CAFETARIA HEAR", "4,30", "", "233,86", "233,86"}},
			output: nil,
			expectedTransactions: []banklib.Entity{
				&Transaction{
					value:  &value,
					Seller: seller.New("COMPRA CAFETARIA HEAR", "COMPRA CAFETARIA HEAR"),
				},
			},
		},
	}

	for _, tc := range testCases {
		r := &Transactions{}

		err := r.buildTransactions(tc.input)

		testkit.AssertEqual(t, tc.output, err)
		testkit.AssertEqual(t, tc.expectedTransactions, r.transactions)
	}
}
