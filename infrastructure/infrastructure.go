package infrastructure

import "github.com/luistm/go-bank-cli/lib/categories"

// Storage ...
type Storage interface {
	categories.IDBHandler
	Close() error
}
