package infrastructure

import "github.com/luistm/go-bank-cli/entities"

// SQLStorage ...
type SQLStorage interface {
	entities.SQLDatabaseHandler
	Close() error
}
