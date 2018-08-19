package changesellername

import "errors"

var (
	// ErrInvalidSellerID ...
	ErrInvalidSellerID   = errors.New("invalid request seller id data")

	// ErrInvalidSellerName ...
	ErrInvalidSellerName = errors.New("invalid request seller name data")

	// ErrDataWasNotCleaned ...
	ErrDataWasNotCleaned = errors.New("data was not cleaned")
)

// NewRequest creates a new request to changesellername
func NewRequest(sellerID string, sellerName string) (*Request, error) {
	if sellerID == "" {
		return &Request{}, ErrInvalidSellerID
	}

	if sellerName == "" {
		return &Request{}, ErrInvalidSellerName
	}

	return &Request{sellerID, sellerName}, nil
}

// Request to changesellername
type Request struct {
	sellerID   string
	sellerName string
}

// SellerID returns the id of the seller
func (r *Request) SellerID() (string, error) {
	if r.sellerID == "" {
		return r.sellerID, ErrDataWasNotCleaned
	}

	return r.sellerID, nil
}

// SellerName returns the name of the seller
func (r *Request) SellerName() (string, error) {
	if r.sellerName == "" {
		return "", ErrDataWasNotCleaned
	}

	return r.sellerName, nil
}
