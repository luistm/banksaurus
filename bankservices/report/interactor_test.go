package report

import (
	"errors"
	"testing"

	"github.com/luistm/banksaurus/bankservices"

	"github.com/luistm/banksaurus/banklib"
	"github.com/luistm/banksaurus/banklib/seller"
	"github.com/luistm/banksaurus/banklib/transaction"
	"github.com/luistm/testkit"
)

func TestUnitReportFromRecordsExecute(t *testing.T) {

	sellerForTransaction := seller.New("sellerSlug", "")
	t1, err := transaction.NewFromString(sellerForTransaction, "")
	testkit.AssertIsNil(t, err)
	transactionsFromRepository := []banklib.Entity{t1}

	sellersFromRepository := []banklib.Entity{seller.New("sellerSlug", "TheSellerName")}
	transactionToPresenter, err := transaction.NewFromString(sellersFromRepository[0].(*seller.Seller), "")
	testkit.AssertIsNil(t, err)
	transactionsToPresenter := []banklib.Entity{transactionToPresenter}

	testCases := []struct {
		name                         string
		output                       error
		transactionRepository        banklib.Repository
		transactionRepositoryReturns []interface{}
		sellersRepository            banklib.Repository
		sellersRepositoryReturns     []interface{}
		presenter                    bankservices.Presenter
		presenterReturns             error
	}{
		{
			name: "Returns nil if success",
			transactionRepository:        &banklib.RepositoryMock{},
			transactionRepositoryReturns: []interface{}{transactionsFromRepository, nil},
			sellersRepository:            &banklib.RepositoryMock{},
			sellersRepositoryReturns:     []interface{}{sellersFromRepository, nil},
			presenter:                    &bankservices.PresenterMock{},
			presenterReturns:             nil,
			output:                       nil,
		},
		{
			name: "If seller has name, fill transaction sellerForTransaction with a proper name",
			transactionRepository:        &banklib.RepositoryMock{},
			transactionRepositoryReturns: []interface{}{transactionsFromRepository, nil},
			sellersRepository:            &banklib.RepositoryMock{},
			sellersRepositoryReturns: []interface{}{
				sellersFromRepository,
				nil,
			},
			presenter:        &bankservices.PresenterMock{},
			presenterReturns: nil,
			output:           nil,
		},
		{
			name: "Handles cases where transaction do not exists",
			transactionRepository:        &banklib.RepositoryMock{},
			transactionRepositoryReturns: []interface{}{[]banklib.Entity{}, nil},
			sellersRepository:            &banklib.RepositoryMock{},
			sellersRepositoryReturns:     []interface{}{sellersFromRepository, nil},
			presenter:                    &bankservices.PresenterMock{},
			presenterReturns:             nil,
			output:                       nil,
		},
		{
			name: "Returns error if transaction repository returns error",
			transactionRepository:        &banklib.RepositoryMock{},
			transactionRepositoryReturns: []interface{}{[]banklib.Entity{}, errors.New("test error")},
			sellersRepository:            &banklib.RepositoryMock{},
			sellersRepositoryReturns:     []interface{}{[]banklib.Entity{}, nil},
			presenter:                    &bankservices.PresenterMock{},
			presenterReturns:             nil,
			output:                       &banklib.ErrRepository{Msg: "test error"},
		},
		{
			name: "Returns error if seller repository returns error",
			transactionRepository:        &banklib.RepositoryMock{},
			transactionRepositoryReturns: []interface{}{transactionsFromRepository, nil},
			sellersRepository:            &banklib.RepositoryMock{},
			sellersRepositoryReturns:     []interface{}{[]banklib.Entity{}, errors.New("test error")},
			presenter:                    &bankservices.PresenterMock{},
			presenterReturns:             nil,
			output:                       &banklib.ErrRepository{Msg: "test error"},
		},
		{
			name: "Returns error if presenter returns error",
			transactionRepository:        &banklib.RepositoryMock{},
			transactionRepositoryReturns: []interface{}{transactionsFromRepository, nil},
			sellersRepository:            &banklib.RepositoryMock{},
			sellersRepositoryReturns:     []interface{}{[]banklib.Entity{}, nil},
			presenter:                    &bankservices.PresenterMock{},
			presenterReturns:             errors.New("test error"),
			output:                       &banklib.ErrPresenter{Msg: "test error"},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)

		tc.transactionRepository.(*banklib.RepositoryMock).On("GetAll").Return(tc.transactionRepositoryReturns...)
		tc.sellersRepository.(*banklib.RepositoryMock).On("GetAll").Return(tc.sellersRepositoryReturns...).Maybe()
		tc.presenter.(*bankservices.PresenterMock).On("Present", transactionsToPresenter).Return(tc.presenterReturns).Maybe()
		i, err := New(tc.transactionRepository, tc.sellersRepository, tc.presenter)
		testkit.AssertIsNil(t, err)

		err = i.Execute()

		tc.transactionRepository.(*banklib.RepositoryMock).AssertExpectations(t)
		tc.sellersRepository.(*banklib.RepositoryMock).AssertExpectations(t)
		tc.presenter.(*bankservices.PresenterMock).AssertExpectations(t)

		testkit.AssertEqual(t, tc.output, err)
	}
}
