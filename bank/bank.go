package bank

// Interactor is the interface each use case must implement
type Interactor interface {
	Execute() error
}
