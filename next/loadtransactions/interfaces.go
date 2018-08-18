package loadtransactions

// TransactionGateway is a collection of entities to be used by the report
type TransactionGateway interface {
	NewFromLine([]string) error
}

// Request to load transactions
type Request interface {
	Lines() ([][]string, error)
}
