package entities

// Row is the interface SQL infrastructure rows must implement
type Row interface {
	Scan(dest ...interface{}) error
	Next() bool
}

// SQLDatabaseHandler is the interface SQL infrastructure must implement to
// be used by entity repositories
type SQLDatabaseHandler interface {
	Execute(statement string, values ...interface{}) error
	Query(statement string) (Row, error)
}

// Repository for entities
type Repository interface {
	Save(Entity) error
	Get(string) (Entity, error)
	GetAll() ([]Entity, error)
}
