package transactions

import (
	"reflect"
	"testing"

	"github.com/luistm/go-bank-cli/lib/customerrors"
)

func TestUnitTransactionRepositoryGetAll(t *testing.T) {

	testCases := []struct {
		name      string
		output    []interface{}
		withMock  bool
		mockOuput []interface{}
	}{
		{
			name:   "Returns error if storage is not defined",
			output: []interface{}{[]*Transaction{}, customerrors.ErrInfrastructureUndefined},
		},
		{
			name:   "Returns error if infrastructure fails",
			output: []interface{}{[]*Transaction{}, customerrors.ErrInfrastructure{Msg: "Test Error"}},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		r := &repository{}
		// var m *storageMock{}

		transactions, err := r.GetAll()

		got := []interface{}{transactions, err}
		if !reflect.DeepEqual(tc.output, got) {
			t.Errorf("Expected '%v', got '%v'", tc.output, got)
		}
	}
}
