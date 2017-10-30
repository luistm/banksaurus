package descriptions

// New creates a new description instance but does not persist it
func New(slug string, friendlyName string) *Description {
	return &Description{
		slug:         slug,
		friendlyName: friendlyName,
	}
}

// Description describes a single entity of many, which an account interacts with
type Description struct {
	slug         string // TODO: Rename this to slug
	friendlyName string
}

// ID returns the ID of the description
func (d *Description) ID() string {
	return d.slug
}

func (d *Description) String() string {
	return d.friendlyName
}
