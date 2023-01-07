package main

import (
	"context"
	"fmt"
	"github.com/onheap/eval"
	"github.com/onheap/eval_lab/data/model"
	"github.com/onheap/eval_lab/data/rule"
)

const size = 10000

func main() {
	cc := eval.NewCompileConfig()

	cc.ConstantMap = rule.ConstantMap()
	cc.SelectorMap = rule.SelectorMap()
	cc.OperatorMap = rule.OperatorMap()

	rules, err := rule.LoadRules()
	if err != nil {
		panic(err)
	}

	exprs := make([]*eval.Expr, len(rules))
	for i, r := range rules {
		exprs[i], err = eval.Compile(cc, r)
		if err != nil {
			panic(err)
		}
	}

	g := model.NewGenerator(1)

	ctx := context.TODO()
	users := make([]*eval.Ctx, 0, size)
	for i := 0; i < size; i++ {
		users = append(users, rule.ToEvalCtx(ctx, g.GenUser()))
	}

	res := make([]int64, len(exprs))
	for i, expr := range exprs {
		for _, user := range users {
			b, err := expr.EvalBool(user)
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
