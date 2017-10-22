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

// ErrInfrastructureUndefined to use when infrastructure is not defined
var ErrInfrastructureUndefined = errors.New("infrastructure is not defined")

// ErrInfrastructure for errors returned by infrastructure
type ErrInfrastructure struct {
	Msg string
}

func (e *ErrInfrastructure) Error() string {
	return fmt.Sprintf("infrastructure error: %s", e.Msg)
}