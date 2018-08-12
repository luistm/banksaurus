package changesellername_test

import (
	"errors"
	"github.com/luistm/banksaurus/next/changesellername"
	"github.com/luistm/banksaurus/next/entity/seller"
	"github.com/luistm/testkit"
	"testing"
)

type repository struct {
	seller         *seller.Entity
	sellerReceived *seller.Entity
	err            error
}

func (r *repository) UpdateSeller(s *seller.Entity) error {
	r.sellerReceived = s
	return nil
}

func (r *repository) GetByID(string) (*seller.Entity, error) {
	if r.err != nil {
		return &seller.Entity{}, r.err
	}
	return r.seller, nil
}

func TestUnitNewInteractor(t *testing.T) {
	t.Run("Returns no error", func(t *testing.T) {
		_, err := changesellername.NewInteractor(&repository{})
		testkit.AssertIsNil(t, err)
	})

	t.Run("Returns error if repository is undefined", func(t *testing.T) {
		_, err := changesellername.NewInteractor(nil)
		testkit.AssertEqual(t, changesellername.ErrSellerRepositoryUndefined, err)
	})
}

func TestUnitChangeSellerName(t *testing.T) {

	sellerName := "SellerName"
	s1, err := seller.New("sellerId", sellerName)
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name           string
		sellerID       string
		sellerName     string
		repository     *repository
		expectedError  error
		expectedSeller *seller.Entity
	}{
		{
			name:           "Changes a seller name",
			sellerID:       s1.ID(),
			sellerName:     sellerName,
			repository:     &repository{seller: s1},
			expectedSeller: s1,
		},
		{
			name:          "Handles repository GetByID error",
			sellerID:      s1.ID(),
			sellerName:    sellerName,
			repository:    &repository{err: errors.New("test error")},
			expectedError: errors.New("test error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			i, err := changesellername.NewInteractor(tc.repository)
			testkit.AssertIsNil(t, err)

			r, err := changesellername.NewRequest(tc.sellerID, tc.sellerName)
			testkit.AssertIsNil(t, err)

			err = i.Execute(r)

			testkit.AssertEqual(t, tc.expectedError, err)
			testkit.AssertEqual(t, tc.expectedSeller, tc.repository.sellerReceived)
		})
	}
}
