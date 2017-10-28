package transactions

import "testing"
import "reflect"

func TestUnitInteractorTransactionsLoad(t *testing.T) {

	testCases := []struct {
		name       string
		output     error
		withMock   bool
		mockOutput []*record
	}{}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := interactor{}

		err := i.Load()

		if !reflect.DeepEqual(tc.output, err) {
			t.Errorf("Expected '%v', got '%v'", tc.output, err)
		}
	}
}
