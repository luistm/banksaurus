package transactions

import (
	"errors"
	"reflect"
	"testing"

	"github.com/shopspring/decimal"

	"github.com/luistm/banksaurus/elib/testkit"

	"github.com/luistm/banksaurus/lib"

	"github.com/luistm/banksaurus/lib/customerrors"
	"github.com/luistm/banksaurus/lib/sellers"
)

func TestUnitTransactionsNew(t *testing.T) {

	t.Error("This test fails because is not finished")

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

	t1 := New(sellers.New("d1", "d1"))
	t2 := New(sellers.New("d2", "d2"))
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
			mockInput:  t1.Seller,
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
			input:  []*Transaction{New(sellers.New("SellerSlug", "SellerName"))},
			output: []lib.Entity{New(sellers.New("SellerSlug", "SellerName"))},
		},
		{
			name: "Input slice has two items, same seller",
			input: []*Transaction{
				&Transaction{value: &decimalOne, Seller: sellers.New("SellerSlug", "SellerName")},
				&Transaction{value: &decimalOne, Seller: sellers.New("SellerSlug", "SellerName")},
			},
			output: []lib.Entity{&Transaction{value: &decimalTwo, Seller: sellers.New("SellerSlug", "SellerName")}},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)

		got, _ := mergeTransactions(tc.input)

		testkit.AssertEqual(t, tc.output, got)
	}
}
