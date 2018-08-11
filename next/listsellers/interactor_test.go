package listsellers_test

import (
	"errors"
	"github.com/luistm/banksaurus/next/entity/seller"
	"github.com/luistm/banksaurus/next/listsellers"
	"github.com/luistm/testkit"
	"testing"
)

type sellerRepository struct {
	sellers []*seller.Entity
	err     error
}

func (sr *sellerRepository) GetAll() ([]*seller.Entity, error) {
	if sr.err != nil {
		return []*seller.Entity{}, sr.err
	}
	return sr.sellers, nil
}

type sellerPresenter struct {
	receivedSellers []map[string]string
	error
}

func (sp *sellerPresenter) Present(sellers []map[string]string) error {
	if sp.error != nil {
		return sp.error
	}
	sp.receivedSellers = sellers
	return nil
}

func TestUnitNewInteractor(t *testing.T) {

	t.Run("Returns error if sellers repository undefined", func(t *testing.T) {
		_, err := listsellers.NewInteractor(nil, nil)
		testkit.AssertEqual(t, listsellers.ErrSellersRepositoryUndefined, err)
	})

	t.Run("Returns error if seller presenter undefined", func(t *testing.T) {
		_, err := listsellers.NewInteractor(&sellerRepository{}, nil)
		testkit.AssertEqual(t, listsellers.ErrPresenterRepositoryUndefined, err)
	})

	t.Run("Does not return error if sellers repository and presenter are defined", func(t *testing.T) {
		_, err := listsellers.NewInteractor(&sellerRepository{}, &sellerPresenter{})
		testkit.AssertIsNil(t, err)
	})
}

func TestUnitListSellers(t *testing.T) {

	s1, err := seller.NewFromID("SellerID1")
	testkit.AssertIsNil(t, err)

	s2, err := seller.NewFromID("SellerID2")
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name            string
		repository      *sellerRepository
		presenter       *sellerPresenter
		expectedSellers []map[string]string
		expectedError   error
	}{
		{
			name:            "List sellers",
			repository:      &sellerRepository{sellers: []*seller.Entity{s1, s2}},
			presenter:       &sellerPresenter{},
			expectedSellers: []map[string]string{{s1.ID(): ""}, {s2.ID(): ""}},
		},
		{
			name:            "No sellers to list",
			repository:      &sellerRepository{sellers: []*seller.Entity{}},
			presenter:       &sellerPresenter{},
			expectedSellers: []map[string]string{},
		},
		{
			name:          "Handles error from repository",
			repository:    &sellerRepository{sellers: []*seller.Entity{}, err: errors.New("test error")},
			presenter:     &sellerPresenter{},
			expectedError: errors.New("test error"),
		},
		{
			name:          "Handles presenter error",
			repository:    &sellerRepository{sellers: []*seller.Entity{s1, s2}},
			presenter:     &sellerPresenter{error: errors.New("test error")},
			expectedError: errors.New("test error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			i, err := listsellers.NewInteractor(tc.repository, tc.presenter)
			testkit.AssertIsNil(t, err)

			err = i.Execute()

			testkit.AssertEqual(t, tc.expectedError, err)
			testkit.AssertEqual(t, tc.expectedSellers, tc.presenter.receivedSellers)
		})
	}
}
