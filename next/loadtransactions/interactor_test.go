package loadtransactions_test

import (
	"errors"
	"github.com/luistm/banksaurus/next/entity/seller"
	"github.com/luistm/banksaurus/next/entity/transaction"
	"github.com/luistm/banksaurus/next/loadtransactions"
	"github.com/luistm/testkit"
	"testing"
	"time"
)

type repository struct {
	Transactions []*transaction.Entity
	Error        error
}

func (r *repository) Factory() ([]*transaction.Entity, error) {
	if r.Error != nil {
		return r.Transactions, r.Error
	}
	return r.Transactions, nil
}

type sellerRepository struct {
	Sellers      []*seller.Entity
	SavedSellers []*seller.Entity
	Error        error
}

func (r *sellerRepository) Save(s *seller.Entity) error {
	if r.Error != nil {
		return r.Error
	}

	r.SavedSellers = append(r.SavedSellers, s)

	return nil
}

func TestUnitNewInteractor(t *testing.T) {

	t.Run("Returns error if transaction repository is undefined", func(t *testing.T) {
		_, err := loadtransactions.NewInteractor(nil, &sellerRepository{})
		testkit.AssertEqual(t, loadtransactions.ErrTransactionRepositoryUndefined, err)
	})

	t.Run("Returns error if seller repository is undefined", func(t *testing.T) {
		_, err := loadtransactions.NewInteractor(&repository{}, nil)
		testkit.AssertEqual(t, loadtransactions.ErrSellerRepositoryUndefined, err)
	})

	t.Run("New interactor receives repository", func(t *testing.T) {
		_, err := loadtransactions.NewInteractor(&repository{}, &sellerRepository{})
		testkit.AssertIsNil(t, err)
	})
}

func TestUnitLoad(t *testing.T) {

	s1, err := seller.New("SellerID", "")
	testkit.AssertIsNil(t, err)
	m1, err := transaction.NewMoney(123456)
	testkit.AssertIsNil(t, err)
	t1, err := transaction.New(1, time.Now(), s1.ID(), m1)
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name            string
		transactionRepo *repository
		sellerRepo      *sellerRepository
		expectedErr     error
		expectedSellers []*seller.Entity
	}{
		{
			name:            "Zero transactions to save",
			transactionRepo: &repository{},
			sellerRepo:      &sellerRepository{},
		},
		{
			name:            "Load creates new sellers",
			transactionRepo: &repository{Transactions: []*transaction.Entity{t1}},
			sellerRepo:      &sellerRepository{Sellers: []*seller.Entity{s1}},
			expectedSellers: []*seller.Entity{s1},
		},
		{
			name:            "Handles transaction repository errors",
			transactionRepo: &repository{Transactions: []*transaction.Entity{}, Error: errors.New("test error")},
			expectedErr:     errors.New("test error"),
			sellerRepo:      &sellerRepository{},
		},
		{
			name:            "Handles seller repository errors",
			transactionRepo: &repository{Transactions: []*transaction.Entity{t1}},
			expectedErr:     errors.New("test error"),
			sellerRepo:      &sellerRepository{Error: errors.New("test error")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := loadtransactions.NewInteractor(tc.transactionRepo, tc.sellerRepo)
			testkit.AssertIsNil(t, err)

			err = r.Execute()

			testkit.AssertEqual(t, tc.expectedErr, err)
			testkit.AssertEqual(t, tc.expectedSellers, tc.sellerRepo.SavedSellers)
		})
	}
}
