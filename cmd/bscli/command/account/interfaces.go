package account

// Commander is the interface commands should implement
type Commander interface {
	Execute(map[string]interface{}) error
}
