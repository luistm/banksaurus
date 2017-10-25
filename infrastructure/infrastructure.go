package infrastructure

import "github.com/luistm/go-bank-cli/entities"
import "github.com/luistm/go-bank-cli/bank/reports"

// SQLStorage for handling SQL databases
type SQLStorage interface {
	entities.SQLDatabaseHandler
	Close() error
}

// CSVStorage for handling CSV files
type CSVStorage interface {
	reports.CSVHandler
	Close() error
}
