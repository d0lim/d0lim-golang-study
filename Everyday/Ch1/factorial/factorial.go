package main

import "fmt"

// fac returns n!.
func fac(n int) int {
	if n <= 0 {
		return 1
	}
	return n * fac(n - 1)
}

func main() {
	fmt.Println(fac(5))
}