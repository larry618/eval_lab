package optimizer

import (
	"dev/eval"
	"sync"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

type Executor struct {
	cc *eval.CompileConfig

	rules []string
	ctxes []*eval.Ctx

	indexMap map[string]int
}

func NewExecutor(
	cc *eval.CompileConfig,
	rules []string,
	ctxes []*eval.Ctx) *Executor {

	o := &Executor{
		cc:       cc,
		rules:    rules,
		ctxes:    ctxes,
		indexMap: make(map[string]int),
	}

	for s := range cc.SelectorMap {
		o.indexMap[s] = len(o.indexMap)
	}

	for s := range cc.OperatorMap {
		o.indexMap[s] = len(o.indexMap)
	}

	// todo: add unknown selector
	return o
}

func (e *Executor) GetInitCosts() []float64 {
	initCosts := make([]float64, len(e.indexMap))
	for s, i := range e.indexMap {
		c := e.cc.CostsMap[s]
		initCosts[i] = c
	}
	return initCosts
}

func (e *Executor) CostsMap(costs []float64) map[string]float64 {
	res := make(map[string]float64, len(e.indexMap))
	for s, i := range e.indexMap {
		res[s] = costs[i]
	}
	return res
}

func (e *Executor) Exec(costs []float64) (float64, error) {
	t := e.newTask(costs)
	err := t.init()
	if err != nil {
		return 0, err
	}

	err = t.exec()
	if err != nil {
		return 0, err
	}

	t.clean()
	return float64(t.counter), nil
}

func (e *Executor) newTask(costs []float64) *task {
	cc := eval.CopyCompileConfig(e.cc)
	cc.CompileOptions[eval.ReportEvent] = true
	cc.CostsMap = e.CostsMap(costs)
	return &task{
		cc:    cc,
		rules: e.rules,
		ctxes: e.ctxes,
	}
}

type task struct {
	counter int64
	wg      sync.WaitGroup
	rules   []string
	ctxes   []*eval.Ctx

	cc    *eval.CompileConfig
	exprs []*eval.Expr
}

func (t *task) init() error {
	t.exprs = make([]*eval.Expr, len(t.rules))
	for i, rule := range t.rules {
		expr, err := eval.Compile(t.cc, rule)
		if err != nil {
			return err
		}

		t.wg.Add(1)
		expr.EventChan = make(chan eval.Event, 1024)
		t.stats(expr)

		t.exprs[i] = expr
	}

	return nil
}

func (t *task) exec() error {
	var eg errgroup.Group
	for i := range t.exprs {
		expr := t.exprs[i]
		eg.Go(func() error {
			for _, ctx := range t.ctxes {
				_, err := expr.Eval(ctx)
				if err != nil {
					return err
				}
			}
			return nil
		})
	}
	return eg.Wait()
}

func (t *task) stats(expr *eval.Expr) {
	go func(e *eval.Expr) {
		var count int64
		for event := range expr.EventChan {
			if event.EventType == eval.LoopEvent {
				count++
			}
		}
		atomic.AddInt64(&t.counter, count)
		t.wg.Done()
	}(expr)
}

func (t *task) clean() {
	for _, e := range t.exprs {
		close(e.EventChan)
	}
	t.wg.Wait()
}
