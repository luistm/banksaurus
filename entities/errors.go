package entities

import "errors"
import "fmt"

// ErrBadInput for inputs wich does not meet the requirements
var ErrBadInput = errors.New("bad input")

// ErrRepositoryIsNil to use when repository is not defined
var ErrRepositoryIsNil = errors.New("repository is undefined")

// ErrRepository for errors returned by repositories
type ErrRepository struct {
	Msg string
}

// Error to return the error message
func (e *ErrRepository) Error() string {
	return fmt.Sprintf("repository error: %s", e.Msg)
}
