package showaccount

import "errors"

var ErrInvalidAccountID = errors.New("invalid account id")

func NewRequest(accountID string) (*Request, error) {
	if accountID == "" {
		return &Request{}, ErrInvalidAccountID
	}
	return &Request{accountID}, nil
}

type Request struct {
	accountID string
}

func (r *Request) AccountID() (string, error) {
	return r.accountID, nil
}
