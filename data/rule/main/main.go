package main

import (
	"context"
	"fmt"
	"local/eval"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/onheap/eval_lab/data/model"
	"github.com/onheap/eval_lab/data/rule"
)

const size = 10000

func main() {
	cc := eval.NewCompileConfig()

	cc.ConstantMap = rule.ConstantMap()
	cc.SelectorMap = rule.SelectorMap()
	cc.OperatorMap = rule.OperatorMap()

	rules, err := rule.LoadAndCompileRules(cc)
	if err != nil {
		panic(err)
	}

	f := gofakeit.Faker{Rand: rand.New(rand.NewSource(1))}

	ctx := context.TODO()
	users := make([]*eval.Ctx, 0, size)
	for i := 0; i < size; i++ {
		users = append(users, rule.ToEvalCtx(ctx, model.GenUser(f)))
	}

	res := make([]int64, len(rules))
	for i, r := range rules {
		for _, user := range users {
			b, err := r.EvalBool(user)
			if err != nil {
				panic(err)
			}
			if b {
				res[i]++
			}
		}
	}

	for i, v := range res {
		fmt.Println(i, ":", v)
	}
}
