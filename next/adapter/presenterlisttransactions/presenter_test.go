package presenterlisttransactions_test

import (
	"github.com/luistm/banksaurus/next/adapter/presenterlisttransactions"
	"testing"

	"github.com/luistm/testkit"
)

func TestUnitPresenterPresent(t *testing.T) {

	testCases := []struct {
		name        string
		input       []map[string]int64
		callPresent bool
		output      []string
		err         error
	}{
		{
			name: "Presenter prepares prepares output data",
			input: []map[string]int64{
				{"key": 1234},
				{"key": 12},
				{"key2": -12345},
				{"key2": -12},
			},
			callPresent: true,
			output:      []string{"12,34€", "key", "0,12€", "key", "-123,45€", "key2", "-0,12€", "key2"},
		},
		{
			name:        "Presenter receives no data",
			input:       []map[string]int64{},
			callPresent: true,
			output:      []string{},
		},
		{
			name:   "Returns error if data was not presented",
			output: []string{},
			err:    presenterlisttransactions.ErrNoDataToPresent,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p, err := presenterlisttransactions.NewPresenter()
			testkit.AssertIsNil(t, err)

			if tc.callPresent {
				err = p.Present(tc.input)
				testkit.AssertIsNil(t, err)
			}

			out, err := p.OutputData()

			testkit.AssertEqual(t, tc.err, err)
			testkit.AssertEqual(t, tc.output, out)
		})
	}
}

func TestUnitPresenterViewModel(t *testing.T) {

	testCases := []struct {
		name        string
		presenter   *presenterlisttransactions.Presenter
		outputError error
	}{
		{
			name:        "Returns error if data if present has no data",
			presenter:   &presenterlisttransactions.Presenter{},
			outputError: presenterlisttransactions.ErrNoDataToPresent,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.presenter.ViewModel()

			testkit.AssertEqual(t, tc.outputError, err)
		})
	}
}
