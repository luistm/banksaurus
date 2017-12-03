package lib

import (
	"fmt"
)

// Identifier is the interface each an object must implement in
// order to be identified
type Identifier interface {
	fmt.Stringer
	ID() string
}
