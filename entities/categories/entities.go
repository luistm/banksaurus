package categories

// IRepository is the interface for repositories which handle categories
type IRepository interface {
	Save(*Category) error
	Get(string) (*Category, error)
	GetAll() ([]*Category, error)
}

// Category ...
type Category struct {
	name string
}

func (c *Category) String() string {
	return c.name
}
