package customerrors

import (
	"errors"
	"fmt"
)

// ErrInteractor for errors returned by interactors
type ErrInteractor struct {
	Msg string
}

func (e *ErrInteractor) Error() string {
	return fmt.Sprintf("interactor error, %s", e.Msg)
}

// ErrInteractorUndefined to use when the interactor is not defined
var ErrInteractorUndefined = errors.New("interactor is not defined")
