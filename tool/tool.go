package tool

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func PrintJson(v interface{}) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func PrintStringKeyMap[T any](m map[string]T) {

	keys := make([]string, 0, len(m))
	for s := range m {
		keys = append(keys, s)
	}

	sort.Strings(keys)

	var sb strings.Builder
	sb.WriteString("{\n")

	for _, s := range keys {
		sb.WriteString(fmt.Sprintf("        `%s`: %v\n", s, m[s]))
	}

	sb.WriteString("}")

	fmt.Println(sb.String())
}

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
