package lib

// Creator interface to create entities
type Creator interface {
	Create(string) (Entity, error)
}
