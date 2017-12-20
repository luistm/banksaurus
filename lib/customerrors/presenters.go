package customerrors

import (
	"errors"
	"fmt"
)

// ErrPresenterUndefined ...
var ErrPresenterUndefined = errors.New("presenter is not defined")

// ErrPresenter for errors returned by Interactors
type ErrPresenter struct {
	Msg string
}

func (e *ErrPresenter) Error() string {
	return fmt.Sprintf("Interactor error, %s", e.Msg)
}
