package cgdcsv_test

import (
	"github.com/luistm/banksaurus/next/application/adapter/cgdcsv"
	"github.com/luistm/banksaurus/next/entity/seller"
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
			err:   cgdcsv.ErrInvalidNumberOfLines,
		},
		{
			name:  "Expects 8 lines",
			input: [][]string{{}, {}, {}, {}, {}, {}, {}, {}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := cgdcsv.New(tc.input)
			testkit.AssertEqual(t, tc.err, err)
		})
	}
}

func TestUnitGetAll(t *testing.T) {

	date, err := time.Parse("02-01-2006", "25-10-2017")
	testkit.AssertIsNil(t, err)
	t1, err := transaction.New(date, "COMPRA CONTINENTE MAI", -7752)
	testkit.AssertIsNil(t, err)
	t2, err := transaction.New(date, "COMPRA CONTINENTE", 7752)
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
			r, err := cgdcsv.New(tc.input)
			testkit.AssertIsNil(t, err)

			ts, err := r.GetAll()

			testkit.AssertEqual(t, tc.err, err)
			testkit.AssertEqual(t, len(tc.output), len(ts))
			testkit.AssertEqual(t, tc.output, ts)
		})
	}
}

func TestUnitReturnGetBySeller(t *testing.T) {

	sellerID := "COMPRA CONTINENTE"

	date, err := time.Parse("02-01-2006", "25-10-2017")
	testkit.AssertIsNil(t, err)
	t2, err := transaction.New(date, sellerID, 7752)
	testkit.AssertIsNil(t, err)

	s1, err := seller.NewFromID(sellerID)
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name   string
		input  [][]string
		output []*transaction.Entity
		seller *seller.Entity
		err    error
	}{
		{
			name: "Returns transactions",
			input: [][]string{{}, {}, {}, {}, {},
				{"25-10-2017", "25-10-2017", "COMPRA CONTINENTE MAI ", "77,52", "", "61,25", "61.25"},
				{"25-10-2017", "25-10-2017", s1.ID(), "", "77,52", "61,25", "61.25"},
				{}, {}},
			seller: s1,
			output: []*transaction.Entity{t2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := cgdcsv.New(tc.input)
			testkit.AssertIsNil(t, err)

			ts, err := r.GetBySeller(tc.seller)

			testkit.AssertEqual(t, tc.err, err)
			testkit.AssertEqual(t, len(tc.output), len(ts))
			testkit.AssertEqual(t, tc.output, ts)
		})
	}
}
