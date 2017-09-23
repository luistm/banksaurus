package infrastructure

import "errors"

var errSetupFailed = errors.New("Failed to setup storage")

// SetupStorage configures the storage to persist data
func SetupStorage() error {

	return errSetupFailed
}
