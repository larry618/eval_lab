package main

import (
	"context"
	"fmt"
	"github.com/MaxHalford/eaopt"
	"github.com/onheap/eval_lab/common"
	"github.com/onheap/eval_lab/data/model"
	"github.com/onheap/eval_lab/data/rule"
	"github.com/onheap/eval_lab/optimizer"
	"local/eval"
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

	executor := optimizer.NewExecutor(config, rules, ctxes)
	initCosts := executor.GetInitCosts()

	count, err := executor.Exec(initCosts)
	if err != nil {
		panic(err)
	}

	fmt.Println("initial execution count:", count)

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
	common.PrintJson(executor.CostsMap(finalCosts))
}
