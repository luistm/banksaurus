package transactions

import (
	"errors"
	"reflect"
	"testing"

	"github.com/shopspring/decimal"

	"github.com/luistm/banksaurus/elib/testkit"
	"github.com/luistm/banksaurus/lib/categories"

	"github.com/luistm/banksaurus/lib"

	"github.com/luistm/banksaurus/lib/customerrors"
	"github.com/luistm/banksaurus/lib/sellers"
)

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
			mockOutput: []interface{}{[]lib.Entity{}, errors.New("Test Error")},
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
	tm.On("GetAll").Return([]lib.Entity{t1, t2}, nil)

	testCasesEntityRepositorySave := []struct {
		name       string
		output     error
		withMock   bool
		mockInput  lib.Entity
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

	for _, tc := range testCasesEntityRepositorySave {
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

func TestUnitReportFromRecords(t *testing.T) {

	seller := sellers.New("sellerSlug", "sellerName")
	transactions := []lib.Entity{&Transaction{seller: seller}}
	presenterMock := new(lib.PresenterMock)
	presenterMock.On("Present", transactions).Return(nil)

	sellersMock := new(lib.RepositoryMock)
	sellersMock.On("GetAll").Return([]lib.Entity{seller}, nil)

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
			mockOutput: []interface{}{[]lib.Entity{}, errors.New("Test Error")},
		},
		{
			name:       "Returns nil if success",
			output:     nil,
			withMock:   true,
			mockOutput: []interface{}{transactions, nil},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := NewInteractor(nil, sellersMock, presenterMock)
		var m *lib.RepositoryMock
		if tc.withMock {
			m = new(lib.RepositoryMock)
			m.On("GetAll").Return(tc.mockOutput...)
			i.transactionsRepository = m
		}

		err := i.ReportFromRecords()

		if tc.withMock {
			m.AssertExpectations(t)
		}
		testkit.AssertEqual(t, tc.output, err)
	}

	trMock := new(lib.RepositoryMock)
	trMock.On("GetAll").Return([]lib.Entity{&Transaction{seller: sellers.New("sellerSlug", "")}}, nil)

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
			mockOutput: []interface{}{[]lib.Entity{}, errors.New("Test error")},
		},
		{
			name:     "Seller from transaction has no pretty name",
			output:   nil,
			withMock: true,
			mockOutput: []interface{}{
				[]lib.Entity{sellers.New("sellerSlug", "sellerName")},
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

	testCasesPresenter := []struct {
		name       string
		output     error
		withMock   bool
		mockInput  []lib.Entity
		mockOutput error
	}{
		{
			name:   "Returns error if presenter is not defined",
			output: customerrors.ErrPresenterUndefined,
		},
		{
			name:       "Returns error if presenter returns error",
			output:     &customerrors.ErrPresenter{Msg: "test error"},
			withMock:   true,
			mockInput:  transactions,
			mockOutput: errors.New("test error"),
		},
		{
			name:      "Returns nil on presenter success",
			withMock:  true,
			mockInput: transactions,
		},
	}

	for _, tc := range testCasesPresenter {
		t.Log(tc.name)
		i := NewInteractor(trMock, sellersMock, nil)
		var pm *lib.PresenterMock
		if tc.withMock {
			pm = new(lib.PresenterMock)
			pm.On("Present", tc.mockInput).Return(tc.mockOutput)
			i.presenter = pm
		}

		err := i.ReportFromRecords()

		if tc.withMock {
			pm.AssertExpectations(t)
		}
		testkit.AssertEqual(t, tc.output, err)
	}
}

func TestUnitMergeTransactionsWithSameSeller(t *testing.T) {

	decimalOne, _ := decimal.NewFromString("1")
	decimalTwo, _ := decimal.NewFromString("2")

	testCases := []struct {
		name   string
		input  []*Transaction
		output []lib.Entity
	}{
		{
			name:   "Input slice is empty",
			input:  []*Transaction{},
			output: []lib.Entity{},
		},
		{
			name:   "Input slice has one item, whitout seller",
			input:  []*Transaction{&Transaction{}},
			output: []lib.Entity{},
		},
		{
			name:   "Input slice has one item, with seller",
			input:  []*Transaction{&Transaction{seller: sellers.New("SellerSlug", "SellerName")}},
			output: []lib.Entity{&Transaction{seller: sellers.New("SellerSlug", "SellerName")}},
		},
		{
			name: "Input slice has two items, same seller",
			input: []*Transaction{
				&Transaction{value: &decimalOne, seller: sellers.New("SellerSlug", "SellerName")},
				&Transaction{value: &decimalOne, seller: sellers.New("SellerSlug", "SellerName")},
			},
			output: []lib.Entity{&Transaction{value: &decimalTwo, seller: sellers.New("SellerSlug", "SellerName")}},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)

		got, _ := mergeTransactions(tc.input)

		testkit.AssertEqual(t, tc.output, got)
	}
}
