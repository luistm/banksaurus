package command

import (
	"testing"

	"github.com/luistm/banksaurus/cmd/banksaurus/command/report"
	"github.com/luistm/testkit"
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
			output: []interface{}{nil, errCommandNotFound},
		},
		{
			name:   "Returns error if empty slice received",
			input:  []string{},
			output: []interface{}{nil, errCommandIsUndefined},
		},
		{
			name:   "Returns command instance if cli input matches",
			input:  []string{"reportgrouped"},
			output: []interface{}{&report.Command{}, nil},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)

		command, err := New(tc.input)

		got := []interface{}{command, err}
		testkit.AssertEqual(t, tc.output, got)
	}
}
