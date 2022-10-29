package local

import (
	evalmain "github.com/onheap/eval"
	"github.com/onheap/eval_lab/benchmark"
	evalloc "local/eval"
	"testing"
)

func BenchmarkEventEvalLocal(b *testing.B) {
	eventChan := make(chan evalloc.Event, 1024)
	go func() {
		for range eventChan {

		}
	}()
	eventChan <- evalloc.Event{} // wait channel ready

	params := benchmark.CreateParams()

	cc := evalloc.NewCompileConfig(evalloc.RegisterVals(params), evalloc.EnableReportEvent)

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
	program.EventChan = eventChan

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
