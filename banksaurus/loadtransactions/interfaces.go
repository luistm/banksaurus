package loadtransactions

// TransactionGateway is a collection of entities to be used by the report
type TransactionGateway interface {
	NewFromLine([]string) error
}

// RequestLoadTransactions to load transactions
type RequestLoadTransactions interface {
	Lines() ([][]string, error)
}
