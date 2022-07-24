package local

import (
	"github.com/larry618/eval_lab/benchmark"
	"testing"

	evalmain "github.com/larry618/eval"
	evalloc "local/eval"
)

func BenchmarkEvalLocal(b *testing.B) {
	params := benchmark.CreateParams()

	cc := evalloc.NewCompileConfig()

	ctx := evalloc.NewCtxWithMap(cc, evalloc.ToValueMap(params))

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

func BenchmarkEvalLocal1(b *testing.B) {
	params := benchmark.CreateParams()

	cc := evalloc.NewCompileConfig()

	ctx := evalloc.NewCtxWithMap(cc, evalloc.ToValueMap(params))

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
