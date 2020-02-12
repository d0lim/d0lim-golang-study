package main

import "fmt"

func fib(n int) int {
	p, q := 0, 1
	for i := 1; i <= n; i++ {
		p, q = q, p + q
	}
	return q
}

func main() {
	fmt.Println(fib(5))
}