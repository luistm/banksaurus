package infrastructure

import (
	"io"

	"github.com/luistm/banksaurus/lib"
)

// SQLStorage for handling SQL databases
type SQLStorage interface {
	lib.SQLInfrastructer
	io.Closer
}

// CSVStorage for handling CSV files
type CSVStorage interface {
	lib.CSVHandler
	io.Closer
}
