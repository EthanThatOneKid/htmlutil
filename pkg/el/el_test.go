package el

import (
	"fmt"

	"github.com/ethanthatonekid/htmlutil/pkg/attr"
	"github.com/ethanthatonekid/htmlutil/pkg/style"
)

func ExampleNew() {
	el := New("div")
	fmt.Println(el)

	// Output: <div></div>
}

func ExampleText() {
	text := Text("abc")
	fmt.Println(text)

	// Output: abc
}

func ExampleWithAttr() {
	a1 := attr.New("id", "my-id")
	a2 := attr.New("href", "https://example.com/")
	a3 := attr.New("abc", "xyz")

	e := New("a", WithAttr(a1, a2), WithAttr(a3))

	for _, a := range e.Attrs() {
		fmt.Println(a)
	}

	// Unordered output:
	// {id my-id}
	// {href https://example.com/}
	// {abc xyz}
}

func ExampleWithStyle() {
	s1 := style.New("color", "red")
	s2 := style.New("border-radius", "50%")
	s3 := style.New("border", "1px solid black")

	e := New("div", WithStyle(s1, s2), WithStyle(s3))

	for _, s := range e.Styles() {
		fmt.Println(s)
	}

	// Unordered output:
	// color:red;
	// border-radius:50%;
	// border:1px solid black;
}

func ExampleWithClass() {
	e := New("div", WithClass("a", "b"), WithClass("c"))

	for _, className := range e.Classes() {
		fmt.Println(className)
	}

	// Unordered output:
	// a
	// b
	// c
}

func ExampleWithChild() {
	e := New("div", WithChild(New("span")))
	fmt.Println(e)

	// Output: <div><span></span></div>
}

func ExampleWithText() {
	parent := New("p", WithText("abc"))
	fmt.Println(parent)

	// Output: <p>abc</p>
}
