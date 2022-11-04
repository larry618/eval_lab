package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/onheap/eval"
	"strings"
)

func main() {

	s := `
;;;; optimize:false
(and
 (or
   (= Origin "MOW") (= Country "RU"))  
 (or (>= Value 100) (= Adults 1)))
`

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

			eval.Reordering:      true,
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

	tui(cc, s, vals, false)
}

func tui(cc *eval.CompileConfig, s string, params map[string]interface{}, tryEval bool) {
	s = eval.IndentByParentheses(s)

	eventChan := make(chan eval.Event)

	ctx := eval.NewCtxWithMap(cc, params)

	expr, err := eval.Compile(cc, s)
	if err != nil {
		panic(err)
	}
	expr.EventChan = eventChan

	if err = ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	const width = 170

	var height = strings.Count(s, "\n") + 3

	leftGraph := widgets.NewParagraph()
	leftGraph.Title = "Original Expression"
	leftGraph.Text = s
	leftGraph.SetRect(0, 0, width/2, height)

	rightGraph := widgets.NewParagraph()
	rightGraph.Title = "Compiled Expression"
	rightGraph.Text = eval.Dump(expr)
	rightGraph.SetRect(width/2+1, 0, width, height)

	tableGraph := widgets.NewParagraph()
	tableGraph.Title = "Table"
	tableGraph.Text = eval.DumpTable(expr, true)
	tableGraph.SetRect(0, height+1, width, height+14)
	height += 15

	eventsList := widgets.NewList()
	eventsList.Title = "Execution Events"
	eventsList.WrapText = false
	eventsList.TextStyle = ui.NewStyle(ui.ColorWhite)
	eventsList.SelectedRowStyle = ui.NewStyle(ui.ColorRed)
	eventsList.SetRect(0, height+1, width-50, height+15)
	//height += 10

	stackGraph := widgets.NewParagraph()
	stackGraph.Title = "Stack"
	stackGraph.PaddingLeft = 25
	stackGraph.SetRect(width-50+1, height+1, width, height+15)

	ui.Render(leftGraph, rightGraph, tableGraph, eventsList, stackGraph)

	var res eval.Value
	go func() {
		if tryEval {
			res, err = expr.TryEval(ctx)
		} else {
			res, err = expr.Eval(ctx)
		}

		if err != nil {
			panic(err)
		}
		close(expr.EventChan)
	}()

	var prev eval.Event
	for uiEvent := range ui.PollEvents() {
		switch uiEvent.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			eventsList.ScrollDown()
		case "k", "<Up>":
			eventsList.ScrollUp()
		case "<Enter>":
			ev, ok := <-eventChan
			if !ok {
				stackGraph.Text = fmt.Sprintf("Final result: %v", res)
				ui.Render(stackGraph)
				continue
			}
			switch ev.EventType {
			case eval.OpExecEvent, eval.FastOpExecEvent:
				data := ev.Data.(eval.OpEventData)
				row := fmt.Sprintf(
					"[%3d] Exec Operator: op: %s, params: %v, res: %v, err: %v",
					len(eventsList.Rows), data.OpName, data.Params, data.Res, data.Err)
				eventsList.Rows = append(eventsList.Rows, row)
			case eval.LoopEvent:
				var curt = ev.Data

				if ev.CurtIdx-prev.CurtIdx > 2 {
					eventsList.Rows = append(eventsList.Rows, fmt.Sprintf("[[%3d] Short Circuit: [%v] jump to [%v]](bg:green)", len(eventsList.Rows), prev.Data, curt))
				}

				eventsList.Rows = append(eventsList.Rows, fmt.Sprintf("[%3d] Current Node: [%v], type:[%s], idx:[%d] ", len(eventsList.Rows), curt, ev.NodeType.String(), ev.CurtIdx))

				stackGraph.Text = drawStack(ev.Stack)
				prev = ev
			}

			eventsList.SelectedRow = len(eventsList.Rows) - 1
		}

		ui.Render(eventsList, stackGraph)
	}
	return
}

func drawStack(stack []eval.Value) string {
	var sb strings.Builder
	sb.WriteString("\n\n")
	for i := 0; i < 8-len(stack); i++ {
		sb.WriteString("│       │\n")
	}

	for i := len(stack) - 1; i >= 0; i-- {
		sb.WriteString(fmt.Sprintf("│ %5v │\n", stack[i]))
	}

	sb.WriteString("└───────┘")
	return sb.String()
}
