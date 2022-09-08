package attr

// Attr is an HTML element attribute.
//
// See: https://pkg.go.dev/golang.org/x/net/html#Attribute
type Attr struct {
	name  string
	value string
}

// Getter for readonly name of the attribute.
func (a Attr) Name() string {
	return a.name
}

// Getter for readonly value of the attribute.
func (a Attr) Value() string {
	return a.value
}

// Creates a new attribute.
func New(name, value string) Attr {
	return Attr{name, value}
}
