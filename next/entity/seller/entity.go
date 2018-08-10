package seller

// NewFromID creates a new seller instance given it's ID
func NewFromID(id string) (*Entity, error) {
	return &Entity{id}, nil
}

// Entity for seller
type Entity struct {
	id string
}

// ID returns the id of the seller
func (e *Entity) ID() string {
	return e.id
}
