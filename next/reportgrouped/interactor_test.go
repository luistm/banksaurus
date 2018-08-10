package reportgrouped_test

import (
	"github.com/luistm/banksaurus/next/entity/seller"
	"github.com/luistm/banksaurus/next/entity/transaction"
	"github.com/luistm/banksaurus/next/reportgrouped"
	"github.com/luistm/testkit"
	"testing"
	"time"
)

type adapterStub struct {
	ReceivedData          []map[string]int64
	Transactions          []*transaction.Entity
	TransactionsForSeller [][]*transaction.Entity
	callNumber            int
}

func (as *adapterStub) GetAll() ([]*transaction.Entity, error) {
	return as.Transactions, nil
}

func (as *adapterStub) GetBySellerID(entity *seller.Entity) ([]*transaction.Entity, error) {
	if as.callNumber == 0 {
		as.callNumber += 1
		return as.TransactionsForSeller[0], nil
	}
	return as.TransactionsForSeller[1], nil
}

func (*adapterStub) GetByID() ([]*seller.Entity, error) {
	panic("implement me")
}

func (as *adapterStub) Present(receivedData []map[string]int64) error {
	as.ReceivedData = receivedData
	return nil
}

func TestUnitReportGroupedNew(t *testing.T) {
	t.Run("Returns error if transactions repository is undefined", func(t *testing.T) {
		_, err := reportgrouped.NewInteractor(nil, nil, nil)
		testkit.AssertEqual(t, reportgrouped.ErrTransactionsRepositoryUndefined, err)
	})

	t.Run("Returns error if presenter is undefined", func(t *testing.T) {
		_, err := reportgrouped.NewInteractor(&adapterStub{}, &adapterStub{}, nil)
		testkit.AssertEqual(t, reportgrouped.ErrPresenterUndefined, err)
	})

	t.Run("Returns error if sellers repository is undefined", func(t *testing.T) {
		_, err := reportgrouped.NewInteractor(&adapterStub{}, nil, &adapterStub{})
		testkit.AssertEqual(t, reportgrouped.ErrSellersRepositoryUndefined, err)
	})

	t.Run("Returns no error if repositories and presenter are defined", func(t *testing.T) {
		_, err := reportgrouped.NewInteractor(&adapterStub{}, &adapterStub{}, &adapterStub{})
		testkit.AssertIsNil(t, err)
	})
}

func TestUnitReportGroupedExecute(t *testing.T) {

	t1, err := transaction.New(time.Now(), "SellerID", 123456789)
	testkit.AssertIsNil(t, err)
	t2, err := transaction.New(time.Now(), "AnotherSellerID", 10)
	testkit.AssertIsNil(t, err)
	t3, err := transaction.New(time.Now(), "SellerID", 123456789)
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name         string
		transactions *adapterStub
		sellers      *adapterStub
		presenter    *adapterStub
		expectedErr  error
		expectedData []map[string]int64
	}{
		{
			name:         "Returns nothing if no data available",
			transactions: &adapterStub{},
			sellers:      &adapterStub{},
			presenter:    &adapterStub{},
			expectedData: []map[string]int64{},
		},
		{
			name: "Returns a single transaction",
			transactions: &adapterStub{
				Transactions:          []*transaction.Entity{t1},
				TransactionsForSeller: [][]*transaction.Entity{{t1}},
			},
			sellers:   &adapterStub{},
			presenter: &adapterStub{},
			expectedData: []map[string]int64{
				{t1.Seller(): t1.Value()},
			},
		},
		{
			name: "Returns a multiple transactions",
			transactions: &adapterStub{
				Transactions:          []*transaction.Entity{t1, t2},
				TransactionsForSeller: [][]*transaction.Entity{{t1}, {t2}},
			},
			sellers:   &adapterStub{},
			presenter: &adapterStub{},
			expectedData: []map[string]int64{
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
			sellers:   &adapterStub{},
			presenter: &adapterStub{},
			expectedData: []map[string]int64{
				{t1.Seller(): t1.Value() + t3.Value()},
				{t2.Seller(): t2.Value()},
			},
		},

		// TODO: Test it can handle errors from repositories and presenter
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := reportgrouped.NewInteractor(tc.transactions, tc.sellers, tc.presenter)
			testkit.AssertIsNil(t, err)

			err = r.Execute()

			testkit.AssertEqual(t, tc.expectedErr, err)
			testkit.AssertEqual(t, tc.expectedData, tc.presenter.ReceivedData)
		})
	}
}
