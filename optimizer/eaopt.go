package optimizer

import (
	"fmt"
	"github.com/MaxHalford/eaopt"
	"github.com/onheap/eval_lab/tool"
	"math/rand"
)

type gaOptimizer struct {
	Executor *Executor
	Costs    []float64
}

func (o *gaOptimizer) Evaluate() (float64, error) {
	return o.Executor.Exec(o.Costs)
}

func (o *gaOptimizer) Mutate(rng *rand.Rand) {
	eaopt.MutNormalFloat64(o.Costs, 0.5, rng)
}

func (o *gaOptimizer) Crossover(genome eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossUniformFloat64(o.Costs, genome.(*gaOptimizer).Costs, rng)
}

func (o *gaOptimizer) Clone() eaopt.Genome {
	costs := make([]float64, len(o.Costs))
	copy(costs, o.Costs)
	return &gaOptimizer{
		Executor: o.Executor,
		Costs:    costs,
	}
}

func (o *gaOptimizer) Factory(rng *rand.Rand) eaopt.Genome {
	return &gaOptimizer{
		Executor: o.Executor,
		Costs:    eaopt.InitUnifFloat64(uint(len(o.Costs)), -100, 100, rng),
	}
}

func (o *gaOptimizer) Callback(ga *eaopt.GA) {
	fmt.Printf("Best fitness at generation %d: %f\n", ga.Generations, ga.HallOfFame[0].Fitness)
}

func DeOpt(executor *Executor) ([]float64, float64) {
	// Instantiate DiffEvo
	//var de, err = eaopt.NewDefaultDiffEvo()
	var de, err = eaopt.NewDiffEvo(50, 100, -5, 5, 0.5, 0.2, false, nil)
	tool.PanicErr(err)

	// Fix random number generation
	de.GA.RNG = rand.New(rand.NewSource(42))

	// Run minimization
	res, min, err := de.Minimize(executor.ExecPanicErr, uint(len(executor.GetInitCosts())))
	tool.PanicErr(err)

	return res, min
}

func OesOpt(executor *Executor) ([]float64, float64) {
	// Instantiate OpenAI Evolution Strategy
	//var oes, err = eaopt.NewDefaultOES()
	var oes, err = eaopt.NewOES(100, 200, 1, 0.1, false, nil)
	tool.PanicErr(err)

	// Fix random number generation
	oes.GA.RNG = rand.New(rand.NewSource(42))

	// Run minimization
	res, min, err := oes.Minimize(executor.ExecPanicErr, executor.GetInitCosts())
	tool.PanicErr(err)

	return res, min
}

func GaOpt(executor *Executor) ([]float64, float64) {
	o := &gaOptimizer{
		Executor: executor,
		Costs:    executor.GetInitCosts(),
	}

	ga, err := eaopt.NewDefaultGAConfig().NewGA()
	tool.PanicErr(err)

	// Set the number of generations to run for
	ga.NGenerations = 10

	// Add a custom print function to track progress
	ga.Callback = o.Callback

	// Find the minimum
	err = ga.Minimize(o.Factory)
	tool.PanicErr(err)

	best := ga.HallOfFame[0]
	finalCosts := best.Genome.(*gaOptimizer).Costs
	return finalCosts, best.Fitness
}
