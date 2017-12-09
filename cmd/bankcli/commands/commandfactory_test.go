package commands

import (
	"testing"

	"github.com/luistm/go-bank-cli/elib/testkit"
)

func TestUnitNewCommand(t *testing.T) {

	testCases := []struct {
		name   string
		input  []string
		output []interface{}
	}{
		{
			name:   "Returns error if command not found",
			input:  []string{"thisCommandDoesNotExist"},
			output: []interface{}{&ReportCommand{}, errCommandNotFound},
		},
		{
			name:   "Returns error if empty slice received",
			input:  []string{},
			output: []interface{}{&ReportCommand{}, errCommandIsUndefined},
		},
		{
			name:   "Returns command instance if cli input matches",
			input:  []string{"report"},
			output: []interface{}{&ReportCommand{commandType: "report"}, nil},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)

		command, err := New(tc.input)

		got := []interface{}{command, err}
		testkit.AssertEqual(t, tc.output, got)
	}
}
