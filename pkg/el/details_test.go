package el

import "fmt"

func ExampleDetails() {
	e := Details()
	fmt.Println(e)

	// Output: <details></details>
}

func ExampleWithSummary() {
	el := Details(WithSummary("abc"))
	fmt.Println(el)

	// Output: <details><summary>abc</summary></details>
}

func ExampleWithDetails() {
	parent := New("main", WithDetails("abc", "xyz"))
	fmt.Println(parent)

	// Output: <main><details><summary>abc</summary>xyz</details></main>
}
