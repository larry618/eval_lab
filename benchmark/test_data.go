package benchmark

const Example = `(Origin == "MOW" || Country == "RU") && (Value >= 100 || Adults == 1)`

func CreateParams() map[string]interface{} {
	return map[string]interface{}{
		"Origin":  "MOW",
		"Country": "RU",
		"Adults":  1,
		"Value":   100,
	}
}
