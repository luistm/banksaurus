package presenterlistsellers_test

import (
	"github.com/luistm/banksaurus/next/application/adapter/presenterlistsellers"
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

func TestUnitNewViewModel(t *testing.T) {
	t.Run("Does not return error", func(t *testing.T) {
		_, err := presenterlistsellers.NewViewModel([]string{})
		testkit.AssertIsNil(t, err)
	})
}

func TestUnitView(t *testing.T) {

	testCases := []struct {
		name           string
		input          []string
		expectedOutput string
		expectedError  error
	}{
		{
			name:           "Sends data to view",
			input:          []string{"a", "b"},
			expectedOutput: "a\nb\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			screen := &mockIOWriter{}

			vm, err := presenterlistsellers.NewViewModel(tc.input)
			testkit.AssertIsNil(t, err)

			err = vm.Write(screen)

			testkit.AssertEqual(t, tc.expectedError, err)
			testkit.AssertEqual(t, tc.expectedOutput, screen.received())
		})
	}
}
