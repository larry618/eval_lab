package projects

import (
	"github.com/onheap/eval_lab/benchmark"
	"testing"
)

func Benchmark_native(b *testing.B) {
	params := benchmark.CreateParams()

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out = (params["Origin"] == "MOW" || params["Country"] == "RU") &&
			(params["Value"].(int) >= 100 || params["Adults"].(int) == 1)
	}
	b.StopTimer()

	if !out.(bool) {
		b.Fail()
	}
}
