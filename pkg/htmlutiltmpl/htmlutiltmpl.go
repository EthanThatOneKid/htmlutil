package htmlutiltmpl

import (
	"context"
	"html/template"
	"os"
)

// Config represents the inputs needed to populate the template.
type Config struct {
	// TemplateFilename is the path to the file where the template is stored.
	TemplateFilename string
	// GeneratedFilename is the path where the generated file should be stored
	// after populating the template.
	GeneratedFilename string
	// Version is the version being released.
	Version string
}

// Run reads a template file, populates it with the configured version and
// writes out the populated template to the generated filename.
func Run(_ context.Context, c Config) error {
	t, err := template.ParseFiles(c.TemplateFilename)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(c.GeneratedFilename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, c)
}
