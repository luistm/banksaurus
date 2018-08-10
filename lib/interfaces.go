package lib

import "fmt"

// Entity is the interface each an object must implement in order to be identified
type Entity interface {
	fmt.Stringer
	ID() string
}

// CSVHandler to handle csv files
type CSVHandler interface {
	Lines() ([][]string, error)
}

// RepositoryFetcher for entities
type RepositoryFetcher interface {
	Get(string) (Entity, error)
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

// Rows is the interface infrastructure rows must implement
type Rows interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}

// Infrastructer is the interface repositories use to read data from the infrastructure
// Implementers should match all data fields specified in the data argument.
// If not data fields are specified, fetch all data.
type Infrastructer interface {
	Get(slug string, data map[string]interface{}) (Rows, error)
}

// TODO: I want to remove these Legacy interfaces -----v

// SQLInfrastructer is the interface relational infrastructure must implement to be used by entity repositories
type SQLInfrastructer interface {
	Execute(string, ...interface{}) error
	Query(string, ...interface{}) (Rows, error)
}

// Fetcher to fetch transaction
type Fetcher interface {
	GetAll() ([]Entity, error)
}
