package branches

import (
	"testing"

	evaldev "dev/eval"
	evalmain "github.com/onheap/eval"
	"github.com/onheap/eval_lab/benchmark"
)

func BenchmarkCompileDev(b *testing.B) {
	params := benchmark.CreateParams()

	cc := evaldev.NewCompileConfig(evaldev.RegisterVals(params))

	s := `
(and
 (or
   (= Origin "MOW")
   (= Country "RU"))
 (or
   (>= Value 100)
   (= Adults 1)))
`

	var program *evaldev.Expr
	var err error
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		program, err = evaldev.Compile(cc, s)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	ctx := evaldev.NewCtxWithMap(cc, params)
	out, err := program.Eval(ctx)

	if err != nil {
		b.Fatal(err)
	}

	if !out.(bool) {
		b.Fail()
	}
}

func BenchmarkCompileMain(b *testing.B) {
	params := benchmark.CreateParams()

	cc := evalmain.NewCompileConfig(evalmain.RegisterVals(params))

	s := `
(and
 (or
   (= Origin "MOW")
   (= Country "RU"))
 (or
   (>= Value 100)
   (= Adults 1)))
`

	var program *evalmain.Expr
	var err error
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		program, err = evalmain.Compile(cc, s)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	ctx := evalmain.NewCtxWithMap(cc, params)
	out, err := program.Eval(ctx)

	if err != nil {
		b.Fatal(err)
	}

	if !out.(bool) {
		b.Fail()
	}
}

func BenchmarkCompileDev1(b *testing.B) {
	params := benchmark.CreateParams()

	cc := evaldev.NewCompileConfig(evaldev.RegisterVals(params))

	s := `
(and
 (or
   (= Origin "MOW")
   (= Country "RU"))
 (or
   (>= Value 100)
   (= Adults 1)))
`

	var program *evaldev.Expr
	var err error
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		program, err = evaldev.Compile(cc, s)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	ctx := evaldev.NewCtxWithMap(cc, params)
	out, err := program.Eval(ctx)

	if err != nil {
		b.Fatal(err)
	}

	if !out.(bool) {
		b.Fail()
	}
}

func BenchmarkCompileMain1(b *testing.B) {
	params := benchmark.CreateParams()

	cc := evalmain.NewCompileConfig(evalmain.RegisterVals(params))

	s := `
(and
 (or
   (= Origin "MOW")
   (= Country "RU"))
 (or
   (>= Value 100)
   (= Adults 1)))
`

	var program *evalmain.Expr
	var err error
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		program, err = evalmain.Compile(cc, s)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	ctx := evalmain.NewCtxWithMap(cc, params)
	out, err := program.Eval(ctx)

	if err != nil {
		b.Fatal(err)
	}

	if !out.(bool) {
		b.Fail()
	}
}
