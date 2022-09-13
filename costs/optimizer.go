package costs

import (
	"local/eval"
	"sync"
	"sync/atomic"
)

const defaultCost = 7

type Optimizer struct {
	cc *eval.CompileConfig

	base map[string]int
	curt map[string]int

	rules []string
	ctxes []*eval.Ctx

	globalCounter int64

	exprs []*eval.Expr

	prev  []float64
	costs []float64
}

func NewOptimizer(
	cc *eval.CompileConfig,
	rules []string,
	ctxes []*eval.Ctx) (*Optimizer, error) {

	cc.CompileOptions[eval.ReportEvent] = true
	o := &Optimizer{
		cc:    cc,
		base:  cc.CostsMap,
		rules: rules,
		ctxes: ctxes,
		exprs: make([]*eval.Expr, len(rules)),
	}

	curt := make(map[string]int, len(o.base))
	for k, v := range o.base {
		curt[k] = v
	}

	// init keys
	for _, rule := range o.rules {
		expr, err := eval.Compile(cc, rule)
		if err != nil {
			return nil, err
		}
		for _, key := range eval.GetSelectorKeys(expr) {
			k := key.StrKey
			if _, exist := curt[k]; !exist {
				curt[k] = defaultCost
			}
			k = "+" + key.StrKey
			if _, exist := curt[k]; !exist {
				curt[k] = defaultCost
			}
			k = "-" + key.StrKey
			if _, exist := curt[k]; !exist {
				curt[k] = defaultCost
			}
		}
	}
	return o, nil
}

func (o *Optimizer) OptimizeCostsMap() {
	o.init()
	o.exec()

}

func (o *Optimizer) init() {
	o.globalCounter = 0
	o.cc.CostsMap = o.curt

	for i, rule := range o.rules {
		expr, err := eval.Compile(o.cc, rule)
		if err != nil {
			panic(err)
		}

		o.stats(expr)

		o.exprs[i] = expr
	}
}

func (o *Optimizer) exec() {
	for _, expr := range o.exprs {
		for _, ctx := range o.ctxes {
			_, err := expr.Eval(ctx)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (o *Optimizer) execParallel() {
	var wg sync.WaitGroup
	for _, expr := range o.exprs {
		wg.Add(1)
		go func(e *eval.Expr) {
			for _, ctx := range o.ctxes {
				_, err := e.Eval(ctx)
				if err != nil {
					panic(err)
				}
			}
			wg.Done()
		}(expr)
	}
	wg.Wait()
}

func (o *Optimizer) stats(e *eval.Expr) {
	go func(e *eval.Expr) {
		var count int64
		for event := range e.EventChan {
			if event.EventType == eval.LoopEvent {
				count++
			}
		}
		atomic.AddInt64(&o.globalCounter, count)
	}(e)
}
