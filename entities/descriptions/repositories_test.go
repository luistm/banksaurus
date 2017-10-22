package descriptions

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luistm/go-bank-cli/entities"
)

func TestUnitRepositorySave(t *testing.T) {

	description := &Description{rawName: "Raw Name", friendlyName: "Friendly Name"}

	testCases := []struct {
		name       string
		input      *Description
		output     error
		withMock   bool
		mockInput  []interface{}
		mockOutput error
	}{
		{
			name:       "Returns error if infrastructure is not defined",
			input:      &Description{},
			output:     entities.ErrInfrastructureUndefined,
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:     "Returns error if infrastructure returns error",
			input:    description,
			output:   &entities.ErrInfrastructure{Msg: "Test error"},
			withMock: true,
			mockInput: []interface{}{
				saveStatement,
				description.rawName,
				description.friendlyName,
			},
			mockOutput: errors.New("Test error"),
		},
		{
			name:     "Saves description",
			input:    description,
			output:   nil,
			withMock: true,
			mockInput: []interface{}{
				saveStatement,
				description.rawName,
				description.friendlyName,
			},
			mockOutput: nil,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		r := &repository{}
		var m *entities.MockSQLStorage
		if tc.withMock {
			m = new(entities.MockSQLStorage)
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
