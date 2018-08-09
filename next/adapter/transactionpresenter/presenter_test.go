package transactionpresenter_test

import (
	"github.com/luistm/banksaurus/next/adapter/transactionpresenter"
	"testing"

	"github.com/luistm/testkit"
)

func TestUnitPresenterPresent(t *testing.T) {

	testCases := []struct {
		name   string
		input  []map[string]int64
		output []string
	}{
		{
			name: "Presenter prepares prepares output data",
			input: []map[string]int64{
				{"key": 1234},
				{"key": 124},
				{"key2": 12345},
			},
			output: []string{"key", "1234", "key", "124", "key2", "12345"},
		},
		{
			name: "Presenter receives no data",
			input: []map[string]int64{},
			output: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p, err := transactionpresenter.NewPresenter()
			testkit.AssertIsNil(t, err)

			err = p.Present(tc.input)

			testkit.AssertIsNil(t, err)
			testkit.AssertEqual(t, tc.output, p.OutputData)
		})
	}
}

func TestUnitPresenterViewModel(t *testing.T){

	testCases := []struct{
		name string
		presenter *transactionpresenter.Presenter
		outputError error
	}{
		{},
	}
	// TODO: Data is not prepared
	// TODO: Data is prepared
	
	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T) {
			t.Error("Test is not finished")
		})
	}
}
