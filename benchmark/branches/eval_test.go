package branches

import (
	"strconv"
	"testing"

	evaldev "dev/eval"
	evalmain "github.com/onheap/eval"
	"github.com/onheap/eval_lab/benchmark"
)

func BenchmarkEvalDev(b *testing.B) {
	params := benchmark.CreateParams()

	cc := evaldev.NewConfig(evaldev.RegVarAndOp(params))

	ctx := evaldev.NewCtxFromVars(cc, params)

	s := `
(and
 (or
   (= Origin "MOW")
   (= Country "RU"))
 (or
   (>= Value 100)
   (= Adults 1)))
`

	program, err := evaldev.Compile(cc, s)

	var out evaldev.Value

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

func BenchmarkEvalMain(b *testing.B) {
	params := benchmark.CreateParams()

	cc := evalmain.NewConfig(evalmain.RegVarAndOp(params))

	ctx := evalmain.NewCtxFromVars(cc, params)

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

func BenchmarkItoa(b *testing.B) {
	var out interface{}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		out = strconv.Itoa(n)
	}
	b.StopTimer()

	if out == nil {
		b.Fatal(out)
	}
}

func BenchmarkEvalDev1(b *testing.B) {
	params := benchmark.CreateParams()

	cc := evaldev.NewConfig(evaldev.RegVarAndOp(params))

	ctx := evaldev.NewCtxFromVars(cc, params)

	s := `
(and
 (or
   (= Origin "MOW")
   (= Country "RU"))
 (or
   (>= Value 100)
   (= Adults 1)))
`

	program, err := evaldev.Compile(cc, s)

	var out evaldev.Value

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

func BenchmarkEvalMain1(b *testing.B) {
	params := benchmark.CreateParams()

	cc := evalmain.NewConfig(evalmain.RegVarAndOp(params))

	ctx := evalmain.NewCtxFromVars(cc, params)

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

func BenchmarkItoa1(b *testing.B) {
	var out interface{}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		out = strconv.Itoa(n)
	}
	b.StopTimer()

	if out == nil {
		b.Fatal(out)
	}
}
