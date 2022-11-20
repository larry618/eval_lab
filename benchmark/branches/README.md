# Benchmark Branches

The tests in this folder are to compare the performance of local version with the latest remote version.

## Usage

Get the latest commits from the main branch
```
go get github.com/onheap/eval@main

or

go get github.com/onheap/eval@commit_id
```

Run benchmark tests
```bash
go test -bench=. -run=none -benchtime=3s -benchmem
go test -bench='BenchmarkEval' -run=none -benchtime=3s -benchmem

go test -bench='BenchmarkEvalDev' -run=none -benchtime=3s -benchmem -memprofile profile.out
go tool pprof -http=:8080 profile.out
```

Code optimization (in the [project eval](https://github.com/onheap/eval) directory)
```bash
# escape analysis
go build -gcflags="-m -m" 2>&1 | grep "engine"

# SSA
env "GOSSAFUNC=(*Expr).Eval" go build

# struct layout
structlayout -json ./engine.go node | structlayout-pretty
structlayout -json ./engine.go node | structlayout-optimize
```

