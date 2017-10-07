package descriptions

import "testing"
import "reflect"

func TestUnitInteractorAdd(t *testing.T) {

	testCases := []struct {
		name                string
		description         string
		expectedDescription *Description
		expectedError       bool
	}{
		{
			name:                "Returns error if repository is not defined",
			description:         "Test description",
			expectedDescription: &Description{},
			expectedError:       true,
		},
		// {
		// 	name:                "Returns error if repository fails",
		// 	description:         "Test description",
		// 	expectedDescription: &Description{},
		// 	expectedError:       true,
		// },
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := &Interactor{}

		d, err := i.Add(tc.description)

		if tc.expectedError && err == nil {
			t.Error("Was expecting and error, but got nil")
		}
		if !tc.expectedError && err != nil {
			t.Error("Was not expecting error, but got one")
		}
		if !reflect.DeepEqual(d, tc.expectedDescription) {
			t.Error("Was expecting a description, but got something else")
		}
	}

}
