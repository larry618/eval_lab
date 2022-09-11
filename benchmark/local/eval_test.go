package local

import (
	"strconv"
	"testing"

	evalmain "github.com/onheap/eval"
	"github.com/onheap/eval_lab/benchmark"
	evalloc "local/eval"
)

func BenchmarkEvalLocal(b *testing.B) {
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

	cc := evalmain.NewCompileConfig(evalmain.RegisterVals(params))

	ctx := evalmain.NewCtxWithMap(cc, params)

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

func BenchmarkReference(b *testing.B) {
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

func BenchmarkEvalLocal1(b *testing.B) {
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

	cc := evalmain.NewCompileConfig(evalmain.RegisterVals(params))

	ctx := evalmain.NewCtxWithMap(cc, params)

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

func BenchmarkReference1(b *testing.B) {
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
