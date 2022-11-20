package tui

import (
	"dev/eval"
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"strings"
	"time"
)

func (tui *TerminalUI) Start() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	var (
		ticker   = time.NewTicker(time.Second).C
		uiEvents = ui.PollEvents()
	)

	tui.init()

	for {
		select {
		case uiEvent := <-uiEvents:
			switch uiEvent.ID {
			case "q", "<C-c>":
				return
			case "r":
				tui.init()
			case "j", "<Down>":
				tui.eventsList.ScrollDown()
			case "k", "<Up>":
				tui.eventsList.ScrollUp()
			case "<Enter>":
				tui.handleExecEvent()
			}
		case <-ticker:
		}
		tui.render()
	}

}

type TerminalUI struct {
	Width int

	Expr          string
	Config        *eval.CompileConfig
	Params        map[string]interface{}
	TryEval       bool
	SkipEventNode bool

	ctx  *eval.Ctx
	expr *eval.Expr
	res  eval.Value

	leftGraph  *widgets.Paragraph
	rightGraph *widgets.Paragraph

	tableGraph *widgets.Paragraph
	eventsList *widgets.List
	stackGraph *widgets.Paragraph

	events []string
	prev   eval.LoopEventData

	exprTable string
	indexRow  string
}

func (tui *TerminalUI) initExpr() {
	tui.Expr = eval.IndentByParentheses(tui.Expr)

	tui.ctx = eval.NewCtxWithMap(tui.Config, tui.Params)

	expr, err := eval.Compile(tui.Config, tui.Expr)
	if err != nil {
		panic(err)
	}

	expr.EventChan = make(chan eval.Event)
	tui.expr = expr
	tui.exprTable = eval.DumpTable(expr, tui.SkipEventNode)
	tui.indexRow = strings.Split(tui.exprTable, "\n")[2]
}

func (tui *TerminalUI) initTui() {
	width := tui.Width
	if width == 0 {
		width = 150
	}

	var height = strings.Count(tui.Expr, "\n") + 3

	leftGraph := widgets.NewParagraph()
	leftGraph.Title = "Original Expression"
	leftGraph.Text = tui.Expr
	leftGraph.SetRect(0, 0, width/2, height)

	rightGraph := widgets.NewParagraph()
	rightGraph.Title = "Compiled Expression"
	rightGraph.Text = eval.Dump(tui.expr)
	rightGraph.SetRect(width/2+1, 0, width, height)

	tableGraph := widgets.NewParagraph()
	tableGraph.Title = "Table"
	tableGraph.Text = tui.exprTable + strings.Repeat(" ", 4) + "▲"
	tableGraph.SetRect(0, height+1, width, height+15)
	height += 15

	eventsList := widgets.NewList()
	eventsList.Title = "Execution Events"
	eventsList.Rows = []string{
		"Press <Enter> to start execution",
		"Press <j> or <k> to scroll",
		"Press <r> to restart",
		"Press <q> to quit",
	}
	eventsList.WrapText = false
	eventsList.TextStyle = ui.NewStyle(ui.ColorWhite)
	eventsList.SelectedRowStyle = ui.NewStyle(ui.ColorRed)
	eventsList.SetRect(0, height+1, width-50, height+15)

	stackGraph := widgets.NewParagraph()
	stackGraph.Title = "Stack"
	stackGraph.PaddingLeft = 25
	stackGraph.Text = drawStack(nil)
	stackGraph.SetRect(width-50+1, height+1, width, height+15)

	tui.leftGraph = leftGraph
	tui.rightGraph = rightGraph
	tui.tableGraph = tableGraph
	tui.eventsList = eventsList
	tui.stackGraph = stackGraph

	tui.render()
}

func (tui *TerminalUI) init() {
	ui.Clear()
	tui.prev = eval.LoopEventData{}
	tui.events = nil
	tui.initExpr()
	tui.initTui()
	tui.execute()
}

func (tui *TerminalUI) render() {
	ui.Render(tui.leftGraph, tui.rightGraph, tui.tableGraph, tui.eventsList, tui.stackGraph)
}

func (tui *TerminalUI) execute() {
	go func() {
		var err error
		var res eval.Value
		if tui.TryEval {
			res, err = tui.expr.TryEval(tui.ctx)
		} else {
			res, err = tui.expr.Eval(tui.ctx)
		}

		if err != nil {
			panic(err)
		}

		tui.res = res
		close(tui.expr.EventChan)
	}()
}

func (tui *TerminalUI) handleExecEvent() bool {
	events := tui.events

	ev, ok := <-tui.expr.EventChan
	if !ok {
		text := fmt.Sprintf("Final result: %v", tui.res)
		text += strings.Repeat("\n", 10)
		text += "[Press <r> to restart\nPress <q> to quit](fg:red)"
		tui.stackGraph.Text = text
		return true
	}
	switch ev.EventType {
	case eval.OpExecEvent:
		data := ev.Data.(eval.OpEventData)
		row := fmt.Sprintf(
			"[%3d] [Exec Operator](bg:green): op: %s, params: %v, res: %v, err: %v",
			len(events), data.OpName, data.Params, data.Res, data.Err)
		events = append(events, row)
	case eval.LoopEvent:
		var (
			prev    = tui.prev
			curt    = ev.Data.(eval.LoopEventData)
			curtIdx = int(curt.CurtIdx)
		)

		idx := strings.Index(tui.indexRow, fmt.Sprintf("|%5d|", curtIdx)) + 3
		cursorRow := strings.Repeat(" ", idx) + "▲"
		tui.tableGraph.Text = tui.exprTable + cursorRow

		var minSteps int16 = 2
		if prev.NodeType == eval.FastOperatorNode {
			minSteps = 4
		}

		if steps := curt.CurtIdx - prev.CurtIdx; steps > minSteps {
			events = append(events, fmt.Sprintf("[%3d] [Short Circuit](bg:red): [%v] jump to [%v] ", len(events), prev.NodeValue, curt.NodeValue))
		}

		events = append(events, fmt.Sprintf("[%3d] Execute Node: [%v], type:[%s], idx:[%d] ", len(events), curt.NodeValue, curt.NodeType.String(), curtIdx))

		tui.stackGraph.Text = drawStack(ev.Stack)
		tui.prev = curt
	}

	tui.eventsList.Rows = events
	tui.eventsList.SelectedRow = len(events) - 1
	tui.events = events

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
