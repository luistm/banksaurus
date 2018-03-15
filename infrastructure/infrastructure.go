package infrastructure

import (
	"io"

	"github.com/luistm/banksaurus/lib/transactions"
	"github.com/luistm/banksaurus/lib"
)

// SQLStorage for handling SQL databases
type SQLStorage interface {
	lib.SQLInfrastructer
	io.Closer
}

// CSVStorage for handling CSV files
type CSVStorage interface {
	transactions.CSVHandler
	io.Closer
}
