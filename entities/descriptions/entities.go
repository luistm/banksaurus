package descriptions

// Description describes a single entity of many, which an account interacts with
type Description struct {
	rawName      string
	friendlyName string
}

func (d *Description) ID() string {
	return d.rawName
}

func (d *Description) String() string {
	return d.friendlyName
}
