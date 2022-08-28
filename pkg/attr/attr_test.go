package attr

import "fmt"

func ExampleAttr() {
	el := New("id", "abc")
	fmt.Println(el.Name())
	fmt.Println(el.Value())
	fmt.Println(el)

	// Output:
	// id
	// abc
	// {id abc}
}
