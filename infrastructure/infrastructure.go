package infrastructure

import (
	"io"

	"github.com/luistm/banksaurus/banklib"
	"github.com/luistm/banksaurus/banklib/transaction"
)

// SQLStorage for handling SQL databases
type SQLStorage interface {
	banklib.SQLInfrastructer
	io.Closer
}

// CSVStorage for handling CSV files
type CSVStorage interface {
	transaction.CSVHandler
	io.Closer
}
