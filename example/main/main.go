package main

import (
	"fmt"
	"github.com/onheap/eval"
)

func main() {
	expr := `(and (>= age 30) (= gender "Male"))`

	vars := map[string]interface{}{
		"age":    30,
		"gender": "Male",
	}

	// new config and register variables
	config := eval.NewConfig(eval.RegVarAndOp(vars))

	// compile string expression to program
	program, err := eval.Compile(config, expr)

	// evaluation expression with variables
	res, err := program.Eval(eval.NewCtxFromVars(config, vars))

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", res)
}

func main1() {
	expr := `(and (>= age 30) (= gender "Male"))`

	vars := map[string]interface{}{
		"age":    30,
		"gender": "Male",
	}

	res, err := eval.Eval(expr, vars)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", res)
}
