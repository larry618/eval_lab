package projects

import (
	"testing"

	"github.com/antonmedv/expr"
	"github.com/onheap/eval_lab/benchmark"
)

func Benchmark_expr(b *testing.B) {
	params := benchmark.CreateParams()

	program, err := expr.Compile(benchmark.Example, expr.Env(params))
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = expr.Run(program, params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
