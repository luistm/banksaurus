package reportfromrecordsgrouped_test

import (
	"testing"

	"github.com/luistm/banksaurus/bank/reportfromrecordsgrouped"
	"github.com/luistm/banksaurus/elib/testkit"
	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/sellers"
	"github.com/luistm/banksaurus/lib/transactions"
)

func TestUnitReportFromRecodesGrouped(t *testing.T) {

	t.Log("Transactions grouped by seller")

	s1 := sellers.New("Seller1Slug", "Seller1Name")
	s2 := sellers.New("Seller2Slug", "Seller2Name")
	sellersFromRepository := []lib.Entity{s1, s2}

	t1, err := transactions.New(s1, "1")
	testkit.AssertIsNil(t, err)
	t2, err := transactions.New(s1, "2")
	testkit.AssertIsNil(t, err)
	t3, err := transactions.New(s2, "1")
	testkit.AssertIsNil(t, err)
	transactionsFromRepository := []lib.Entity{t1, t2, t3}

	summedTransaction, err := transactions.New(s1, "3")
	testkit.AssertIsNil(t, err)
	transactionsToPresenter := []lib.Entity{summedTransaction, t3}

	transactionRepository := &lib.RepositoryMock{}
	transactionRepository.On("GetAll").Return(transactionsFromRepository, nil)
	sellersRepository := &lib.RepositoryMock{}
	sellersRepository.On("GetAll").Return(sellersFromRepository, nil)
	presenter := &lib.PresenterMock{}
	presenter.On("Present", transactionsToPresenter).Return(nil)
	i, err := reportfromrecordsgrouped.New(transactionRepository, sellersRepository, presenter)
	testkit.AssertIsNil(t, err)

	err = i.Execute()

	testkit.AssertIsNil(t, err)
	transactionRepository.AssertExpectations(t)
	sellersRepository.AssertExpectations(t)
	presenter.AssertExpectations(t)
}
