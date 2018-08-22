package createaccount

// NewInteractor creates a new instance of the create account interactor
func NewInteractor() (*Interactor, error){
	return &Interactor{}, nil
}

// Interactor for creating an account
type Interactor struct{}