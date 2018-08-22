package account

// New creates a new instance of account
func New() (*Entity, error) {
	return &Entity{}, nil
}

// Entity account
type Entity struct{}
