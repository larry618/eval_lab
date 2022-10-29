package projects

import (
	"testing"

	"github.com/Knetic/govaluate"
	"github.com/onheap/eval_lab/benchmark"
)

func Benchmark_govaluate(b *testing.B) {
	params := benchmark.CreateParams()

	expression, err := govaluate.NewEvaluableExpression(benchmark.Example)

	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = expression.Evaluate(params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
