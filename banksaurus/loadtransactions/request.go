package loadtransactions

// NewRequest creates an instance of a request
func NewRequest(lines [][]string) (*Request, error) {
	return &Request{lines}, nil
}

// Request for load transactions
type Request struct {
	lines [][]string
}

func (r *Request) AccountID() (string, error) {
	panic("implement me")
}

// Lines ...
func (r *Request) Lines() ([][]string, error) {
	return r.lines, nil
}
