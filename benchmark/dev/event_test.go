package dev

import (
	evaldev "dev/eval"
	evalmain "github.com/onheap/eval"
	"github.com/onheap/eval_lab/benchmark"
	"testing"
)

func BenchmarkEventEvalDev(b *testing.B) {
	eventChan := make(chan evaldev.Event, 1024)
	go func() {
		for range eventChan {

		}
	}()
	eventChan <- evaldev.Event{} // wait channel ready

	params := benchmark.CreateParams()

	cc := evaldev.NewCompileConfig(evaldev.RegisterVals(params), evaldev.EnableReportEvent)

	ctx := evaldev.NewCtxWithMap(cc, params)

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
	program.EventChan = eventChan

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

func BenchmarkEventEvalMain(b *testing.B) {

	eventChan := make(chan evalmain.Event, 1024)
	go func() {
		for range eventChan {
		}
	}()
	eventChan <- evalmain.Event{} // wait channel ready

	params := benchmark.CreateParams()

	cc := evalmain.NewCompileConfig(evalmain.RegisterVals(params), evalmain.EnableReportEvent)

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
	program.EventChan = eventChan

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
