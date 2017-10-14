package infrastructure

import "github.com/luistm/go-bank-cli/lib/categories"

// Storage ...
type Storage interface {
	Execute(statement string, values ...interface{}) error
	Query(statement string) (categories.IRow, error)
	Close() error
}
