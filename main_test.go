package main

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	if testing.Short() {
		// setupDatabase()
	}
	result := m.Run()
	if !testing.Short() {
		// 	teardownDatabase()
	}
	os.Exit(result)
}
