package transactions

import (
	"errors"
	"reflect"
	"testing"

	"github.com/shopspring/decimal"

	"github.com/luistm/go-bank-cli/elib/testkit"
	"github.com/luistm/go-bank-cli/lib/categories"

	"github.com/luistm/go-bank-cli/lib"

	"github.com/luistm/go-bank-cli/lib/customerrors"
	"github.com/luistm/go-bank-cli/lib/sellers"
)

// type testMock struct {
// 	mock.Mock
// }

// func (m *testMock) GetAll() ([]lib.Identifier, error) {
// 	args := m.Called()
// 	return args.Get(0).([]lib.Identifier), args.Error(1)
// }

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
			mockOutput: []interface{}{[]lib.Identifier{}, errors.New("Test Error")},
		},
	}

	for _, tc := range testCasesRepository {
		t.Log(tc.name)
		i := Interactor{}
		var m *lib.RepositoryMock
		if tc.withMock {
			m = new(lib.RepositoryMock)
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
	t1 := &Transaction{seller: sellers.New("d1", "d1"), category: c}
	t2 := &Transaction{seller: sellers.New("d2", "d2"), category: c}
	i := Interactor{}
	tm := new(lib.RepositoryMock)
	i.transactionsRepository = tm
	tm.On("GetAll").Return([]lib.Identifier{t1, t2}, nil)

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

// type repositoryMock struct {
// 	mock.Mock
// }

// func (m *repositoryMock) GetAll() ([]*Transaction, error) {
// 	args := m.Called()
// 	return args.Get(0).([]*Transaction), args.Error(1)
// }

func TestUnitReportFromRecords(t *testing.T) {

	presenterMock := new(lib.PresenterMock)

	testCases := []struct {
		name       string
		output     error
		withMock   bool
		mockOutput []interface{}
	}{
		{
			name:   "Returns error if repository is undefined",
			output: customerrors.ErrRepositoryUndefined,
		},
		{
			name:       "Returns error if repository returns error",
			output:     &customerrors.ErrRepository{Msg: "Test Error"},
			withMock:   true,
			mockOutput: []interface{}{[]lib.Identifier{}, errors.New("Test Error")},
		},
		{
			name:       "Returns nil if success",
			output:     nil,
			withMock:   true,
			mockOutput: []interface{}{[]lib.Identifier{&Transaction{}}, nil},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := NewInteractor(nil, nil, presenterMock)
		i.presenter = presenterMock
		var m *lib.RepositoryMock
		if tc.withMock {
			m = new(lib.RepositoryMock)
			m.On("GetAll").Return(tc.mockOutput...)
			i.transactionsRepository = m

			sellersMock := new(lib.RepositoryMock)
			sellersMock.On("GetAll").Return([]lib.Identifier{}, nil)
			i.sellersRepository = sellersMock
		}

		err := i.ReportFromRecords()

		if tc.withMock {
			m.AssertExpectations(t)
		}
		testkit.AssertEqual(t, tc.output, err)
	}

	trMock := new(lib.RepositoryMock)
	trMock.On("GetAll").Return([]lib.Identifier{&Transaction{seller: sellers.New("sellerSlug", "")}}, nil)

	testCasesSellerWithPrettyName := []struct {
		name       string
		output     error
		withMock   bool
		mockOutput []interface{}
	}{
		{
			name:   "Returns error if repository is undefined",
			output: customerrors.ErrRepositoryUndefined,
		},
		{
			name:       "Seller repository returns error",
			output:     &customerrors.ErrRepository{Msg: "Test error"},
			withMock:   true,
			mockOutput: []interface{}{[]lib.Identifier{}, errors.New("Test error")},
		},
		{
			name:     "Seller from transaction has no pretty name",
			output:   nil,
			withMock: true,
			mockOutput: []interface{}{
				[]lib.Identifier{sellers.New("sellerSlug", "sellerName")},
				nil,
			},
		},
	}

	for _, tc := range testCasesSellerWithPrettyName {
		t.Log(tc.name)
		i := NewInteractor(trMock, nil, presenterMock)
		var m *lib.RepositoryMock
		if tc.withMock {
			m = new(lib.RepositoryMock)
			m.On("GetAll").Return(tc.mockOutput...)
			i.sellersRepository = m
		}

		err := i.ReportFromRecords()

		if tc.withMock {
			m.AssertExpectations(t)
		}
		testkit.AssertEqual(t, tc.output, err)
	}

	// Define seller repository mock
	sellersMock := new(lib.RepositoryMock)
	sellersMock.On("GetAll").Return([]lib.Identifier{sellers.New("sellerSlug", "sellerName")}, nil)

	testCasesPresenter := []struct {
		name       string
		output     error
		withMock   bool
		mockInput  *lib.Identifier
		mockOutput error
	}{
		{
			name:   "Returns error if presenter is not defined",
			output: customerrors.ErrPresenterUndefined,
		},
	}

	for _, tc := range testCasesPresenter {
		t.Log(tc.name)
		i := NewInteractor(trMock, sellersMock, nil)

		err := i.ReportFromRecords()

		testkit.AssertEqual(t, tc.output, err)
	}
}

func TestUnitMergeTransactionsWithSameSeller(t *testing.T) {

	decimalOne, _ := decimal.NewFromString("1")
	decimalTwo, _ := decimal.NewFromString("2")

	testCases := []struct {
		name   string
		input  []*Transaction
		output []*Transaction
	}{
		{
			name:   "Input slice is empty",
			input:  []*Transaction{},
			output: []*Transaction{},
		},
		{
			name:   "Input slice has one item, whitout seller",
			input:  []*Transaction{&Transaction{}},
			output: []*Transaction{},
		},
		{
			name:   "Input slice has one item, with seller",
			input:  []*Transaction{&Transaction{seller: sellers.New("SellerSlug", "SellerName")}},
			output: []*Transaction{&Transaction{seller: sellers.New("SellerSlug", "SellerName")}},
		},
		{
			name: "Input slice has two items, same seller",
			input: []*Transaction{
				&Transaction{value: &decimalOne, seller: sellers.New("SellerSlug", "SellerName")},
				&Transaction{value: &decimalOne, seller: sellers.New("SellerSlug", "SellerName")},
			},
			output: []*Transaction{&Transaction{value: &decimalTwo, seller: sellers.New("SellerSlug", "SellerName")}},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)

		got, _ := mergeTransactions(tc.input)

		testkit.AssertEqual(t, tc.output, got)
	}
}
