package lib

import (
	"fmt"
)

// Entity is the interface each an object must implement in
// order to be identified
type Entity interface {
	fmt.Stringer
	ID() string
}
