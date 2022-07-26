package example

import (
	"fmt"
	"github.com/onheap/eval"
)

func ExampleEval() {
	output, err := eval.Eval(`(+ 1 v1)`, map[string]interface{}{
		"v1": 1,
	})
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	fmt.Printf("%v", output)

	// Output: 2
}
