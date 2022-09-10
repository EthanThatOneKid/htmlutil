package data

// See: https://github.com/a-h/generate#readme

//go:generate schema-generate -o properties.gen.go -p data properties.schema.json
func main() {
	cssProperty := Property{}

	cssProperty
}
