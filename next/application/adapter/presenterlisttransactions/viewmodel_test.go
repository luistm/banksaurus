package presenterlisttransactions_test

import (
	"github.com/luistm/banksaurus/next/application/adapter/presenterlisttransactions"
	"github.com/luistm/testkit"
	"testing"
)

type mockIOWriter struct {
	data []byte
}

func (miow *mockIOWriter) Write(p []byte) (n int, err error) {
	miow.data = append(miow.data, p...)
	return len(p), nil
}

func (miow *mockIOWriter) received() string {
	return string(miow.data)
}

func TestUnitNewVieModel(t *testing.T) {

	t.Run("Returns error if data has not an event length", func(t *testing.T) {
		_, err := presenterlisttransactions.NewViewModel([]string{"key", "1234", "key2"})
		testkit.AssertEqual(t, presenterlisttransactions.ErrDataHasOddLength, err)
	})

	t.Run("Returns no error if data has zero length", func(t *testing.T) {
		_, err := presenterlisttransactions.NewViewModel([]string{})
		testkit.AssertIsNil(t, err)
	})
}

func TestUnitViewModel(t *testing.T) {

	testCases := []struct {
		name   string
		input  []string
		output string
	}{
		{
			name:   "View model",
			input:  []string{"key", "1234", "key2", "12345"},
			output: "key  1234\nkey2 12345\n",
		},
		{
			name:   "View model has not data",
			input:  []string{},
			output: "",
		},
		{
			name:   "View model empty slice",
			input:  []string{},
			output: "",
		},
		{
			name:   "View model no input",
			output: "",
		},
	}

	// TODO: fix this! vm.View is not handling errors

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			screen := &mockIOWriter{}

			vm, err := presenterlisttransactions.NewViewModel(tc.input)
			testkit.AssertIsNil(t, err)

			vm.Write(screen)

			testkit.AssertEqual(t, tc.output, screen.received())
		})
	}
}
