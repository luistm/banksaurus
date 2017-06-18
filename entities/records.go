package entities

import "strings"

// Record represents a single line of an importend transaction.
// Data will be imported most likely from a text base format.
// Therefore a record will be the representation of each of of those lines.
// Forthe time being, nothis will be closely coupled to the CSV format.
type Record struct {
	Record []string
}

// Valid returns a boolean regarding if the record is valid or not
func (r *Record) Valid() bool {
	// TODO: Allow the extention with  multiple types of validators
	// TODO: Use regular expressions where appropriated
	if r.Record[0] == " " {
		return false
	}

	if len(r.Record) == 8 && !strings.HasPrefix(r.Record[0], "Data mov") {
		return true
	}

	return false
}
