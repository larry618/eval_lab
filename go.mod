module github.com/onheap/eval_lab

go 1.18

require (
	github.com/antonmedv/expr v1.9.0
	github.com/google/cel-go v0.12.4
	github.com/onheap/eval v0.0.0-20220726112430-f31dc2d0bae9
    local/eval v1.0.0
)

require (
	github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20220418222510-f25a4f6275ed // indirect
	github.com/stoewer/go-strcase v1.2.0 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220502173005-c8bf987b8c21 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace local/eval v1.0.0 => ../eval
