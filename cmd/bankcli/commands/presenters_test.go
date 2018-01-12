package commands

import (
	"bytes"
	"testing"

	"github.com/luistm/go-bank-cli/lib"

	"github.com/luistm/go-bank-cli/elib/testkit"
)

func TestUnitCLIPresenterPresent(t *testing.T) {

	testCases := []struct {
		name     string
		output   error
		withMock bool
	}{
		{
			name:   "Returns error if output pipe is not defined",
			output: errOutputPipeUndefined,
		},
		{
			name:     "Returns expected output in pipe",
			output:   nil,
			withMock: true,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		em := &lib.EntityMock{}
		p := &CLIPresenter{}
		var ioWriterMock bytes.Buffer
		if tc.withMock {
			p.output = &ioWriterMock

			em.On("String").Return("write this to output pipe")
		}

		err := p.Present(em)

		if tc.withMock {
			testkit.AssertEqual(t, "write this to output pipe\n", ioWriterMock.String())
		}
		testkit.AssertEqual(t, tc.output, err)
	}
}
