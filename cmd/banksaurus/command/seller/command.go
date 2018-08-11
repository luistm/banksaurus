package seller

// Command seller
type Command struct{}

// Execute the seller command with arguments
func (s *Command) Execute(arguments map[string]interface{}) error {

	if arguments["seller"].(bool) && arguments["new"].(bool) {
		panic("seller new not implemented")
	}

	if arguments["seller"].(bool) && arguments["show"].(bool) {
		panic("seller show not implemented")
	}

	if arguments["seller"].(bool) && arguments["change"].(bool) {
		panic("seller change not implemented")
	}

	return nil
}
