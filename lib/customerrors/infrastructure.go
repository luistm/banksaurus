package customerrors

import "errors"
import "fmt"

// ErrInfrastructureUndefined to use when infrastructure is not defined
var ErrInfrastructureUndefined = errors.New("infrastructure is not defined")

// ErrInfrastructure for errors returned by infrastructure
type ErrInfrastructure struct {
	Msg string
}

func (e *ErrInfrastructure) Error() string {
	return fmt.Sprintf("infrastructure error, %s", e.Msg)
}
