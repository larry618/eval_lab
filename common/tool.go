package common

import (
	"encoding/json"
	"fmt"
)

func PrintJson(v interface{}) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
