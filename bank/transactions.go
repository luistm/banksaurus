package bank


// New creates a transaction show interactor
func New() (Interactor, error) {
	return &TransactionsShow{}, nil
}

// TransactionsShow shows transaction
type TransactionsShow struct{}

// Execute ...
func (ts *TransactionsShow) Execute() error {
	return nil
}
