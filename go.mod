module github.com/larry618/eval_lab

require (
	github.com/antonmedv/expr v1.8.9
	github.com/google/cel-go v0.7.3
	github.com/larry618/eval v0.0.0-20220724121141-07c93153366f
	local/eval v1.0.0
)

require (
	github.com/antlr/antlr4 v0.0.0-20200503195918-621b933c7a7f // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/stoewer/go-strcase v1.2.0 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20201102152239-715cce707fb0 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

replace local/eval v1.0.0 => ../eval

go 1.18
