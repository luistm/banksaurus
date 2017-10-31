package customerrors

import (
	"errors"
	"fmt"
)

// ErrRepositoryUndefined to use when repository is not defined
var ErrRepositoryUndefined = errors.New("repository is not defined")

// ErrRepository for errors returned by repositories
type ErrRepository struct {
	Msg string
}

// Error to return the error message
func (e *ErrRepository) Error() string {
	return fmt.Sprintf("repository error, %s", e.Msg)
}
