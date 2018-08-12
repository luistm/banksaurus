package transaction

// Command handles transaction command
type Command struct{}

// Executes the command instance
func (tc *Command) Execute(arguments map[string]interface{}) error {
	panic("transaction execute is not implemented")
}
