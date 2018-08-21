package loadtransactions

// TransactionGateway is a collection of entities to be used by the report
type TransactionGateway interface {
	NewFromLine([]string) error
}

// Requester to load transactions
type Requester interface {
	Lines() ([][]string, error)
}
