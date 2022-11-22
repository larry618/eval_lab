package main

import (
	"context"
	"dev/eval"
	"fmt"
	"github.com/onheap/eval_lab/data/model"
	"github.com/onheap/eval_lab/data/rule"
	"github.com/onheap/eval_lab/optimizer"
	"github.com/onheap/eval_lab/tool"
	"math"
	"sync/atomic"
)

func main() {
	executor := initRuleAndExecutor(true)

	initCosts := executor.GetInitCosts()

	count, err := executor.Exec(initCosts)
	if err != nil {
		panic(err)
	}

	fmt.Println("initial execution count:", count)

	finalCosts, min := optimizer.GopTuna(executor)

	fmt.Println("min", min)

	tool.PrintStringKeyMap(executor.ToCostsMap(finalCosts))
}

func initRuleAndExecutor(callback bool) *optimizer.Executor {
	const size = 10000

	rules, err := rule.LoadRules()
	if err != nil {
		panic(err)
	}

	g := model.NewGenerator(1)
	ctxes := make([]*eval.Ctx, 0, size)
	for i := 0; i < size; i++ {
		ctxes = append(ctxes, rule.ToEvalCtx(context.TODO(), g.GenUser()))
	}

	config := rule.CompileConfig()
	config.CompileOptions[eval.ContextBasedReordering] = true

	executor, err := optimizer.NewExecutor(config, rules, ctxes)
	if err != nil {
		panic(err)
	}

	if callback {
		var execution int64
		var min int64 = math.MaxInt64
		executor.Callback = func(_ []float64, score float64) {
			curt := int64(score)
			if curt < atomic.LoadInt64(&min) {
				atomic.StoreInt64(&min, curt)
			}

			atomic.AddInt64(&execution, 1)
			fmt.Printf(
				"[exec callback] No:%5d, curt: %6d, min: %6d\n",
				atomic.LoadInt64(&execution), curt, atomic.LoadInt64(&min))
		}
	}

	return executor
}
