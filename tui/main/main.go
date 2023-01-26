package main

import (
	"github.com/onheap/eval"
	"github.com/onheap/eval_lab/tui"
)

func main() {

	s := `
   (and
        (or
          (>= credit Good)   ;; cost 100
          (overlap ("top" "high_value") user_tags))
        (not
          (in "F" user_tags)))
	`

	//s := `(and (> Adults 1) (now))`

	//vals := benchmark.CreateParams()
	//cc := eval.NewCompileConfig(eval.RegisterVals(vals), eval.EnableReportEvent, eval.Optimizations(false))

	vals := map[string]interface{}{
		"Origin":  "MOW",
		"Country": "RU",
		"Value":   100,
		"Adults":  1,
		//"Value":   -100,
		//"Adults":  -1,
	}

	cc := &eval.Config{
		VariableKeyMap: map[string]eval.VariableKey{
			"Origin":  eval.VariableKey(0),
			"Country": eval.VariableKey(1),
			"Value":   eval.VariableKey(2),
			"Adults":  eval.VariableKey(3),
		},
		CompileOptions: map[eval.CompileOption]bool{
			eval.ReportEvent:            true, // enable debug
			eval.AllowUndefinedVariable: true,

			eval.Reordering:      true,
			eval.FastEvaluation:  true,
			eval.ConstantFolding: true,
			eval.ReduceNesting:   true,
		},
		CostsMap: map[string]float64{
			"credit": 100,
		},
		OperatorMap: map[string]eval.Operator{
			"now": nil,
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
