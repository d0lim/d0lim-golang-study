package stackeval_test

import (
	"fmt"

	"github.com/imdigo/stack-evaluator-go"
)

func Example_eval() {
	fmt.Println(stackeval.Eval("5"))
	fmt.Println(stackeval.Eval("1 + 2"))
	fmt.Println(stackeval.Eval("1 - 2 + 3"))
	fmt.Println(stackeval.Eval("3 * ( 3 + 1 * 3 ) / 2"))
	fmt.Println(stackeval.Eval("3 * ( ( 3 + 1 ) * 3 ) / 2"))
	// Output:
	// 5
	// 3
	// 2
	// 9
	// 18
}

