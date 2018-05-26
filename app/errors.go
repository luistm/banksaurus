package app

import "errors"

var (
	// ErrPathIsFile ...
	ErrPathIsFile = errors.New("the path is a file")

	// ErrMalformedPath ...
	ErrMalformedPath = errors.New("path is malformed")

	// ErrPathDoesNotExist ...
	ErrPathDoesNotExist = errors.New("path does not exist")

	// ErrInvalidAppStructure ...
	ErrInvalidAppStructure = errors.New("invalid application structure")
)
