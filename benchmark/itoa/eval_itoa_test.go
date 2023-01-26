package itoa

import (
	"strconv"
	"testing"

	"github.com/onheap/eval"
	"github.com/onheap/eval_lab/benchmark"
)

func Benchmark_eval(b *testing.B) {
	params := benchmark.CreateParams()

	cc := eval.NewConfig(eval.RegVarAndOp(params))

	ctx := eval.NewCtxFromVars(cc, params)

	s := `
(and
  (or
    (= Origin "MOW")
    (= Country "RU"))
  (or
    (>= Value 100)
    (= Adults 1)))
`

	program, err := eval.Compile(cc, s)

	var out eval.Value

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

func _Benchmark_tryEval(b *testing.B) {
	params := benchmark.CreateParams()

	cc := eval.NewConfig(eval.RegVarAndOp(params))

	ctx := eval.NewCtxFromVars(cc, params)

	s := `
(and
  (or
    (= Origin "MOW")
    (= Country "RU"))
  (or
    (>= Value 100)
    (= Adults 1)))
`

	program, err := eval.Compile(cc, s)

	var out eval.Value

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

func Benchmark_itoa(b *testing.B) {
	var out interface{}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		out = strconv.Itoa(12345678)
	}
	b.StopTimer()

	if out == nil {
		b.Fatal(out)
	}
}
