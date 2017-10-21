package entities

// Row ...
type Row interface {
	Scan(dest ...interface{}) error
	Next() bool
}

type InfrastructureHandler interface {
	Execute(statement string, values ...interface{}) error
	Query(statement string) (Row, error)
}
