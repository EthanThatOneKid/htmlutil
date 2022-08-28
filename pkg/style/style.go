package style

// Style is an HTML element style.
type Style struct {
	prop  string
	value string
}

// Getter for readonly property of the style.
func (s Style) Prop() string {
	return s.prop
}

// Getter for readonly value of the style.
func (s Style) Value() string {
	return s.value
}

// String returns human-readable representation of the style.
func (s Style) String() string {
	return s.prop + ":" + s.value + ";"
}

// Creates a new style.
func New(prop, value string) Style {
	return Style{prop, value}
}
