package listtransactionsgrouped_test

import (
	"errors"
	"github.com/luistm/banksaurus/banksaurus/listtransactionsgrouped"
	"github.com/luistm/banksaurus/money"
	"github.com/luistm/banksaurus/seller"
	"github.com/luistm/banksaurus/transaction"
	"github.com/luistm/testkit"
	"testing"
	"time"
)

type adapterStub struct {
	ReceivedData          []map[string]*money.Money
	Transactions          []*transaction.Entity
	TransactionsForSeller [][]*transaction.Entity
	callNumber            int
	getAllError           error
	presentError          error
}

func (as *adapterStub) GetAll() ([]*transaction.Entity, error) {
	if as.getAllError != nil {
		return []*transaction.Entity{}, as.getAllError
	}
	return as.Transactions, nil
}

func (as *adapterStub) GetBySeller(entity string) ([]*transaction.Entity, error) {
	if as.callNumber == 0 {
		as.callNumber += 1
		return as.TransactionsForSeller[0], nil
	}
	return as.TransactionsForSeller[1], nil
}

func (as *adapterStub) Present(receivedData []map[string]*money.Money) error {
	if as.presentError != nil {
		return as.presentError
	}

	as.ReceivedData = receivedData
	return nil
}

func TestUnitReportGroupedNew(t *testing.T) {
	t.Run("Returns error if transactions repository is undefined", func(t *testing.T) {
		_, err := listtransactionsgrouped.NewInteractor(nil, nil)
		testkit.AssertEqual(t, listtransactionsgrouped.ErrTransactionsRepositoryUndefined, err)
	})

	t.Run("Returns error if presenter is undefined", func(t *testing.T) {
		_, err := listtransactionsgrouped.NewInteractor(&adapterStub{}, nil)
		testkit.AssertEqual(t, listtransactionsgrouped.ErrPresenterUndefined, err)
	})

	t.Run("Returns no error if repositories and presenter are defined", func(t *testing.T) {
		_, err := listtransactionsgrouped.NewInteractor(&adapterStub{}, &adapterStub{})
		testkit.AssertIsNil(t, err)
	})
}

func TestUnitReportGroupedExecute(t *testing.T) {

	s1, err := seller.New("SellerID", "")
	testkit.AssertIsNil(t, err)

	s2, err := seller.New("AnotherSellerID", "")
	testkit.AssertIsNil(t, err)

	m1, err := money.NewMoney(123456789)
	testkit.AssertIsNil(t, err)
	t1, err := transaction.New(1, time.Now(), s1, m1)
	testkit.AssertIsNil(t, err)

	m2, err := money.NewMoney(10)
	testkit.AssertIsNil(t, err)
	t2, err := transaction.New(2, time.Now(), s2, m2)
	testkit.AssertIsNil(t, err)

	t3, err := transaction.New(3, time.Now(), s1, m1)
	testkit.AssertIsNil(t, err)

	m1plusm3, err := t1.Value().Add(t3.Value())
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name         string
		transactions *adapterStub
		presenter    *adapterStub
		expectedErr  error
		expectedData []map[string]*money.Money
	}{
		{
			name:         "Returns nothing if no data available",
			transactions: &adapterStub{},
			presenter:    &adapterStub{},
			expectedData: []map[string]*money.Money{},
		},
		{
			name: "Returns a single transaction",
			transactions: &adapterStub{
				Transactions:          []*transaction.Entity{t1},
				TransactionsForSeller: [][]*transaction.Entity{{t1}},
			},
			presenter:    &adapterStub{},
			expectedData: []map[string]*money.Money{{t1.Seller().ID(): m1}},
		},
		{
			name: "Returns a multiple transactions",
			transactions: &adapterStub{
				Transactions:          []*transaction.Entity{t1, t2},
				TransactionsForSeller: [][]*transaction.Entity{{t1}, {t2}},
			},
			presenter: &adapterStub{},
			expectedData: []map[string]*money.Money{
				{t1.Seller().ID(): t1.Value()},
				{t2.Seller().ID(): t2.Value()},
			},
		},
		{
			name: "Returns transactions grouped by seller",
			transactions: &adapterStub{
				Transactions:          []*transaction.Entity{t1, t2, t3},
				TransactionsForSeller: [][]*transaction.Entity{{t1, t3}, {t2}},
			},
			presenter: &adapterStub{},
			expectedData: []map[string]*money.Money{
				{t1.Seller().ID(): m1plusm3},
				{t2.Seller().ID(): m2},
			},
		},
		{
			name: "Handles repository error",
			transactions: &adapterStub{
				Transactions: []*transaction.Entity{},
				getAllError:  errors.New("test error"),
			},
			presenter:   &adapterStub{},
			expectedErr: errors.New("test error"),
		},
		{
			name: "Handles presenter error",
			transactions: &adapterStub{
				Transactions:          []*transaction.Entity{t1},
				TransactionsForSeller: [][]*transaction.Entity{{t1}},
			},
			presenter:   &adapterStub{presentError: errors.New("test error")},
			expectedErr: errors.New("test error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := listtransactionsgrouped.NewInteractor(tc.transactions, tc.presenter)
			testkit.AssertIsNil(t, err)

			err = r.Execute()

			testkit.AssertEqual(t, tc.expectedErr, err)
			testkit.AssertEqual(t, tc.expectedData, tc.presenter.ReceivedData)
		})
	}
}
