package optimizer

import (
	"github.com/c-bata/goptuna"
	"github.com/c-bata/goptuna/cmaes"
	"github.com/c-bata/goptuna/tpe"
)

func GopTuna(executor *Executor) ([]float64, float64) {
	_ = tpe.NewSampler()
	cmaesSampler := cmaes.NewSampler(
		cmaes.SamplerOptionNStartupTrials(5))

	study, err := goptuna.CreateStudy(
		"goptuna",
		goptuna.StudyOptionLogger(nil),
		//goptuna.StudyOptionSampler(tpeSampler),
		goptuna.StudyOptionRelativeSampler(cmaesSampler),
	)
	if err != nil {
		panic(err)
	}

	indexMap := executor.GetIndexMap()

	var objective goptuna.FuncObjective = func(trial goptuna.Trial) (float64, error) {

		var params = make([]float64, len(indexMap))

		for key, idx := range indexMap {
			params[idx], _ = trial.SuggestFloat(key, -100, 100)
		}

		return executor.Exec(params)
	}

	err = study.Optimize(objective, 20000)
	if err != nil {
		panic(err)
	}

	v, _ := study.GetBestValue()
	p, _ := study.GetBestParams()

	var params = make([]float64, len(indexMap))
	for s, i := range indexMap {
		params[i] = p[s].(float64)
	}

	return params, v
}
