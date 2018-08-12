package changesellername

import "errors"

var ErrInvalidSellerID = errors.New("invalid request seller id data")
var ErrInvalidSellerName = errors.New("invalid request seller name data")
var ErrDataWasNotCleaned = errors.New("data was not cleaned")

func NewRequest(sellerID string, sellerName string) (*Request, error) {
	if sellerID == "" {
		return &Request{}, ErrInvalidSellerID
	}

	if sellerName == "" {
		return &Request{}, ErrInvalidSellerName
	}

	return &Request{sellerID, sellerName}, nil
}

type Request struct {
	sellerID   string
	sellerName string
}

func (r *Request) SellerID() (string, error) {
	if r.sellerID == "" {
		return r.sellerID, ErrDataWasNotCleaned
	}

	return r.sellerID, nil
}

func (r *Request) SellerName() (string, error) {
	if r.sellerName == "" {
		return "", ErrDataWasNotCleaned
	}

	return r.sellerName, nil
}
