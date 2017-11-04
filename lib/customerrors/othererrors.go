package customerrors

import "errors"

// ErrBadInput for inputs which does not meet the requirements
var ErrBadInput = errors.New("bad input")
