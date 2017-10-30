package lib

// Creator interface to create entities
type Creator interface {
	Create() (Entity, error)
}
