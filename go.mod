module github.com/onheap/eval_lab

go 1.18

require (
	github.com/MaxHalford/eaopt v0.4.2
	github.com/antonmedv/expr v1.9.0
	github.com/brianvoe/gofakeit/v6 v6.19.0
	github.com/google/cel-go v0.12.4
	github.com/gurkankaymak/hocon v1.2.7
	github.com/onheap/eval v0.0.0-20220913070844-19f42ba2a58c
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
	local/eval v1.0.0
)

require (
	github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20220418222510-f25a4f6275ed // indirect
	github.com/stoewer/go-strcase v1.2.0 // indirect
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220502173005-c8bf987b8c21 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace local/eval v1.0.0 => ../eval
