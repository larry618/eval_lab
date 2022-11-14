module github.com/onheap/eval_lab

go 1.18

require (
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/MaxHalford/eaopt v0.4.2
	github.com/PaesslerAG/gval v1.2.1
	github.com/antonmedv/expr v1.9.0
	github.com/brianvoe/gofakeit/v6 v6.19.0
	github.com/dop251/goja v0.0.0-20221025165401-cb5011b539fe
	github.com/gizak/termui/v3 v3.1.0
	github.com/google/cel-go v0.12.4
	github.com/gurkankaymak/hocon v1.2.7
	github.com/hashicorp/go-bexpr v0.1.11
	github.com/onheap/eval v1.0.0
	github.com/robertkrimen/otto v0.0.0-20221025135307-511d75fba9f8
	github.com/skx/evalfilter/v2 v2.1.19
	go.starlark.net v0.0.0-20221028183056-acb66ad56dd2
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
	local/eval v1.0.0
)

require (
	github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20220418222510-f25a4f6275ed // indirect
	github.com/dlclark/regexp2 v1.7.0 // indirect
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/mattn/go-runewidth v0.0.8 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/mitchellh/pointerstructure v1.2.1 // indirect
	github.com/nsf/termbox-go v0.0.0-20190121233118-02980233997d // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/stoewer/go-strcase v1.2.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220502173005-c8bf987b8c21 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
)

replace local/eval v1.0.0 => ../eval
