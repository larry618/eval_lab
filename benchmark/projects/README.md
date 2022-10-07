# Benchmark Projects


The tests in this folder are cloned from the repo [golang-expression-evaluation-comparison](https://github.com/antonmedv/golang-expression-evaluation-comparison) to compare the performance of different projects.


## Usage
```
go test -bench=. -run=none -benchtime=3s -benchmem
```