package seller

// New creates a new seller instance
func New(slug string, name string) *Seller {
	return &Seller{slug: slug, name: name}
}

// Seller is the destiny of the money spent in a transaction
type Seller struct {
	slug string
	name string
}

// ID returns the ID of the seller
func (s *Seller) ID() string {
	return s.slug
}

// String returns a string representing a Seller
func (s *Seller) String() string {
	if s.name == "" {
		return s.slug
	}
	return s.name
}
