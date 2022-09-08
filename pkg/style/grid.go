package style

const (
	PROP_GRID   = "grid"
	TAG_SUMMARY = "summary"
)

// Grid is a type that represents a CSS grid.
//
// See: https://developer.mozilla.org/en-US/docs/Web/CSS/grid
func GridStyle(value string) Style {
	// Check that grid is a valid grid property value.
	return New(PROP_GRID, value)
}

// TODO: WithGridStyle
