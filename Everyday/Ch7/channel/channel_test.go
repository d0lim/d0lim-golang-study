package channel_test

import "fmt"

func Example_simpleChannel() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// Output:
	// 1
	// 2
	// 3
}

func Example_simpleChannel2() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for num := range c {
		fmt.Println(num)
	}
	// Output:
	// 1
	// 2
	// 3
}

func Example_simpleChannel3() {
	c := func() <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			c <- 1
			c <- 2
			c <- 3
		}()
		return c
	}()
	for num := range c {
		fmt.Println(num)
	}
	// Output:
	// 1
	// 2
	// 3
}

func Example_closedChannel() {
	c := make(chan int)
	close(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// Output:
	// 0
	// 0
	// 0
}