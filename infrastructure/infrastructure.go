package infrastructure

import (
	"io"

	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/transaction"
)

// SQLStorage for handling SQL databases
type SQLStorage interface {
	lib.SQLInfrastructer
	io.Closer
}

// CSVStorage for handling CSV files
type CSVStorage interface {
	transaction.CSVHandler
	io.Closer
}
