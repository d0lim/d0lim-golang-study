package pipeline

import "fmt"

func ExamplePlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	// for num := range PlusOne(PlusOne(c)) {
	// 	fmt.Println(num)
	// }
	for num := range Chain(PlusOne, PlusOne)(c) {
		fmt.Println(num)
	}

	// Output:
	// 7 
	// 5
	// 10
}

