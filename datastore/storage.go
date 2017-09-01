package datastore

import "errors"

var errSetupFailed = errors.New("Failed to setup storage")

// SetupStorage setups up the specified storage mechanism
func SetupStorage(storageType string) error {

	return errSetupFailed
}
