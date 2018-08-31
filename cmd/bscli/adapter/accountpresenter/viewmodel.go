package accountpresenter

func NewViewModel(value string) (*ViewModel, error) {
	return &ViewModel{value}, nil
}

type ViewModel struct {
	Value string
}
