package model

import (
	"github.com/onheap/eval_lab/tool"
	"testing"
)

func TestGenerate(t *testing.T) {
	g := NewGenerator(1)

	for i := 0; i < 10; i++ {
		u := g.GenUser()
		tool.PrintJson(u)
	}
}
