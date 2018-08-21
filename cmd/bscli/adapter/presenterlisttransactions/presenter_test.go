package presenterlisttransactions_test

import (
	"testing"

	"github.com/luistm/banksaurus/cmd/bscli/adapter/presenterlisttransactions"
	"github.com/luistm/banksaurus/transaction"
	"github.com/luistm/testkit"
)

func TestUnitPresenterPresent(t *testing.T) {

	m1, err := transaction.NewMoney(1234)
	testkit.AssertIsNil(t, err)
	m2, err := transaction.NewMoney(12)
	testkit.AssertIsNil(t, err)
	m3, err := transaction.NewMoney(-12345)
	testkit.AssertIsNil(t, err)
	m4, err := transaction.NewMoney(-12)
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name        string
		input       []map[string]*transaction.Money
		callPresent bool
		output      []string
		err         error
	}{
		{
			name: "Presenter prepares prepares output data",
			input: []map[string]*transaction.Money{
				{"key": m1},
				{"key": m2},
				{"key2": m3},
				{"key2": m4},
			},
			callPresent: true,
			output:      []string{"12,34€", "key", "0,12€", "key", "-123,45€", "key2", "-0,12€", "key2"},
		},
		{
			name:        "Presenter receives no data",
			input:       []map[string]*transaction.Money{},
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
