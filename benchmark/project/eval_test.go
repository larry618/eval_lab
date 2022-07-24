package project

import (
	evalmain "github.com/larry618/eval"
	"github.com/larry618/eval_lab/benchmark"
	"testing"
)

func Benchmark_eval(b *testing.B) {
	params := benchmark.CreateParams()

	cc := evalmain.NewCompileConfig()

	ctx := evalmain.NewCtxWithMap(cc, evalmain.ToValueMap(params))

	s := `
(and
  (or
    (= Origin "MOW")
    (= Country "RU"))
  (or
    (>= Value 100)
    (= Adults 1)))
`

	program, err := evalmain.Compile(cc, s)

	var out evalmain.Value

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = program.Eval(ctx)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
