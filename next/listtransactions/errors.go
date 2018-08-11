package listtransactions

import "fmt"

type customError struct {
	msg string
	err error
}

// AppendError allows to append an error to the current one
func (e *customError) AppendError(err error) error {
	e.err = err
	return e
}

func (e *customError) Error() string {
	return fmt.Errorf("%s: %s", e.msg, e.err.Error()).Error()
}
