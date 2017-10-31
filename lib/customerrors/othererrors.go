package customerrors

import "errors"

// ErrBadInput for inputs wich does not meet the requirements
var ErrBadInput = errors.New("bad input")
