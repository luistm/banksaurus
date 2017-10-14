package descriptions

// IRepository ...
type IRepository interface{}

// Description describes a single entity of many, which an account interacts with
type Description struct {
	rawName      string
	friendlyName string
}
