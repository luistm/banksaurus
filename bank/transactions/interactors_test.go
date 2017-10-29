package transactions

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luistm/go-bank-cli/entities"
)

func TestUnitInteractorTransactionsLoad(t *testing.T) {

	testCases := []struct {
		name       string
		output     error
		withMock   bool
		mockOutput []interface{}
	}{
		{
			name:       "Returns error on repository error",
			output:     &entities.ErrRepository{Msg: "Test Error"},
			withMock:   true,
			mockOutput: []interface{}{[]*record{}, errors.New("Test Error")},
		},
		// {
		// 	name: "Creates description",
		// }
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := interactor{}

		err := i.Load()

		if !reflect.DeepEqual(tc.output, err) {
			t.Errorf("Expected '%v', got '%v'", tc.output, err)
		}
	}
}
