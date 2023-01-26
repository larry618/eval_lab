package optimizer

import (
	"sync"
	"sync/atomic"

	"github.com/onheap/eval"
	"golang.org/x/sync/errgroup"
)

type Executor struct {
	cc *eval.Config

	rules []string
	ctxes []*eval.Ctx

	indexMap map[string]int

	Callback func(costs []float64, score float64)
}

func NewExecutor(
	cc *eval.Config,
	rules []string,
	ctxes []*eval.Ctx) (*Executor, error) {

	e := &Executor{
		cc:    cc,
		rules: rules,
		ctxes: ctxes,
	}

	err := e.initIndexMap()
	return e, err
}

func (e *Executor) initIndexMap() (err error) {
	indexMap := make(map[string]int)
	for s := range e.cc.VariableKeyMap {
		indexMap[s] = len(indexMap)
	}

	for s := range e.cc.OperatorMap {
		indexMap[s] = len(indexMap)
	}

	e.indexMap = indexMap
	return nil
}

func (e *Executor) GetInitCosts() []float64 {
	initCosts := make([]float64, len(e.indexMap))
	for s, i := range e.indexMap {
		c := e.cc.CostsMap[s]
		initCosts[i] = c
	}
	return initCosts
}

func (e *Executor) GetIndexMap() map[string]int {
	return e.indexMap
}

func (e *Executor) ToCostsMap(costs []float64) map[string]float64 {
	res := make(map[string]float64, len(e.indexMap))
	for s, i := range e.indexMap {
		res[s] = costs[i]
	}
	return res
}

func (e *Executor) ExecPanicErr(costs []float64) float64 {
	res, err := e.Exec(costs)
	if err != nil {
		panic(err)
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
	res := float64(t.counter)

	if e.Callback != nil {
		e.Callback(costs, res)
	}
	return res, nil
}

func (e *Executor) newTask(costs []float64) *task {
	cc := eval.CopyConfig(e.cc)
	cc.CompileOptions[eval.ReportEvent] = true
	cc.CostsMap = e.ToCostsMap(costs)
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

	cc    *eval.Config
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
