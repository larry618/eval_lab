package main

import (
	"github.com/onheap/eval_lab/tui"
	"local/eval"
)

func main() {

	s := `
(and
 (or
   (= Origin "MOW1") (= Country "RU"))  
 (or (>= Value (+ 33 67)) (= Adults 1)))
`

	//vals := benchmark.CreateParams()
	//cc := eval.NewCompileConfig(eval.RegisterVals(vals), eval.EnableReportEvent, eval.Optimizations(false))

	vals := map[string]interface{}{
		"Origin":  "MOW",
		"Country": "RU",
		//"Value":   100,
		"Adults": 1,
		//"Value":   -100,
		//"Adults":  -1,
	}

	cc := &eval.CompileConfig{
		SelectorMap: map[string]eval.SelectorKey{
			"Origin":  eval.SelectorKey(0),
			"Country": eval.SelectorKey(1),
			"Value":   eval.SelectorKey(2),
			"Adults":  eval.SelectorKey(3),
		},
		CompileOptions: map[eval.Option]bool{
			eval.ReportEvent:           true, // enable debug
			eval.AllowUnknownSelectors: true,

			eval.Reordering:      false,
			eval.FastEvaluation:  true,
			eval.ConstantFolding: true,
			eval.ReduceNesting:   true,
		},
		CostsMap: map[string]float64{
			//"Origin": 100,
			//"Value":  -100,
			//"Adults": -100,
		},
	}

	t := &tui.TerminalUI{
		Width: 150,

		Expr:          s,
		Config:        cc,
		Params:        vals,
		TryEval:       true,
		SkipEventNode: true,
	}

	t.Start()
}
