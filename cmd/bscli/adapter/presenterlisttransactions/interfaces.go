package presenterlisttransactions

import "io"

// Viewer interface
type Viewer interface {
	io.Writer
}
