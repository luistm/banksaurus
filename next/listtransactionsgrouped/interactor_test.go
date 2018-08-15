package listtransactionsgrouped_test

import (
	"github.com/luistm/banksaurus/next/entity/seller"
	"github.com/luistm/banksaurus/next/entity/transaction"
	"github.com/luistm/banksaurus/next/listtransactionsgrouped"
	"github.com/luistm/testkit"
	"testing"
	"time"
)

type adapterStub struct {
	ReceivedData          []map[string]*transaction.Money
	Transactions          []*transaction.Entity
	TransactionsForSeller [][]*transaction.Entity
	callNumber            int
}

func (as *adapterStub) Factory() ([]*transaction.Entity, error) {
	return as.Transactions, nil
}

func (as *adapterStub) GetBySeller(entity *seller.Entity) ([]*transaction.Entity, error) {
	if as.callNumber == 0 {
		as.callNumber += 1
		return as.TransactionsForSeller[0], nil
	}
	return as.TransactionsForSeller[1], nil
}

func (*adapterStub) GetByID() ([]*seller.Entity, error) {
	panic("implement me")
}

func (as *adapterStub) Present(receivedData []map[string]*transaction.Money) error {
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

	// How to get a transaction with an ID?
	//factory, err := NewTransactionFactory(transactionRepository)
	//transaction ,err := factory.NewTransaction(time.Now(), "SellerID", m1)

	m1, err := transaction.NewMoney(123456789)
	testkit.AssertIsNil(t, err)
	t1, err := transaction.New(1, time.Now(), "SellerID", m1)
	testkit.AssertIsNil(t, err)

	m2, err := transaction.NewMoney(10)
	testkit.AssertIsNil(t, err)
	t2, err := transaction.New(2, time.Now(), "AnotherSellerID", m2)
	testkit.AssertIsNil(t, err)

	t3, err := transaction.New(3, time.Now(), "SellerID", m1)
	testkit.AssertIsNil(t, err)

	m1plusm3, err := t1.Value().Add(t3.Value())
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name         string
		transactions *adapterStub
		presenter    *adapterStub
		expectedErr  error
		expectedData []map[string]*transaction.Money
	}{
		{
			name:         "Returns nothing if no data available",
			transactions: &adapterStub{
				//Transactions:          []*transaction.Entity{},
				////TransactionsForSeller: *transaction.Entity{},
			},
			presenter:    &adapterStub{},
			expectedData: []map[string]*transaction.Money{},
		},
		{
			name: "Returns a single transaction",
			transactions: &adapterStub{
				Transactions:          []*transaction.Entity{t1},
				TransactionsForSeller: [][]*transaction.Entity{{t1}},
			},
			presenter:    &adapterStub{},
			expectedData: []map[string]*transaction.Money{{t1.Seller(): m1}},
		},
		{
			name: "Returns a multiple transactions",
			transactions: &adapterStub{
				Transactions:          []*transaction.Entity{t1, t2},
				TransactionsForSeller: [][]*transaction.Entity{{t1}, {t2}},
			},
			presenter: &adapterStub{},
			expectedData: []map[string]*transaction.Money{
				{t1.Seller(): t1.Value()},
				{t2.Seller(): t2.Value()},
			},
		},
		{
			name: "Returns transactions grouped by seller",
			transactions: &adapterStub{
				Transactions:          []*transaction.Entity{t1, t2, t3},
				TransactionsForSeller: [][]*transaction.Entity{{t1, t3}, {t2}},
			},
			presenter: &adapterStub{},
			expectedData: []map[string]*transaction.Money{
				{t1.Seller(): m1plusm3},
				{t2.Seller(): m2},
			},
		},

		// TODO: Test it can handle errors from repositories and presenter
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
