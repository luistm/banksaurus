package accountpresenter_test

import (
	"github.com/luistm/banksaurus/cmd/bscli/adapter/accountpresenter"
	"github.com/luistm/banksaurus/money"
	"github.com/luistm/testkit"
	"strconv"
	"testing"
)

func TestUnitNewPresenter(t *testing.T) {
	t.Run("Returns no error", func(t *testing.T) {
		_, err := accountpresenter.New()
		testkit.AssertIsNil(t, err)
	})
}

func TestUnitPresenter(t *testing.T) {

	m1, err := money.NewMoney(12345)
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name          string
		input         *money.Money
		callPresent   bool
		expectedView  *accountpresenter.ViewModel
		expectedError error
	}{
		{
			name:         "Presenter creates view model",
			callPresent:  true,
			input:        m1,
			expectedView: &accountpresenter.ViewModel{Value: strconv.FormatInt(m1.Value(), 10)},
		},
		{
			name:          "View model returns error if presenter did not received data",
			expectedView:  &accountpresenter.ViewModel{},
			expectedError: accountpresenter.ErrNoDataAvailable,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p, err := accountpresenter.New()
			testkit.AssertIsNil(t, err)

			if tc.callPresent {
				err = p.Present(tc.input)
				testkit.AssertIsNil(t, err)
			}
			vm, err := p.ViewModel()

			testkit.AssertEqual(t, tc.expectedError, err)
			testkit.AssertEqual(t, tc.expectedView, vm)
		})
	}
}
