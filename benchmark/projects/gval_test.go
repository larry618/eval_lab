package projects

import (
	"context"
	"github.com/onheap/eval_lab/benchmark"
	"testing"

	"github.com/PaesslerAG/gval"
)

func Benchmark_gval(b *testing.B) {
	params := benchmark.CreateParams()
	ctx := context.Background()

	var out interface{}
	var err error

	eval, err := gval.Full().NewEvaluable(benchmark.Example)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = eval(ctx, params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
