package datastore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupStorage(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	name := "Fails to created storage if type is not defined"
	err := SetupStorage("")
	assert.Equal(t, err, errSetupFailed, name)
}
