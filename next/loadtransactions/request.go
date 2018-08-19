package loadtransactions

func NewRequest(lines [][]string) (*Request, error) {
	// TODO: Should i validate data here?
	return &Request{lines}, nil
}

// Request ...
type Request struct {
	lines [][]string
}

func (r *Request) Lines() ([][]string, error) {
	return r.lines, nil
}
