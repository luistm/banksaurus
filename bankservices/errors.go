package bankservices

import "fmt"

// ErrInteractor for errors returned by Interactors
type ErrInteractor struct {
	Msg string
}

// Error ...
func (e *ErrInteractor) Error() string {
	return fmt.Sprintf("Interactor error, %s", e.Msg)
}
