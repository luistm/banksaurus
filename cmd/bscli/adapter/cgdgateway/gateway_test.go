package cgdgateway_test

import (
	"github.com/luistm/banksaurus/cmd/bscli/adapter/cgdgateway"
	"github.com/luistm/banksaurus/money"
	"github.com/luistm/banksaurus/seller"
	"github.com/luistm/banksaurus/transaction"
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
			err:   cgdgateway.ErrInvalidNumberOfLines,
		},
		{
			name:  "Expects 8 lines",
			input: [][]string{{}, {}, {}, {}, {}, {}, {}, {}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := cgdgateway.New(tc.input)
			testkit.AssertEqual(t, tc.err, err)
		})
	}
}

func TestUnitGetAll(t *testing.T) {

	date, err := time.Parse("02-01-2006", "25-10-2017")
	testkit.AssertIsNil(t, err)

	s1, err := seller.New("COMPRA CONTINENTE MAI", "")
	testkit.AssertIsNil(t, err)

	s2, err := seller.New("COMPRA CONTINENTE", "")
	testkit.AssertIsNil(t, err)

	m1, err := money.NewMoney(-7752)
	testkit.AssertIsNil(t, err)
	t1, err := transaction.New(1, date, s1, m1)
	testkit.AssertIsNil(t, err)
	m2, err := money.NewMoney(7752)
	testkit.AssertIsNil(t, err)
	t2, err := transaction.New(1, date, s2, m2)
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
			r, err := cgdgateway.New(tc.input)
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
	s1, err := seller.New(sellerID, "")
	testkit.AssertIsNil(t, err)

	date, err := time.Parse("02-01-2006", "25-10-2017")
	testkit.AssertIsNil(t, err)

	m1, err := money.NewMoney(7752)
	testkit.AssertIsNil(t, err)
	t2, err := transaction.New(1, date, s1, m1)
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
			r, err := cgdgateway.New(tc.input)
			testkit.AssertIsNil(t, err)

			ts, err := r.GetBySeller(tc.seller.ID())

			testkit.AssertEqual(t, tc.err, err)
			testkit.AssertEqual(t, len(tc.output), len(ts))
			testkit.AssertEqual(t, tc.output, ts)
		})
	}
}
