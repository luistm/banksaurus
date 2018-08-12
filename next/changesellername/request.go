package changesellername

import "errors"

var ErrInvalidInput = errors.New("invalid request sellerID data")

func NewRequest(sellerID string) (*Request, error){
	if sellerID == ""{
		return &Request{}, ErrInvalidInput
	}
	return &Request{sellerID}, nil
}

type Request struct{
	sellerID string
}

func (r *Request) SellerID() string{
	// TODO: Return error if sellerID is empty
	return r.sellerID
}