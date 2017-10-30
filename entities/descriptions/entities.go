package descriptions

// New creates a new description instance but does not persist it
func New(rawName string, friendlyName string) *Description {
	return &Description{
		rawName:      rawName,
		friendlyName: friendlyName,
	}
}

// Description describes a single entity of many, which an account interacts with
type Description struct {
	rawName      string // TODO: Rename this to slug
	friendlyName string
}

// ID returns the ID of the description
func (d *Description) ID() string {
	return d.rawName
}

func (d *Description) String() string {
	return d.friendlyName
}
