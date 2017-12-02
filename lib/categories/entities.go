package categories

// New creates a new category instance but does not persist it
func New(name string) *Category {
	return &Category{name: name}
}

// Category ...
type Category struct {
	name string
}

// ID returns the ID of the category
func (c *Category) ID() string {
	return c.name
}

// String returns a string representing a Category instance
func (c *Category) String() string {
	return c.name
}
