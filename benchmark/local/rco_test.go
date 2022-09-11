package local

import (
	"github.com/onheap/eval_lab/benchmark"
	evalloc "local/eval"
	"testing"
)

func BenchmarkEvalLocalRCO(b *testing.B) {
	params := benchmark.CreateParams()

	cc := evalloc.NewCompileConfig(evalloc.RegisterVals(params))

	ctx := evalloc.NewCtxWithMap(cc, params)

	s := `
(and
 (or
   (= Origin "MOW")
   (= Country "RU"))
 (or
   (>= Value 100)
   (= Adults 1)))
`

	program, err := evalloc.Compile(cc, s)

	var out evalloc.Value

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = program.TryEval(ctx)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
