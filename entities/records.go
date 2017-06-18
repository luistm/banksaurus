package entities

// Record represents a single line the the file
type Record struct {
	Record []string
}

// Valid returns a boolean regarding if the record is valid or not
func (r *Record) Valid() bool {
	if len(r.Record) == 8 {
		return true
	}

	return false
}
