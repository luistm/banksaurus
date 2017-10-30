package descriptions

// New creates a new description instance but does not persist it
func New(slug string, name string) *Description {
	return &Description{slug: slug, name: name}
}

// Description describes a single entity of many, which an account interacts with
type Description struct {
	slug string
	name string
}

// ID returns the ID of the description
func (d *Description) ID() string {
	return d.slug
}

func (d *Description) String() string {
	return d.name
}
