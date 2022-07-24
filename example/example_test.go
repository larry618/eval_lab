package example

import (
	"fmt"
	evalmain "github.com/larry618/eval"
)

func ExampleEval() {
	output, err := evalmain.Eval(`(+ 1 v1)`, map[string]interface{}{
		"v1": 1,
	})
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	fmt.Printf("%v", output)

	// Output: 2
}
