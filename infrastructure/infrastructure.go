package infrastructure

import (
	"io"

	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/transactions"
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
