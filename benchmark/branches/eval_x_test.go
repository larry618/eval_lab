package branches

import (
	evaldev "dev/eval"
	evalmain "github.com/onheap/eval"
	"github.com/onheap/eval_lab/benchmark"
	"testing"
)

func BenchmarkEvalDev2(b *testing.B) {
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

func BenchmarkEvalMain2(b *testing.B) {
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

func BenchmarkEvalDev3(b *testing.B) {
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

func BenchmarkEvalMain3(b *testing.B) {
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

func BenchmarkEvalDev4(b *testing.B) {
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

func BenchmarkEvalMain4(b *testing.B) {
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

func BenchmarkEvalDev5(b *testing.B) {
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

func BenchmarkEvalMain5(b *testing.B) {
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

func BenchmarkEvalDev6(b *testing.B) {
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

func BenchmarkEvalMain6(b *testing.B) {
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

func BenchmarkEvalDev7(b *testing.B) {
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

func BenchmarkEvalMain7(b *testing.B) {
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

func BenchmarkEvalDev8(b *testing.B) {
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

func BenchmarkEvalMain8(b *testing.B) {
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

func BenchmarkEvalDev9(b *testing.B) {
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

func BenchmarkEvalMain9(b *testing.B) {
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
