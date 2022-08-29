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
	id := attr.New("id", "my-id")
	el := New("div", WithAttr(id))
	fmt.Println(el)

	// Output: <div id="my-id"></div>
}

func ExampleWithStyle() {
	style := style.New("color", "red")
	el := New("div", WithStyle(style))
	fmt.Println(el)

	// Output: <div style="color:red;"></div>
}

func ExampleWithClass() {
	el := New("div", WithClass("a", "b", "c"))
	fmt.Println(el)

	// Output: <div class="a b c"></div>
}

func ExampleWithChild() {
	el := New("div", WithChild(New("span")))
	fmt.Println(el)

	// Output: <div><span></span></div>
}

func ExampleWithText() {
	parent := New("p", WithText("abc"))
	fmt.Println(parent)

	// Output: <p>abc</p>
}
