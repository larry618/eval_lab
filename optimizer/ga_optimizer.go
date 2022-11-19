package optimizer

import (
	"fmt"
	"github.com/MaxHalford/eaopt"
	"math/rand"
)

type GAOptimizer struct {
	Executor *Executor
	Costs    []float64
}

func (o *GAOptimizer) Evaluate() (float64, error) {
	return o.Executor.Exec(o.Costs)
}

func (o *GAOptimizer) Mutate(rng *rand.Rand) {
	eaopt.MutNormalFloat64(o.Costs, 0.5, rng)
}

func (o *GAOptimizer) Crossover(genome eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossUniformFloat64(o.Costs, genome.(*GAOptimizer).Costs, rng)
}

func (o *GAOptimizer) Clone() eaopt.Genome {
	costs := make([]float64, len(o.Costs))
	copy(costs, o.Costs)
	return &GAOptimizer{
		Executor: o.Executor,
		Costs:    costs,
	}
}

func (o *GAOptimizer) Factory(rng *rand.Rand) eaopt.Genome {
	return &GAOptimizer{
		Executor: o.Executor,
		Costs:    eaopt.InitUnifFloat64(uint(len(o.Costs)), -100, 100, rng),
	}
}

func (o *GAOptimizer) Callback(ga *eaopt.GA) {
	fmt.Printf("Best fitness at generation %d: %f\n", ga.Generations, ga.HallOfFame[0].Fitness)
}
