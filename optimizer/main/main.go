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
		"address":         -40.63870316206402,
		"address.city":    0.0710227058252233,
		"address.country": 30.210058686854367,
		"address.state":   -0.02823716811558348,
		"age":             0.40326041397793244,
		"app_version":     19.667577682778045,
		"balance":         -8.363933544396618,
		"birth_date":      1.0643036802851924,
		"created_at":      31.825885259849215,
		"credit":          -30.165281458972295,
		"credit_limit":    -314.549068477568,
		"discount":        0.057350554961002276,
		"distance":        3.7300158761369175,
		"gender":          -4.863874209068786,
		"interests":       -3.524997444603568,
		"is_birthday":     38.04945055128983,
		"is_student":      -0.02783173895578729,
		"is_vip":          34.92114795250204,
		"language":        -5044.663266113189,
		"now":             8.514616864224068,
		"os_version":      5.04653449094668,
		"platform":        27.22224483876182,
		"updated_at":      24.830042105193144,
		"user_id":         10.27928489016543,
		"user_tags":       -0.7498110735744463,
	}

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
	ga.NGenerations = 1000

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
