package datastore

import (
	"log"
	"os"
)

// OpenFile opens and returns a file handler for a CSV file
func OpenFile(inputFilePath string) (*os.File, error) {
	_, err := os.Stat(inputFilePath)
	if err != nil {
		log.Fatal("Input error: ", err)
	}

	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}

	return file, err
}
