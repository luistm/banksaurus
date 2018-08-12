package seller

// NewFromID creates a new seller instance given it's ID
func NewFromID(id string) (*Entity, error) {
	return &Entity{id: id}, nil
}

func New(id string, name string) (*Entity, error) {
	return &Entity{id, name}, nil
}

// Entity for seller
type Entity struct {
	id   string
	name string
}

func (e *Entity) GoString() string {
	return e.id
}

// ID returns the id of the seller
func (e *Entity) ID() string {
	return e.id
}
