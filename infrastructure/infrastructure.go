package infrastructure

import (
	"io"

	"github.com/luistm/banksaurus/banklib"
)

// SQLStorage for handling SQL databases
type SQLStorage interface {
	banklib.SQLInfrastructer
	io.Closer
}

// CSVStorage for handling CSV files
type CSVStorage interface {
	banklib.CSVHandler
	io.Closer
}
