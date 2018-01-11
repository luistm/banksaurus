package lib

// Presenter is use
type Presenter interface {
	Present(...Entity) error
}
