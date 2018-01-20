package testkit

import (
	"reflect"
	"testing"
)

// TODO: unit tests for this package

// AssertEqual asserts that two values are equal
func AssertEqual(t *testing.T, expected interface{}, got interface{}) {
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("\n\n---- Expected:\n%v\n++++ Got:\n%v\n", expected, got)
	}
}
