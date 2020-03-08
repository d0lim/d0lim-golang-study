package pipeline

// import (
// 	"context"
// 	"sync"
// )

// IntPipe is the type of function that compose int pipeline
type IntPipe func(<-chan int) <-chan int

// Chain get IntPipes and connect pipelines recursively using function closure and return IntPipe
func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}

// PlusOne returns a channel of num + 1 for nums received from in.
func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}
