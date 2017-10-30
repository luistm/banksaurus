package infrastructure

import (
	"io"

	"github.com/luistm/go-bank-cli/bank/reports"
	"github.com/luistm/go-bank-cli/lib"
)

// SQLStorage for handling SQL databases
type SQLStorage interface {
	lib.SQLDatabaseHandler
	io.Closer
}

// CSVStorage for handling CSV files
type CSVStorage interface {
	reports.CSVHandler
	io.Closer
}
