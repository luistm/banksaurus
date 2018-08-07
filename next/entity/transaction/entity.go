package transaction

import "time"

// New creates a new transaction
func New(date time.Time, sellerID string, value int64) (*Entity, error) {
	return &Entity{date: date, sellerID: sellerID, value: value}, nil
}

// Entity represents a transaction
type Entity struct {
	id       uint64
	date     time.Time
	sellerID string
	value    int64
}

// ID returns the identification of the transaction
func (t *Entity) ID() uint64 {
	return t.id
}
