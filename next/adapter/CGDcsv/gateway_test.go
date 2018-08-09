package CGDcsv_test

import (
	"github.com/luistm/banksaurus/next/adapter/CGDcsv"
	"github.com/luistm/banksaurus/next/entity/transaction"
	"github.com/luistm/testkit"
	"testing"
	"time"
)

func TestUnitNewGateway(t *testing.T) {

	testCases := []struct {
		name   string
		input  [][]string
		output []*transaction.Entity
		err    error
	}{
		{
			name:  "Returns error if line number does not match",
			input: [][]string{},
			err:   CGDcsv.ErrInvalidNumberOfLines,
		},
		{
			name:  "Expects 8 lines",
			input: [][]string{{}, {}, {}, {}, {}, {}, {}, {}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := CGDcsv.New(tc.input)
			testkit.AssertEqual(t, tc.err, err)
		})
	}
}

func TestUnitReturnTransactions(t *testing.T) {

	date, err := time.Parse("02-01-2006", "25-10-2017")
	testkit.AssertIsNil(t, err)
	t1, err := transaction.New(date, "COMPRA CONTINENTE MAI ", -7752)
	testkit.AssertIsNil(t, err)
	t2, err := transaction.New(date, "COMPRA CONTINENTE ", 7752)
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name   string
		input  [][]string
		output []*transaction.Entity
		err    error
	}{
		{
			name: "Returns transactions",
			input: [][]string{{}, {}, {}, {}, {},
				{"25-10-2017", "25-10-2017", "COMPRA CONTINENTE MAI ", "77,52", "", "61,25", "61.25"},
				{"25-10-2017", "25-10-2017", "COMPRA CONTINENTE ", "", "77,52", "61,25", "61.25"},
				{}, {}},
			output: []*transaction.Entity{t1, t2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := CGDcsv.New(tc.input)
			testkit.AssertIsNil(t, err)

			ts, err := r.GetAll()

			testkit.AssertEqual(t, tc.err, err)
			testkit.AssertEqual(t, len(tc.output), len(ts))
			testkit.AssertEqual(t, tc.output, ts)
		})
	}
}
