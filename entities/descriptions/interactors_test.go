package descriptions

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luistm/go-bank-cli/entities"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (m *repositoryMock) Save(c *Description) error {
	args := m.Called(c)
	return args.Error(0)
}

func TestUnitInteractorAdd(t *testing.T) {

	var description = "TestDescrition"

	testCases := []struct {
		name       string
		input      string
		output     []interface{}
		withMock   bool
		mockInput  *Description
		mockOutput error
	}{
		{
			name:       "Returns error if repository is not defined",
			input:      description,
			output:     []interface{}{&Description{}, entities.ErrRepositoryIsNil},
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:       "Returns error if description is empty string",
			input:      "",
			output:     []interface{}{&Description{}, entities.ErrBadInput},
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:  "Returns error on repository error",
			input: description,
			output: []interface{}{
				&Description{},
				&entities.ErrRepository{Msg: "Test Error"},
			},
			withMock:   true,
			mockInput:  &Description{rawName: description},
			mockOutput: errors.New("Test Error"),
		},
		{
			name:  "Returns description entity created",
			input: description,
			output: []interface{}{
				&Description{rawName: description},
				nil,
			},
			withMock:   true,
			mockInput:  &Description{rawName: description},
			mockOutput: nil,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := &Interactor{}
		var m *repositoryMock
		if tc.withMock {
			m = new(repositoryMock)
			m.On("Save", tc.mockInput).Return(tc.mockOutput)
			i.Repository = m
		}

		d, err := i.Add(tc.input)

		if tc.withMock {
			m.AssertExpectations(t)
		}
		got := []interface{}{d, err}
		if !reflect.DeepEqual(tc.output, got) {
			t.Errorf("Expected %v, got %v", tc.output, got)
		}

	}

}
