package optimizer

import (
	"github.com/c-bata/goptuna"
	"github.com/c-bata/goptuna/cmaes"
	"github.com/c-bata/goptuna/rdb.v2"
	"github.com/c-bata/goptuna/tpe"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GopTuna(executor *Executor) ([]float64, float64) {
	//study, err := createStudy()
	study, err := createCmaesStudyWithDB()
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

	err = study.Optimize(objective, 10000)
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

func createTpeStudyWithDB() (*goptuna.Study, error) {
	const dsn = "eval_db.sqlite3"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}
	err = rdb.RunAutoMigrate(db)
	if err != nil {
		return nil, err
	}

	return goptuna.CreateStudy(
		"eval_tpe",
		goptuna.StudyOptionLogger(nil),
		goptuna.StudyOptionStorage(rdb.NewStorage(db)),
		goptuna.StudyOptionSampler(tpe.NewSampler()),
		goptuna.StudyOptionDirection(goptuna.StudyDirectionMinimize),
		goptuna.StudyOptionLoadIfExists(true),
	)
}

func createCmaesStudyWithDB() (*goptuna.Study, error) {
	const dsn = "eval_db.sqlite3"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}
	err = rdb.RunAutoMigrate(db)
	if err != nil {
		return nil, err
	}

	cmaesSampler := cmaes.NewSampler(
		cmaes.SamplerOptionNStartupTrials(5))

	return goptuna.CreateStudy(
		"eval_cmaes",
		//goptuna.StudyOptionLogger(nil),
		goptuna.StudyOptionStorage(rdb.NewStorage(db)),
		goptuna.StudyOptionRelativeSampler(cmaesSampler),
		goptuna.StudyOptionLoadIfExists(true),
	)
}

func createStudy() (*goptuna.Study, error) {
	_ = tpe.NewSampler()
	cmaesSampler := cmaes.NewSampler(
		cmaes.SamplerOptionNStartupTrials(5))

	return goptuna.CreateStudy(
		"goptuna",
		goptuna.StudyOptionLogger(nil),
		//goptuna.StudyOptionSampler(tpeSampler),
		goptuna.StudyOptionRelativeSampler(cmaesSampler),
	)
}
