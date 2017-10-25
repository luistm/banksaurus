package infrastructure

import "github.com/luistm/go-bank-cli/entities"

// SQLStorage for handling SQL databases
type SQLStorage interface {
	entities.SQLDatabaseHandler
	Close() error
}

// CSVStorage for handling CSV files
type CSVStorage interface {
	Close() error
}
