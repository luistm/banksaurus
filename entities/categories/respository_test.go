package categories

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luistm/go-bank-cli/entities"
	"github.com/stretchr/testify/assert"
)

func TestCategoryRepositoryGetAll(t *testing.T) {

	testCases := []struct {
		name string
	}{}

	for _, tc := range testCases {
		m := new(entities.MockSQLStorage)
		cr := repository{SQLStorage: m}
		_, err := cr.GetAll()
		assert.Error(t, err, tc.name)
	}
}

func TestUnitCategoryRepositorySave(t *testing.T) {
	category := &Category{name: "Test Category"}

	testCases := []struct {
		name      string
		input     *Category
		output    error
		withMock  bool
		mockInput []interface{}
		mockOuput error
	}{
		{
			name:      "Returns error if infrastructure not defined",
			input:     category,
			output:    entities.ErrInfrastructureUndefined,
			withMock:  false,
			mockInput: nil,
			mockOuput: nil,
		},
		{
			name:     "Returns error if infrastructure fails",
			input:    category,
			output:   &entities.ErrInfrastructure{Msg: "Test Error"},
			withMock: true,
			mockInput: []interface{}{
				insertStatement,
				category.name,
			},
			mockOuput: errors.New("Test Error"),
		},
		{
			name:     "Saves category to infrastructure",
			input:    category,
			output:   nil,
			withMock: true,
			mockInput: []interface{}{
				insertStatement,
				category.name,
			},
			mockOuput: nil,
		},
	}
	for _, tc := range testCases {
		t.Log(tc.name)
		r := repository{}
		var m *entities.MockSQLStorage
		if tc.withMock {
			m = new(entities.MockSQLStorage)
			m.On("Execute", tc.mockInput...).Return(tc.mockOuput)
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
