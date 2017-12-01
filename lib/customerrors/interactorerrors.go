package customerrors

import (
	"errors"
	"fmt"
)

// ErrInteractor for errors returned by Interactors
type ErrInteractor struct {
	Msg string
}

func (e *ErrInteractor) Error() string {
	return fmt.Sprintf("Interactor error, %s", e.Msg)
}

// ErrInteractorUndefined to use when the Interactor is not defined
var ErrInteractorUndefined = errors.New("Interactor is not defined")
