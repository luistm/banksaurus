package lib

import (
	"fmt"
)

// Entity generic type for domain entities
type Entity interface {
	fmt.Stringer
	ID() string
}
