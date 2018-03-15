package reportfromrecords_test

import (
	"github.com/luistm/banksaurus/bank/reportfromrecords"
	"github.com/luistm/banksaurus/elib/testkit"
	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/sellers"
	"github.com/luistm/banksaurus/lib/transactions"
	"testing"
	"errors"
	"github.com/luistm/banksaurus/lib/customerrors"
)

func TestUnitReportFromRecordsExecute(t *testing.T) {

	sellerForTransaction := sellers.New("sellerSlug", "")
	transaction := transactions.New(sellerForTransaction)
	transactionsFromRepository := []lib.Entity{transaction}

	sellersFromRepository := []lib.Entity{sellers.New("sellerSlug", "TheSellerName")}
	transactionsToPresenter := []lib.Entity{transactions.New(sellersFromRepository[0].(*sellers.Seller))}

	testCases := []struct {
		name                         string
		output                       error
		transactionRepository        lib.Repository
		transactionRepositoryReturns []interface{}
		sellersRepository            lib.Repository
		sellersRepositoryReturns     []interface{}
		presenter                    lib.Presenter
		presenterReturns             error
	}{
		{
			name: "Returns nil if success",
			transactionRepository:        &lib.RepositoryMock{},
			transactionRepositoryReturns: []interface{}{transactionsFromRepository, nil},
			sellersRepository:            &lib.RepositoryMock{},
			sellersRepositoryReturns:     []interface{}{sellersFromRepository, nil},
			presenter:                    &lib.PresenterMock{},
			presenterReturns:             nil,
			output:                       nil,
		},
		{
			name: "If seller has name, fill transaction sellerForTransaction with a proper name",
			transactionRepository:        &lib.RepositoryMock{},
			transactionRepositoryReturns: []interface{}{transactionsFromRepository, nil},
			sellersRepository:            &lib.RepositoryMock{},
			sellersRepositoryReturns:     []interface{}{
				sellersFromRepository,
				nil,
			},
			presenter:                    &lib.PresenterMock{},
			presenterReturns:             nil,
			output:                       nil,
		},
		{
			name: "Returns error if transaction repository returns error",
			transactionRepository:        &lib.RepositoryMock{},
			transactionRepositoryReturns: []interface{}{[]lib.Entity{}, errors.New("test error")},
			sellersRepository:            &lib.RepositoryMock{},
			sellersRepositoryReturns:     []interface{}{[]lib.Entity{}, nil},
			presenter:                    &lib.PresenterMock{},
			presenterReturns:             nil,
			output:                       &customerrors.ErrRepository{Msg: "test error"},
		},
		{
			name: "Returns error if seller repository returns error",
			transactionRepository:        &lib.RepositoryMock{},
			transactionRepositoryReturns: []interface{}{transactionsFromRepository, nil},
			sellersRepository:            &lib.RepositoryMock{},
			sellersRepositoryReturns:     []interface{}{[]lib.Entity{}, errors.New("test error")},
			presenter:                    &lib.PresenterMock{},
			presenterReturns:             nil,
			output:                       &customerrors.ErrRepository{Msg: "test error"},
		},
		{
			name: "Returns error if presenter returns error",
			transactionRepository:        &lib.RepositoryMock{},
			transactionRepositoryReturns: []interface{}{transactionsFromRepository, nil},
			sellersRepository:            &lib.RepositoryMock{},
			sellersRepositoryReturns:     []interface{}{[]lib.Entity{}, nil},
			presenter:                    &lib.PresenterMock{},
			presenterReturns:             errors.New("test error"),
			output:                       &customerrors.ErrPresenter{Msg: "test error"},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)

		tc.transactionRepository.(*lib.RepositoryMock).On("GetAll").Return(tc.transactionRepositoryReturns...)
		tc.sellersRepository.(*lib.RepositoryMock).On("GetAll").Return(tc.sellersRepositoryReturns...).Maybe()
		tc.presenter.(*lib.PresenterMock).On("Present", transactionsToPresenter).Return(tc.presenterReturns).Maybe()
		i, err := reportfromrecords.New(tc.transactionRepository, tc.sellersRepository, tc.presenter)
		testkit.AssertIsNil(t, err)

		err = i.Execute()

		tc.transactionRepository.(*lib.RepositoryMock).AssertExpectations(t)
		tc.sellersRepository.(*lib.RepositoryMock).AssertExpectations(t)
		tc.presenter.(*lib.PresenterMock).AssertExpectations(t)

		testkit.AssertEqual(t, tc.output, err)
	}
}
