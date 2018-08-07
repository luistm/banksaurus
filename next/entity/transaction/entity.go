package transaction

// New creates a new transaction
func New() (*Entity, error) {
	return &Entity{}, nil
}

// Entity represents a transaction
type Entity struct {
	id uint64
}

// ID returns the identification of the transaction
func (t *Entity) ID() uint64 {
	return t.id
}
