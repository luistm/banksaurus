package load_test

import (
	"github.com/luistm/banksaurus/next/entity/seller"
	"github.com/luistm/banksaurus/next/entity/transaction"
	"github.com/luistm/banksaurus/next/load"
	"github.com/luistm/testkit"
	"testing"
	"time"
)

type repository struct {
	Transactions []*transaction.Entity
}

func (r *repository) GetAll() ([]*transaction.Entity, error) {
	return r.Transactions, nil
}

type sellerRepository struct {
	Sellers      []*seller.Entity
	SavedSellers []*seller.Entity
}

func (r *sellerRepository) Save(s *seller.Entity) error {
	r.SavedSellers = append(r.SavedSellers, s)
	return nil
}

func TestUnitNewInteractor(t *testing.T) {

	t.Run("Returns error if transaction repository is undefined", func(t *testing.T) {
		_, err := load.NewInteractor(nil, &sellerRepository{})
		testkit.AssertEqual(t, load.ErrTransactionRepositoryUndefined, err)
	})

	t.Run("Returns error if seller repository is undefined", func(t *testing.T) {
		_, err := load.NewInteractor(&repository{}, nil)
		testkit.AssertEqual(t, load.ErrSellerRepositoryUndefined, err)
	})

	t.Run("New interactor receives repository", func(t *testing.T) {
		_, err := load.NewInteractor(&repository{}, &sellerRepository{})
		testkit.AssertIsNil(t, err)
	})
}

func TestUnitLoad(t *testing.T) {

	s1, err := seller.NewFromID("SellerID")
	testkit.AssertIsNil(t, err)
	t1, err := transaction.New(time.Now(), s1.ID(), 123456)
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

		// TODO: Test error paths
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := load.NewInteractor(tc.transactionRepo, tc.sellerRepo)
			testkit.AssertIsNil(t, err)

			err = r.Execute()

			testkit.AssertEqual(t, tc.expectedErr, err)
			testkit.AssertEqual(t, tc.expectedSellers, tc.sellerRepo.SavedSellers)
		})
	}
}
