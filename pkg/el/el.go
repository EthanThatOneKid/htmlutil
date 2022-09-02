package el

import (
	"strings"

	"github.com/ethanthatonekid/htmlutil/pkg/attr"
	"github.com/ethanthatonekid/htmlutil/pkg/style"
)

const (
	TAG_TEXT_CONTENT = "__TEXT_CONTENT"
)

// El is the base of all HTML elements.
type El struct {
	tag     string
	attrs   map[string]attr.Attr
	styles  map[string]style.Style
	classes map[string]struct{}

	// An El may either text content or child elements.
	TextContent string
	Children    []El
}

// Element transformer function type.
type Tf func(El) El

// Getter for readonly tag of the element.
func (e El) Tag() string {
	return e.tag
}

// Getter for attributes of the element.
func (e El) Attrs() (attrs []attr.Attr) {
	for _, value := range e.attrs {
		attrs = append(attrs, value)
	}
	return
}

// Getter for styles of the element.
func (e El) Styles() (styles []style.Style) {
	for _, value := range e.styles {
		styles = append(styles, value)
	}
	return
}

// Getter for readonly classes of the element.
func (e El) Classes() (classes []string) {
	for className := range e.classes {
		classes = append(classes, className)
	}
	return
}

// Apply element transformer(s) to the element.
// TODO: Update e.Apply(el.WithX()) pattern to e.With(el.X())
func (e *El) Apply(tfs ...Tf) El {
	for _, tf := range tfs {
		*e = tf(*e)
	}
	return *e
}

// Instantiates new element with tag name and empty attribute/style/class maps.
func New(tag string, tfs ...Tf) El {
	emptyAttrs := make(map[string]attr.Attr)
	emptyStyles := make(map[string]style.Style)
	emptyClasses := make(map[string]struct{})

	e := El{
		tag:     tag,
		attrs:   emptyAttrs,
		styles:  emptyStyles,
		classes: emptyClasses,
	}

	return e.Apply(tfs...)
}

// Instantiates new text node from text content.
func Text(content string) El {
	return New(TAG_TEXT_CONTENT, func(e El) El {
		e.TextContent = content
		return e
	})
}

// Returns an element transformer that adds a variadic number of attributes to the element.
func WithAttr(attrs ...attr.Attr) Tf {
	return func(e El) El {
		for _, a := range attrs {
			e.attrs[a.Name()] = a
		}
		return e
	}
}

// Returns an element transformer that adds a variadic number of styles to the element.
func WithStyle(styles ...style.Style) Tf {
	return func(e El) El {
		for _, s := range styles {
			e.styles[s.Prop()] = s
		}
		return e
	}
}

// Returns an element transformer that adds a variadic number of children to the element.
func WithChild(children ...El) Tf {
	return func(e El) El {
		e.Children = append(e.Children, children...)
		return e
	}
}

// Returns an element transformer that adds a variadic number of classes to the element.
func WithClass(classes ...string) Tf {
	return func(e El) El {
		for _, c := range classes {
			e.classes[c] = struct{}{}
		}
		return e
	}
}

// Appends a child text node.
func WithText(content string) Tf {
	return func(e El) El {
		return e.Apply(WithChild(Text(content)))
	}
}

// String returns the HTML representation of the element.
func (e El) String() string {
	if e.Tag() == TAG_TEXT_CONTENT {
		return e.TextContent
	}

	var b strings.Builder

	b.WriteString("<")
	b.WriteString(e.tag)

	for k, v := range e.attrs {
		b.WriteString(" ")
		b.WriteString(k)
		b.WriteString("=\"")
		b.WriteString(v.Value())
		b.WriteString("\"")
	}

	if len(e.classes) > 0 {
		b.WriteString(" class=\"")
		b.WriteString(strings.Join(e.Classes(), " "))
		b.WriteString("\"")
	}

	if len(e.styles) > 0 {
		b.WriteString(" style=\"")
		for _, style := range e.styles {
			b.WriteString(style.String())
		}
		b.WriteString("\"")
	}

	b.WriteString(">")

	for _, child := range e.Children {
		b.WriteString(child.String())
	}

	b.WriteString("</")
	b.WriteString(e.tag)
	b.WriteString(">")
	return b.String()
}
