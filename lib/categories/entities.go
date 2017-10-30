package categories

// Category ...
type Category struct {
	name string
}

func (c *Category) ID() string {
	return c.name
}

func (c *Category) String() string {
	return c.name
}
