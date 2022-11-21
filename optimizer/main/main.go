package main

import (
	"context"
	"dev/eval"
	"fmt"
	"github.com/MaxHalford/eaopt"
	"github.com/onheap/eval_lab/common"
	"github.com/onheap/eval_lab/data/model"
	"github.com/onheap/eval_lab/data/rule"
	"github.com/onheap/eval_lab/optimizer"
	"math"
	"math/rand"
	"sync/atomic"
)

func main() {
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

	initCosts := executor.GetInitCosts()

	count, err := executor.Exec(initCosts)
	if err != nil {
		panic(err)
	}

	fmt.Println("initial execution count:", count)

	finalCosts := deOpt(initCosts, executor)

	common.PrintStringKeyMap(executor.ToCostsMap(finalCosts))
}

func deOpt(initCosts []float64, executor *optimizer.Executor) []float64 {
	// Instantiate DiffEvo
	//var de, err = eaopt.NewDefaultDiffEvo()
	var de, err = eaopt.NewDiffEvo(50, 100, -5, 5, 0.5, 0.2, false, nil)
	if err != nil {
		panic(err)
	}

	// Fix random number generation
	de.GA.RNG = rand.New(rand.NewSource(42))

	// Run minimization
	res, min, err := de.Minimize(executor.ExecPanicErr, uint(len(initCosts)))
	if err != nil {
		panic(err)
	}

	fmt.Println("min score", min)
	return res
}

func oesOpt(initCosts []float64, executor *optimizer.Executor) []float64 {
	// Instantiate OpenAI Evolution Strategy
	//var oes, err = eaopt.NewDefaultOES()
	var oes, err = eaopt.NewOES(100, 200, 1, 0.1, false, nil)

	if err != nil {
		panic(err)
	}

	// Fix random number generation
	oes.GA.RNG = rand.New(rand.NewSource(42))

	// Run minimization
	res, min, err := oes.Minimize(executor.ExecPanicErr, initCosts)
	if err != nil {
		panic(err)
	}

	fmt.Println("min score", min)

	return res
}

func gaOpt(initCosts []float64, executor *optimizer.Executor) []float64 {
	o := &optimizer.GAOptimizer{
		Executor: executor,
		Costs:    initCosts,
	}

	ga, err := eaopt.NewDefaultGAConfig().NewGA()
	if err != nil {
		panic(err)
	}

	// Set the number of generations to run for
	ga.NGenerations = 10

	// Add a custom print function to track progress
	ga.Callback = o.Callback

	// Find the minimum
	err = ga.Minimize(o.Factory)
	if err != nil {
		panic(err)
	}

	finalCosts := ga.HallOfFame[0].Genome.(*optimizer.GAOptimizer).Costs
	return finalCosts
}
