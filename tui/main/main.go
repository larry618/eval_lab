package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/onheap/eval"
	"strings"
	"time"
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

	e := &tui{
		width: 150,

		s:       s,
		cc:      cc,
		params:  vals,
		tryEval: true,
	}

	e.start()
}

func (e *tui) start() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	var (
		ticker   = time.NewTicker(time.Second).C
		uiEvents = ui.PollEvents()
	)

	e.init()

	for {
		select {
		case uiEvent := <-uiEvents:
			switch uiEvent.ID {
			case "q", "<C-c>":
				return
			case "r":
				e.init()
			case "j", "<Down>":
				e.eventsList.ScrollDown()
			case "k", "<Up>":
				e.eventsList.ScrollUp()
			case "<Enter>":
				e.handleExecEvent()
			}
		case <-ticker:
		}
		e.render()
	}

}

type tui struct {
	width int

	tryEval bool
	s       string
	cc      *eval.CompileConfig
	params  map[string]interface{}

	exprStr string

	ctx  *eval.Ctx
	expr *eval.Expr
	res  eval.Value

	leftGraph  *widgets.Paragraph
	rightGraph *widgets.Paragraph

	tableGraph *widgets.Paragraph
	eventsList *widgets.List
	stackGraph *widgets.Paragraph

	events []string
	prev   eval.Event

	exprTable string
	indexRow  string
}

func (e *tui) initExpr() {
	e.exprStr = eval.IndentByParentheses(e.s)

	e.ctx = eval.NewCtxWithMap(e.cc, e.params)

	expr, err := eval.Compile(e.cc, e.exprStr)
	if err != nil {
		panic(err)
	}

	expr.EventChan = make(chan eval.Event)
	e.expr = expr
	e.exprTable = eval.DumpTable(expr, true)
	e.indexRow = strings.Split(e.exprTable, "\n")[2]
}

func (e *tui) initTui() {
	width := e.width
	if width == 0 {
		width = 150
	}

	var height = strings.Count(e.exprStr, "\n") + 3

	leftGraph := widgets.NewParagraph()
	leftGraph.Title = "Original Expression"
	leftGraph.Text = e.exprStr
	leftGraph.SetRect(0, 0, width/2, height)

	rightGraph := widgets.NewParagraph()
	rightGraph.Title = "Compiled Expression"
	rightGraph.Text = eval.Dump(e.expr)
	rightGraph.SetRect(width/2+1, 0, width, height)

	tableGraph := widgets.NewParagraph()
	tableGraph.Title = "Table"
	tableGraph.Text = e.exprTable + strings.Repeat(" ", 4) + "▲"
	tableGraph.SetRect(0, height+1, width, height+15)
	height += 15

	eventsList := widgets.NewList()
	eventsList.Title = "Execution Events"
	eventsList.Rows = []string{"Press <Enter> to start execution"}
	eventsList.WrapText = false
	eventsList.TextStyle = ui.NewStyle(ui.ColorWhite)
	eventsList.SelectedRowStyle = ui.NewStyle(ui.ColorRed)
	eventsList.SetRect(0, height+1, width-50, height+15)

	stackGraph := widgets.NewParagraph()
	stackGraph.Title = "Stack"
	stackGraph.PaddingLeft = 25
	stackGraph.Text = drawStack(nil)
	stackGraph.SetRect(width-50+1, height+1, width, height+15)

	e.leftGraph = leftGraph
	e.rightGraph = rightGraph
	e.tableGraph = tableGraph
	e.eventsList = eventsList
	e.stackGraph = stackGraph

	e.render()
}

func (e *tui) init() {
	ui.Clear()
	e.prev = eval.Event{}
	e.events = nil
	e.initExpr()
	e.initTui()
	e.execute()
}

func (e *tui) render() {
	ui.Render(e.leftGraph, e.rightGraph, e.tableGraph, e.eventsList, e.stackGraph)
}

func (e *tui) execute() {
	go func() {
		var err error
		var res eval.Value
		if e.tryEval {
			res, err = e.expr.TryEval(e.ctx)
		} else {
			res, err = e.expr.Eval(e.ctx)
		}

		if err != nil {
			panic(err)
		}

		e.res = res
		close(e.expr.EventChan)
	}()
}

func (e *tui) handleExecEvent() bool {
	events := e.events

	ev, ok := <-e.expr.EventChan
	if !ok {
		text := fmt.Sprintf("Final result: %v", e.res)
		text += strings.Repeat("\n", 10)
		text += "[Press <q> to quit\nPress <r> to restart](fg:red)"
		e.stackGraph.Text = text
		return true
	}
	switch ev.EventType {
	case eval.OpExecEvent, eval.FastOpExecEvent:
		data := ev.Data.(eval.OpEventData)
		row := fmt.Sprintf(
			"[%3d] [Exec Operator](bg:green): op: %s, params: %v, res: %v, err: %v",
			len(events), data.OpName, data.Params, data.Res, data.Err)
		events = append(events, row)
	case eval.LoopEvent:
		var (
			prev    = e.prev
			curt    = ev.Data
			curtIdx = int(ev.CurtIdx)
		)

		idx := strings.Index(e.indexRow, fmt.Sprintf("|%4d|", curtIdx)) + 3
		cursorRow := strings.Repeat(" ", idx) + "▲"
		e.tableGraph.Text = e.exprTable + cursorRow

		var minSteps int16 = 2
		if prev.NodeType == eval.FastOperatorNode {
			minSteps = 4
		}

		if steps := ev.CurtIdx - prev.CurtIdx; steps > minSteps {
			events = append(events, fmt.Sprintf("[%3d] [Short Circuit](bg:red): [%v] jump to [%v] ", len(events), prev.Data, curt))
		}

		events = append(events, fmt.Sprintf("[%3d] Execute Node: [%v], type:[%s], idx:[%d] ", len(events), curt, ev.NodeType.String(), curtIdx))

		e.stackGraph.Text = drawStack(ev.Stack)
		e.prev = ev
	}

	e.eventsList.Rows = events
	e.eventsList.SelectedRow = len(events) - 1
	e.events = events

	return false
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
