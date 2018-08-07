package cgd_csv_test

import (
	"github.com/luistm/banksaurus/next/entity/transaction"
	"github.com/luistm/banksaurus/next/gateway/cgd_csv"
	"github.com/luistm/testkit"
	"testing"
	"time"
)

func TestUnitNewGateway(t *testing.T) {

	timeNow := time.Now()
	t1, err := transaction.New(timeNow, "", 1)
	testkit.AssertIsNil(t, err)
	t2, err := transaction.New(timeNow, "", 1)
	testkit.AssertIsNil(t, err)
	t3, err := transaction.New(timeNow, "", 1)
	testkit.AssertIsNil(t, err)
	t4, err := transaction.New(timeNow, "", 1)
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name   string
		input  string
		output []*transaction.Entity
		err    error
	}{
		{
			name:   "Returns transactions",
			input:  "../../../data/fixtures/sample_records_load.csv",
			output: []*transaction.Entity{t1, t2, t3, t4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := cgd_csv.New(tc.input)
			testkit.AssertIsNil(t, err)

			ts, err := r.GetAll()

			testkit.AssertEqual(t, tc.err, err)
			testkit.AssertEqual(t, len(tc.output), len(ts))
		})
	}
}
