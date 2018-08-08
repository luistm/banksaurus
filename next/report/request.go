package report

// NewRequest creates a new report request
func NewRequest() (*Request, error) {
	return &Request{}, nil
}

// Request model for the report interactor
type Request struct{}
