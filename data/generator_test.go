package data

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
	"testing"
)

func TestGenerate(t *testing.T) {
	f := gofakeit.Faker{Rand: rand.New(rand.NewSource(1))}

	for i := 0; i < 10; i++ {
		u := GenUser(f)
		printJson(u)
	}
}

func printJson(v interface{}) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
