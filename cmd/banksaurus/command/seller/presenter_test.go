package seller

import (
	"bytes"
	"testing"

	"github.com/luistm/banksaurus/banklib"

	"github.com/luistm/testkit"
)

func TestUnitCLIPresenterPresent(t *testing.T) {

	testCases := []struct {
		name         string
		output       error
		withMock     bool
		entityOutput string
	}{
		{
			name:   "Returns error if output pipe is not defined",
			output: errOutputPipeUndefined,
		},
		{
			name:         "Returns expected output in pipe",
			output:       nil,
			withMock:     true,
			entityOutput: "write this to output pipe",
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		em := &banklib.EntityMock{}
		p := &CLIPresenter{}
		var ioWriterMock bytes.Buffer
		if tc.withMock {
			p.output = &ioWriterMock
			em.On("String").Return(tc.entityOutput)
		}

		err := p.Present(em)

		if tc.withMock {
			testkit.AssertEqual(t, tc.entityOutput+"\n", ioWriterMock.String())
		}
		testkit.AssertEqual(t, tc.output, err)
	}
}
