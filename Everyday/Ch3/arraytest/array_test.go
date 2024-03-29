package arraytest

import "fmt"

func Example_array() {
	fruits := [3]string{"사과", "바나나", "토마토"}
	// You can use like
	// fruits := [...]string{"사과", "바나나", "토마토"}
	// to get Compiler know the length of array
	for _, fruit := range fruits {
		fmt.Printf("%s는 맛있다.\n", fruit)
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
}