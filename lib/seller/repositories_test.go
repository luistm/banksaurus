package seller

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luistm/banksaurus/lib"
	sqlite3 "github.com/mattn/go-sqlite3"
)

func TestUnitRepositorySave(t *testing.T) {

	seller := &Seller{slug: "Raw name", name: "Friendly name"}

	testCases := []struct {
		name            string
		input           *Seller
		output          error
		withMock        bool
		mockInput       []interface{}
		mockOutput      error
		withUpdateMock  bool
		mockInputUpdate []interface{}
	}{
		{
			name:       "Returns error if infrastructure is not defined",
			input:      &Seller{},
			output:     lib.ErrInfrastructureUndefined,
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:     "Returns error if infrastructure returns error",
			input:    seller,
			output:   &lib.ErrInfrastructure{Msg: "Test error"},
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
		{
			name:           "Updates seller if error is UNIQUE constraint",
			input:          seller,
			output:         nil,
			withMock:       true,
			withUpdateMock: true,
			mockInput: []interface{}{
				saveStatement,
				seller.slug,
				seller.name,
			},
			mockInputUpdate: []interface{}{
				updateStatement,
				seller.name,
				seller.slug,
			},
			mockOutput: sqlite3.Error{Code: sqlite3.ErrNo(19)},
		},
	}

	// TODO: To much logic in this test, maybe it's an opportunity to refactor
	for _, tc := range testCases {
		t.Log(tc.name)
		r := &Sellers{}
		var m *lib.SQLStorageMock
		if tc.withMock {
			m = new(lib.SQLStorageMock)
			if tc.withUpdateMock {
				m.On("Execute", tc.mockInput...).Return(tc.mockOutput).
					On("Execute", tc.mockInputUpdate...).Return(nil)
			} else {
				m.On("Execute", tc.mockInput...).Return(tc.mockOutput)
			}
			r.SQLStorage = m
		}

		err := r.Save(tc.input)

		if tc.withMock {
			m.AssertExpectations(t)
			if tc.withUpdateMock {
				m.AssertNumberOfCalls(t, "Execute", 2)
			}
		}
		if !reflect.DeepEqual(tc.output, err) {
			t.Errorf("Expected '%v', got '%v'", tc.output, err)
		}
	}
}
