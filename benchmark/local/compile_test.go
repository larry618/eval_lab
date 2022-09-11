package local

import (
	"testing"

	evalmain "github.com/onheap/eval"
	"github.com/onheap/eval_lab/benchmark"
	evalloc "local/eval"
)

func BenchmarkCompileLocal(b *testing.B) {
	params := benchmark.CreateParams()

	cc := evalloc.NewCompileConfig(evalloc.RegisterVals(params))

	s := `
(and
 (or
   (= Origin "MOW")
   (= Country "RU"))
 (or
   (>= Value 100)
   (= Adults 1)))
`

	var program *evalloc.Expr
	var err error
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		program, err = evalloc.Compile(cc, s)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	ctx := evalloc.NewCtxWithMap(cc, params)
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

func BenchmarkCompileLocal1(b *testing.B) {
	params := benchmark.CreateParams()

	cc := evalloc.NewCompileConfig(evalloc.RegisterVals(params))

	s := `
(and
 (or
   (= Origin "MOW")
   (= Country "RU"))
 (or
   (>= Value 100)
   (= Adults 1)))
`

	var program *evalloc.Expr
	var err error
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		program, err = evalloc.Compile(cc, s)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	ctx := evalloc.NewCtxWithMap(cc, params)
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
