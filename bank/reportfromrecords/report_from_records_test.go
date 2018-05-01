package reportfromrecords_test

import (
	"errors"
	"testing"

	"github.com/luistm/banksaurus/bank/reportfromrecords"
	"github.com/luistm/banksaurus/elib/testkit"
	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/seller"
	"github.com/luistm/banksaurus/lib/transaction"
	"github.com/luistm/banksaurus/bank"
)

func TestUnitReportFromRecordsExecute(t *testing.T) {

	sellerForTransaction := seller.New("sellerSlug", "")
	t1, err := transaction.New(sellerForTransaction, "")
	testkit.AssertIsNil(t, err)
	transactionsFromRepository := []lib.Entity{t1}

	sellersFromRepository := []lib.Entity{seller.New("sellerSlug", "TheSellerName")}
	transactionToPresenter, err := transaction.New(sellersFromRepository[0].(*seller.Seller), "")
	testkit.AssertIsNil(t, err)
	transactionsToPresenter := []lib.Entity{transactionToPresenter}

	testCases := []struct {
		name                         string
		output                       error
		transactionRepository        lib.Repository
		transactionRepositoryReturns []interface{}
		sellersRepository            lib.Repository
		sellersRepositoryReturns     []interface{}
		presenter                    bank.Presenter
		presenterReturns             error
	}{
		{
			name: "Returns nil if success",
			transactionRepository:        &lib.RepositoryMock{},
			transactionRepositoryReturns: []interface{}{transactionsFromRepository, nil},
			sellersRepository:            &lib.RepositoryMock{},
			sellersRepositoryReturns:     []interface{}{sellersFromRepository, nil},
			presenter:                    &bank.PresenterMock{},
			presenterReturns:             nil,
			output:                       nil,
		},
		{
			name: "If seller has name, fill transaction sellerForTransaction with a proper name",
			transactionRepository:        &lib.RepositoryMock{},
			transactionRepositoryReturns: []interface{}{transactionsFromRepository, nil},
			sellersRepository:            &lib.RepositoryMock{},
			sellersRepositoryReturns: []interface{}{
				sellersFromRepository,
				nil,
			},
			presenter:        &bank.PresenterMock{},
			presenterReturns: nil,
			output:           nil,
		},
		{
			name: "Handles cases where transaction do not exists",
			transactionRepository:        &lib.RepositoryMock{},
			transactionRepositoryReturns: []interface{}{[]lib.Entity{}, nil},
			sellersRepository:            &lib.RepositoryMock{},
			sellersRepositoryReturns:     []interface{}{sellersFromRepository, nil},
			presenter:                    &bank.PresenterMock{},
			presenterReturns:             nil,
			output:                       nil,
		},
		{
			name: "Returns error if transaction repository returns error",
			transactionRepository:        &lib.RepositoryMock{},
			transactionRepositoryReturns: []interface{}{[]lib.Entity{}, errors.New("test error")},
			sellersRepository:            &lib.RepositoryMock{},
			sellersRepositoryReturns:     []interface{}{[]lib.Entity{}, nil},
			presenter:                    &bank.PresenterMock{},
			presenterReturns:             nil,
			output:                       &lib.ErrRepository{Msg: "test error"},
		},
		{
			name: "Returns error if seller repository returns error",
			transactionRepository:        &lib.RepositoryMock{},
			transactionRepositoryReturns: []interface{}{transactionsFromRepository, nil},
			sellersRepository:            &lib.RepositoryMock{},
			sellersRepositoryReturns:     []interface{}{[]lib.Entity{}, errors.New("test error")},
			presenter:                    &bank.PresenterMock{},
			presenterReturns:             nil,
			output:                       &lib.ErrRepository{Msg: "test error"},
		},
		{
			name: "Returns error if presenter returns error",
			transactionRepository:        &lib.RepositoryMock{},
			transactionRepositoryReturns: []interface{}{transactionsFromRepository, nil},
			sellersRepository:            &lib.RepositoryMock{},
			sellersRepositoryReturns:     []interface{}{[]lib.Entity{}, nil},
			presenter:                    &bank.PresenterMock{},
			presenterReturns:             errors.New("test error"),
			output:                       &lib.ErrPresenter{Msg: "test error"},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)

		tc.transactionRepository.(*lib.RepositoryMock).On("GetAll").Return(tc.transactionRepositoryReturns...)
		tc.sellersRepository.(*lib.RepositoryMock).On("GetAll").Return(tc.sellersRepositoryReturns...).Maybe()
		tc.presenter.(*bank.PresenterMock).On("Present", transactionsToPresenter).Return(tc.presenterReturns).Maybe()
		i, err := reportfromrecords.New(tc.transactionRepository, tc.sellersRepository, tc.presenter)
		testkit.AssertIsNil(t, err)

		err = i.Execute()

		tc.transactionRepository.(*lib.RepositoryMock).AssertExpectations(t)
		tc.sellersRepository.(*lib.RepositoryMock).AssertExpectations(t)
		tc.presenter.(*bank.PresenterMock).AssertExpectations(t)

		testkit.AssertEqual(t, tc.output, err)
	}
}
