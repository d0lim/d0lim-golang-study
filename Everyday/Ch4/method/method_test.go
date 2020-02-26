package method_test

import "fmt"

type VertexID int

func (id VertexID) String() string {
	return fmt.Sprintf("VertexID(%d)", id)
}

func ExampleVertexID_print() {
	i := VertexID(100)
	fmt.Println(i)
	// Output:
	// 100
}

func ExampleVertexID_String() {
	i := VertexID(100)
	// This works because of Interface of Go.
	fmt.Println(i)
	// Output:
	// VertexID(100)
}