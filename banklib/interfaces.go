package banklib

import "fmt"

// Entity is the interface each an object must implement in order to be identified
type Entity interface {
	fmt.Stringer
	ID() string
}

// Rows is the interface SQL infrastructure rows must implement
type Rows interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}

// CSVHandler to handle csv files
type CSVHandler interface {
	Lines() ([][]string, error)
}

// SQLInfrastructer is the interface relational infrastructure must implement to be used by entity repositories
type SQLInfrastructer interface {
	Execute(string, ...interface{}) error
	Query(string, ...interface{}) (Rows, error)
}

// RepositoryFetcher for entities
type RepositoryFetcher interface {
	Get(string) (Entity, error)
	GetAll() ([]Entity, error)
}

// Fetcher to fetch transaction
type Fetcher interface {
	GetAll() ([]Entity, error)
}

// RepositoryCreator interface to create entities
type RepositoryCreator interface {
	Save(Entity) error
}

// Repository for entities
type Repository interface {
	RepositoryCreator
	RepositoryFetcher
}
