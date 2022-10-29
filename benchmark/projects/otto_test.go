package projects

import (
	"github.com/onheap/eval_lab/benchmark"
	"testing"

	"github.com/robertkrimen/otto"
)

func Benchmark_otto(b *testing.B) {
	params := benchmark.CreateParams()

	vm := otto.New()

	script, err := vm.Compile("", benchmark.Example)
	if err != nil {
		b.Fatal(err)
	}

	_ = vm.Set("Origin", params["Origin"])
	_ = vm.Set("Country", params["Country"])
	_ = vm.Set("Adults", params["Adults"])
	_ = vm.Set("Value", params["Value"])

	var out otto.Value

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(script)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	ok, err := out.ToBoolean()
	if err != nil {
		b.Fatal(err)
	}
	if !ok {
		b.Fail()
	}
}
