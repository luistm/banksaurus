package presenterlistsellers_test

import (
	"github.com/luistm/banksaurus/cmd/bscli/adapter/presenterlistsellers"
	"github.com/luistm/testkit"
	"testing"
)

func TestUnitNewPresenter(t *testing.T) {
	t.Run("No error on new presenter", func(t *testing.T) {
		_, err := presenterlistsellers.NewPresenter()
		testkit.AssertIsNil(t, err)
	})
}

func TestUnitPresenterPresent(t *testing.T) {

	testCases := []struct {
		name          string
		input         []string
		callPresent   bool
		expectedData  []string
		expectedError error
	}{
		{
			name:         "Presenter receives data",
			input:        []string{"seller1", "seller2"},
			callPresent:  true,
			expectedData: []string{"seller1", "seller2"},
		},
		{
			name:          "Presenter returns error if not data was received",
			expectedError: presenterlistsellers.ErrNoDataToPresent,
			expectedData:  []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p, err := presenterlistsellers.NewPresenter()
			testkit.AssertIsNil(t, err)

			if tc.callPresent {
				err = p.Present(tc.input)
				testkit.AssertIsNil(t, err)
			}

			raw, err := p.OutputData()

			testkit.AssertEqual(t, tc.expectedError, err)
			testkit.AssertEqual(t, tc.expectedData, raw)
		})
	}
}

func TestUnitViewModel(t *testing.T) {

	testCases := []struct {
		name          string
		input         []string
		callPresent   bool
		expectedError error
	}{
		{
			name:        "Returns no error if present was called",
			callPresent: true,
		},
		{
			name:          "Returns  error if present was not called",
			expectedError: presenterlistsellers.ErrNoDataToPresent,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p, err := presenterlistsellers.NewPresenter()
			testkit.AssertIsNil(t, err)

			if tc.callPresent {
				err = p.Present(tc.input)
				testkit.AssertIsNil(t, err)
			}

			_, err = p.ViewModel()

			testkit.AssertEqual(t, tc.expectedError, err)
		})
	}

}
