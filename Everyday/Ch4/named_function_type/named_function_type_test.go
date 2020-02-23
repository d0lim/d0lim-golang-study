package named_function_type

import "fmt"

type BinOp func(int, int) int
type BinSub func(int, int) int

func BinOpToBinSub(f BinOp) BinSub {
	var count int
	return func(a, b int) int {
		fmt.Println(f(a, b))
		count++
		return count
	}
}


func ExampleBinOpToBinSub() {
	sub := BinOpToBinSub(func(a, b int) int {
		return a + b
	})
	sub(5, 7)
	sub(5, 7)
	count := sub(5, 7)
	fmt.Println("count:", count)
	// Output:
	// 12
	// 12
	// 12
	// count: 3
}
