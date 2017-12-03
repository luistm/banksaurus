package transactions

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luistm/go-bank-cli/elib/testkit"
	"github.com/luistm/go-bank-cli/lib/categories"

	"github.com/luistm/go-bank-cli/lib"

	"github.com/luistm/go-bank-cli/lib/customerrors"
	"github.com/luistm/go-bank-cli/lib/sellers"
	"github.com/stretchr/testify/mock"
)

type testMock struct {
	mock.Mock
}

func (m *testMock) GetAll() ([]*Transaction, error) {
	args := m.Called()
	return args.Get(0).([]*Transaction), args.Error(1)
}

func (m *testMock) Create(s string) (lib.Identifier, error) {
	args := m.Called(s)
	return args.Get(0).(lib.Identifier), args.Error(1)
}

func TestUnitInteractorTransactionsLoadDataFromRecords(t *testing.T) {

	testCasesRepository := []struct {
		name       string
		output     error
		withMock   bool
		mockOutput []interface{}
	}{
		{
			name:       "Returns error if repository is not defined",
			output:     customerrors.ErrRepositoryUndefined,
			withMock:   false,
			mockOutput: nil,
		},
		{
			name:       "Returns error on repository error",
			output:     &customerrors.ErrRepository{Msg: "Test Error"},
			withMock:   true,
			mockOutput: []interface{}{[]*Transaction{}, errors.New("Test Error")},
		},
	}

	for _, tc := range testCasesRepository {
		t.Log(tc.name)
		i := Interactor{}
		var m *testMock
		if tc.withMock {
			m = new(testMock)
			i.transactionsRepository = m
			m.On("GetAll").Return(tc.mockOutput...)
		}

		err := i.LoadDataFromRecords()

		if tc.withMock {
			m.AssertExpectations(t)
		}
		if !reflect.DeepEqual(tc.output, err) {
			t.Errorf("Expected '%v', got '%v'", tc.output, err)
		}
	}

	c := categories.New("Test Category")
	t1 := &Transaction{seller: sellers.New("d1", "d1"), c: c}
	t2 := &Transaction{seller: sellers.New("d2", "d2"), c: c}
	i := Interactor{}
	tm := new(testMock)
	i.transactionsRepository = tm
	tm.On("GetAll").Return([]*Transaction{t1, t2}, nil)

	testCasesIdentifierRepositorySave := []struct {
		name       string
		output     error
		withMock   bool
		mockInput  lib.Identifier
		mockOutput error
	}{
		{
			name:       "Returns error if entity repository is not defined",
			output:     customerrors.ErrInteractorUndefined,
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:       "Returns error if entity save method returns fail",
			output:     &customerrors.ErrInteractor{Msg: "Test Error"},
			withMock:   true,
			mockInput:  t1.seller,
			mockOutput: errors.New("Test Error"),
		},
	}

	for _, tc := range testCasesIdentifierRepositorySave {
		t.Log(tc.name)
		var im *lib.RepositoryMock
		if tc.withMock {
			im = new(lib.RepositoryMock)
			im.On("Save", tc.mockInput).Return(tc.mockOutput)
			i.sellersRepository = im
		}

		err := i.LoadDataFromRecords()

		if tc.withMock {
			im.AssertExpectations(t)
		}
		if !reflect.DeepEqual(tc.output, err) {
			t.Errorf("Expected '%v', got '%v'", tc.output, err)
		}
	}

}

type repositoryMock struct {
	mock.Mock
}

func (m *repositoryMock) GetAll() ([]*Transaction, error) {
	args := m.Called()
	return args.Get(0).([]*Transaction), args.Error(1)
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
			mockOutput: []interface{}{[]*Transaction{}, errors.New("Test Error")},
		},
		{
			name: "Report has transactions",
			output: []interface{}{
				&Report{transactions: []*Transaction{&Transaction{}}},
				nil,
			},
			withMock:   true,
			mockOutput: []interface{}{[]*Transaction{&Transaction{}}, nil},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := Interactor{}
		var m *repositoryMock
		if tc.withMock {
			m = new(repositoryMock)
			m.On("GetAll").Return(tc.mockOutput...)
			i.transactionsRepository = m

			sellersMock := new(lib.RepositoryMock)
			sellersMock.On("GetAll").Return([]lib.Identifier{}, nil)
			i.sellersRepository = sellersMock
		}

		r, err := i.ReportFromRecords()

		if tc.withMock {
			m.AssertExpectations(t)
		}
		testkit.AssertEqual(t, tc.output, []interface{}{r, err})
	}

	trMock := new(repositoryMock)
	trMock.On("GetAll").Return([]*Transaction{&Transaction{seller: sellers.New("sellerSlug", "")}}, nil)

	testCasesSellerWithPrettyName := []struct {
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
			name:       "Seller repository returns error",
			output:     []interface{}{&Report{}, &customerrors.ErrRepository{Msg: "Test error"}},
			withMock:   true,
			mockOutput: []interface{}{[]lib.Identifier{}, errors.New("Test error")},
		},
		{
			name: "Seller from transaction has no pretty name",
			output: []interface{}{
				&Report{[]*Transaction{&Transaction{seller: sellers.New("sellerSlug", "sellerName")}}},
				nil,
			},
			withMock: true,
			mockOutput: []interface{}{
				[]lib.Identifier{sellers.New("sellerSlug", "sellerName")},
				nil,
			},
		},
	}

	for _, tc := range testCasesSellerWithPrettyName {
		t.Log(tc.name)
		i := Interactor{}
		i.transactionsRepository = trMock
		var m *lib.RepositoryMock
		if tc.withMock {
			m = new(lib.RepositoryMock)
			m.On("GetAll").Return(tc.mockOutput...)
			i.sellersRepository = m
		}

		r, err := i.ReportFromRecords()

		if tc.withMock {
			m.AssertExpectations(t)
		}
		testkit.AssertEqual(t, tc.output, []interface{}{r, err})
	}

}
