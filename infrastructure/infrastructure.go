package infrastructure

import (
	"io"

	"github.com/luistm/go-bank-cli/bank"
	"github.com/luistm/go-bank-cli/lib"
)

// SQLStorage for handling SQL databases
type SQLStorage interface {
	lib.SQLInfrastructer
	io.Closer
}

// CSVStorage for handling CSV files
type CSVStorage interface {
	bank.CSVHandler
	io.Closer
}
