package main

import (
	"context"
	"fmt"
	"github.com/MaxHalford/eaopt"
	"github.com/onheap/eval_lab/common"
	"github.com/onheap/eval_lab/costs"
	"github.com/onheap/eval_lab/data/model"
	"github.com/onheap/eval_lab/data/rule"
	"local/eval"
)

const size = 10000

func main() {
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

	config.CostsMap = map[string]float64{
		"address":         -0.2673894241818548,
		"address.city":    -1.8902959372128334,
		"address.country": 2.3333135648838343,
		"address.state":   6.9171105028635065,
		"age":             -0.8637966451577009,
		"app_version":     -0.5899768958649574,
		"balance":         -316.0174344562734,
		"birth_date":      92.38754463554814,
		"created_at":      247.830395062358,
		"credit":          0.5078775146828337,
		"credit_limit":    -55.6878709902171,
		"discount":        442.6481166781716,
		"distance":        2.2207711628575257,
		"gender":          -4.797643104628769,
		"interests":       -50.102679219089566,
		"is_birthday":     -2.189705588154541,
		"is_student":      -0.003908537785935158,
		"is_vip":          0.04995903444326317,
		"language":        -122.3751203353385,
		"now":             889.1232423530479,
		"os_version":      50.515414415518435,
		"platform":        1.3150028957337696,
		"updated_at":      12.899403383756011,
		"user_id":         -8.520977119183572,
		"user_tags":       0.975682156603751,
	}

	executor := costs.NewExecutor(config, rules, ctxes)
	initCosts := executor.GetInitCosts()

	count, err := executor.Exec(initCosts)
	if err != nil {
		panic(err)
	}

	fmt.Println("initial execution count:", count)

	o := &costs.GAOptimizer{
		Executor: executor,
		Costs:    initCosts,
	}

	ga, err := eaopt.NewDefaultGAConfig().NewGA()
	if err != nil {
		panic(err)
	}

	// Set the number of generations to run for
	ga.NGenerations = 100

	// Add a custom print function to track progress
	ga.Callback = o.Callback

	// Find the minimum
	err = ga.Minimize(o.Factory)
	if err != nil {
		panic(err)
	}

	finalCosts := ga.HallOfFame[0].Genome.(*costs.GAOptimizer).Costs
	common.PrintJson(executor.CostsMap(finalCosts))
}
