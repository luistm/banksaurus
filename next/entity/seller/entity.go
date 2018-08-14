package seller

import "fmt"

// NewFromID creates a new seller instance given it's ID
func NewFromID(id string) (*Entity, error) {
	return &Entity{id: id}, nil
}

// New creates a new instance of seller
func New(id string, name string) (*Entity, error) {
	return &Entity{id, name}, nil
}

// Entity for seller
type Entity struct {
	id   string
	name string
}

// GoString to satisfy the fmt.GoStringer interface
func (e *Entity) GoString() string {
	return fmt.Sprintf(">%s, %s", e.id, e.name)
}

// ID returns the id of the seller
func (e *Entity) ID() string {
	return e.id
}

// Name returns the name of the seller
func (e *Entity) Name() string {
	return e.name
}

// HasName ...
func (e *Entity) HasName() bool {
	if e.name == "" {
		return false
	}

	return true
}

// String to satisfy the string interface
func (e *Entity) String() string {
	if e.name != "" {
		return e.Name()
	}

	return e.ID()

}
