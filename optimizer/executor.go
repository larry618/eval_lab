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
	ctxes []*eval.Ctx) (*Executor, error) {

	e := &Executor{
		cc:    cc,
		rules: rules,
		ctxes: ctxes,
	}

	err := e.initIndexMap()
	return e, err
}

func (e *Executor) initIndexMap() error {
	indexMap := make(map[string]int)
	for s := range e.cc.SelectorMap {
		indexMap[s] = len(indexMap)
	}

	for s := range e.cc.OperatorMap {
		indexMap[s] = len(indexMap)
	}

	e.indexMap = indexMap

	err := e.initCostIdentifiers()
	return err
}

func (e *Executor) initCostIdentifiers() error {
	for _, rule := range e.rules {
		expr, err := eval.Compile(e.cc, rule)
		if err != nil {
			return err
		}
		tree := eval.GetAstTree(expr)

		for _, n := range tree {
			if n.ParentIdx == -1 {
				continue
			}

			p := tree[n.ParentIdx]
			if !isBoolOp(p) {
				continue
			}

			switch n.NodeType {
			case eval.SelectorNode, eval.OperatorNode:
				s := getCostIdentifier(p, n, nil)
				if _, exist := e.indexMap[s]; !exist {
					e.indexMap[s] = len(e.indexMap)
				}
			case eval.FastOperatorNode:
				s := getCostIdentifier(p, n, nil)
				if _, exist := e.indexMap[s]; !exist {
					e.indexMap[s] = len(e.indexMap)
				}

				cIdx := n.ChildIdx
				cCnt := n.ChildCnt
				s = getCostIdentifier(p, n, tree[cIdx:cIdx+cCnt])
				if _, exist := e.indexMap[s]; !exist {
					e.indexMap[s] = len(e.indexMap)
				}
			}
		}
	}

	missed := []string{
		`(and (!= address.country "US"))`,
		`(and (= credit 3))`,
		`(and (= gender 1))`,
		`(and (= gender 2))`,
		`(and (> balance 3000))`,
		`(and (> discount 88))`,
		`(and (> updated_at 1162166400))`,
		`(and (>= age 21))`,
		`(and (>= age 22))`,
		`(and (>= age 30))`,
		`(and (>= balance 1000))`,
		`(and (>= credit 2))`,
		`(and (>= credit_limit 3000))`,
		`(and (>= credit_limit 4000))`,
		`(and (in "active" user_tags))`,
		`(and (in "return" user_tags))`,
		`(and (in "sports" interests))`,
		`(and (in "top" user_tags))`,
		`(and (in "video_games" interests))`,
		`(and (in language ("zh-CN" "zh-HK" "zh-TW")))`,
		`(and (overlap ("active" "new") user_tags))`,
		`(and (overlap ("high_value" "top") user_tags))`,
		`(and (overlap ("video_games" "travel") interests))`,
		`(and is_student)`,
		`(and is_vip)`,
		`(or (= address.country "CA"))`,
		`(or (= address.country "US"))`,
		`(or (= credit 3))`,
		`(or (>= balance 4000))`,
		`(or (>= credit 1))`,
		`(or (>= credit 2))`,
		`(or (>= credit_limit 4000))`,
		`(or (in "celebrity" user_tags))`,
		`(or (in "top" user_tags))`,
		`(or (in "video_games" interests))`,
		`(or (overlap ("top" "high_value") user_tags))`,
	}

	for _, s := range missed {
		if _, exist := e.indexMap[s]; !exist {
			e.indexMap[s] = len(e.indexMap)
		}
	}

	return nil
}

func isBoolOp(node eval.TreeNode) bool {
	nt, v := node.NodeType, node.Value
	return (nt == eval.OperatorNode || nt == eval.FastOperatorNode) &&
		(v == "or" || v == "|" || v == "||" || v == "and" || v == "&" || v == "&&")
}

func getCostIdentifier(parent eval.TreeNode, curt eval.TreeNode, children []eval.TreeNode) string {
	var tree eval.Tree
	var appendChildren = func(parentIdx int, children ...eval.TreeNode) {
		if parentIdx != -1 {
			tree[parentIdx].ChildCnt = len(children)
			tree[parentIdx].ChildIdx = len(tree)
		}

		for _, n := range children {
			tree = append(tree,
				eval.TreeNode{
					NodeType: n.NodeType,
					Value:    n.Value,

					Idx:       len(tree),
					ChildCnt:  0,
					ChildIdx:  -1,
					ParentIdx: parentIdx,
				},
			)
		}
	}

	appendChildren(-1, parent)
	appendChildren(0, curt)
	if len(children) != 0 {
		appendChildren(1, children...)
	}

	code := tree.DumpCode(false)
	//fmt.Printf("gen `%s`\n", code)
	return code
}

func (e *Executor) GetInitCosts() []float64 {
	initCosts := make([]float64, len(e.indexMap))
	for s, i := range e.indexMap {
		c := e.cc.CostsMap[s]
		initCosts[i] = c
	}
	return initCosts
}

func (e *Executor) ToCostsMap(costs []float64) map[string]float64 {
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
