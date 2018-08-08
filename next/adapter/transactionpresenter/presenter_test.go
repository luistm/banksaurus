package transactionpresenter_test

import (
	"github.com/luistm/banksaurus/next/adapter/transactionpresenter"
	"testing"

	"github.com/luistm/testkit"
)

func TestUnitPresenterViewModel(t *testing.T) {

	testCases := []struct {
		name   string
		input  []map[string]int64
		output string
	}{
		{
			name: "View model",
			input: []map[string]int64{
				{"key": 1234},
				{"key2": 12345},
			},
			output: " key 1234 \nkey2 12345\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p, err := transactionpresenter.NewPresenter()
			testkit.AssertIsNil(t, err)

			err = p.Present(tc.input)
			testkit.AssertIsNil(t, err)

			vm, err := p.ViewModel()

			testkit.AssertEqual(t, tc.output, vm)
		})
	}
}
