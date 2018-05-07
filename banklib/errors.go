package banklib

import "errors"
import "fmt"

// ErrInfrastructureUndefined to use when infrastructure is not defined
var ErrInfrastructureUndefined = errors.New("infrastructure is not defined")

// ErrInfrastructure for errors returned by infrastructure
type ErrInfrastructure struct {
	Msg string
}

// Error ...
func (e *ErrInfrastructure) Error() string {
	return fmt.Sprintf("infrastructure error, %s", e.Msg)
}

// ErrBadInput for inputs which does not meet the requirements
var ErrBadInput = errors.New("bad input")

// ErrInteractorUndefined to use when the Interactor is not defined
var ErrInteractorUndefined = errors.New("interactor is not defined")

// ErrPresenterUndefined ...
var ErrPresenterUndefined = errors.New("presenter is not defined")

// ErrPresenter for errors returned by Interactors
type ErrPresenter struct {
	Msg string
}

// Error ...
func (e *ErrPresenter) Error() string {
	return fmt.Sprintf("presenter error, %s", e.Msg)
}

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
