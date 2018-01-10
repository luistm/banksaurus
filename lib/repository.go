package lib

// Rows is the interface SQL infrastructure rows must implement
type Rows interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}

// SQLInfrastructer is the interface SQL infrastructure must implement to
// be used by entity repositories
type SQLInfrastructer interface {
	Execute(statement string, values ...interface{}) error
	Query(statement string, args ...interface{}) (Rows, error)
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
