package reportgrouped

import (
	"testing"

	"github.com/luistm/banksaurus/bankservices"

	"github.com/luistm/banksaurus/elib/testkit"
	"github.com/luistm/banksaurus/banklib"
	"github.com/luistm/banksaurus/banklib/seller"
	"github.com/luistm/banksaurus/banklib/transaction"
)

func TestUnitReportFromRecodesGrouped(t *testing.T) {

	t.Log("Transactions grouped by seller")

	s1 := seller.New("Seller1Slug", "Seller1Name")
	s2 := seller.New("Seller2Slug", "Seller2Name")
	sellersFromRepository := []banklib.Entity{s1, s2}

	t1, err := transaction.New(s1, "1")
	testkit.AssertIsNil(t, err)
	t2, err := transaction.New(s1, "2")
	testkit.AssertIsNil(t, err)
	t3, err := transaction.New(s2, "1")
	testkit.AssertIsNil(t, err)
	transactionsFromRepository := []banklib.Entity{t1, t2, t3}

	summedTransaction, err := transaction.New(s1, "3")
	testkit.AssertIsNil(t, err)
	transactionsToPresenter := []banklib.Entity{summedTransaction, t3}

	transactionRepository := &banklib.RepositoryMock{}
	transactionRepository.On("GetAll").Return(transactionsFromRepository, nil)
	sellersRepository := &banklib.RepositoryMock{}
	sellersRepository.On("GetAll").Return(sellersFromRepository, nil)
	presenter := &bankservices.PresenterMock{}
	presenter.On("Present", transactionsToPresenter).Return(nil)
	i, err := New(transactionRepository, sellersRepository, presenter)
	testkit.AssertIsNil(t, err)

	err = i.Execute()

	testkit.AssertIsNil(t, err)
	transactionRepository.AssertExpectations(t)
	sellersRepository.AssertExpectations(t)
	presenter.AssertExpectations(t)
}
