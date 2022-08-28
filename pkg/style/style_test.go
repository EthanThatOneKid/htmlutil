package style

import "fmt"

func ExampleStyle() {
	el := New("color", "red")
	fmt.Println(el.Prop())
	fmt.Println(el.Value())
	fmt.Println(el)

	// Output:
	// color
	// red
	// color:red;
}
