package changesellername_test

import (
	"github.com/luistm/banksaurus/next/changesellername"
	"github.com/luistm/banksaurus/next/entity/seller"
	"github.com/luistm/testkit"
	"testing"
)

type repository struct{
	seller *seller.Entity
}

func (r *repository) UpdateSeller(*seller.Entity) error {
	panic("implement me")
}

func (r *repository) GetByID(string) (*seller.Entity, error) {
	return r.seller, nil
}

func (r *repository) updatedSeller() *seller.Entity {
	return nil
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

	sellerName := "SellerID"
	s1, err := seller.New("sellerId", sellerName)
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name           string
		input string
		repository     *repository
		expectedError  error
		expectedSeller *seller.Entity
	}{
		{
			name:           "Changes a seller name",
			input: sellerName,
			repository:     &repository{},
			expectedSeller: s1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			i, err := changesellername.NewInteractor(tc.repository)
			testkit.AssertIsNil(t, err)

			r , err := changesellername.NewRequest(tc.input)
			testkit.AssertIsNil(t, err)

			err = i.Execute(r)

			testkit.AssertEqual(t, tc.expectedError, err)
			testkit.AssertEqual(t, tc.expectedSeller, tc.repository.updatedSeller())
		})
	}
}
