package sellers

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luistm/go-bank-cli/lib"
	"github.com/luistm/go-bank-cli/lib/customerrors"
)

func TestUnitRepositorySave(t *testing.T) {

	seller := &Seller{slug: "Raw name", name: "Friendly name"}

	testCases := []struct {
		name       string
		input      *Seller
		output     error
		withMock   bool
		mockInput  []interface{}
		mockOutput error
	}{
		{
			name:       "Returns error if infrastructure is not defined",
			input:      &Seller{},
			output:     customerrors.ErrInfrastructureUndefined,
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:     "Returns error if infrastructure returns error",
			input:    seller,
			output:   &customerrors.ErrInfrastructure{Msg: "Test error"},
			withMock: true,
			mockInput: []interface{}{
				saveStatement,
				seller.slug,
				seller.name,
			},
			mockOutput: errors.New("Test error"),
		},
		{
			name:     "Saves seller",
			input:    seller,
			output:   nil,
			withMock: true,
			mockInput: []interface{}{
				saveStatement,
				seller.slug,
				seller.name,
			},
			mockOutput: nil,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		r := &repository{}
		var m *lib.MockSQLStorage
		if tc.withMock {
			m = new(lib.MockSQLStorage)
			m.On("Execute", tc.mockInput...).Return(tc.mockOutput)
			r.SQLStorage = m
		}

		err := r.Save(tc.input)

		if tc.withMock {
			m.AssertExpectations(t)
		}
		if !reflect.DeepEqual(tc.output, err) {
			t.Errorf("Expected '%v', got '%v'", tc.output, err)
		}
	}
}
