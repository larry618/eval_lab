package costs

import "github.com/MaxHalford/eaopt"

func test() {
	ga, err := eaopt.NewDefaultGAConfig().NewGA()

	if err != nil {
		panic(err)
	}
	_ = ga
}
