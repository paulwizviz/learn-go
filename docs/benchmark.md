# Benchmark

This section discusses techniques to use benchmarking tool.

## Working example

This [example](../example/benchmark/benchmark_test.go) demonstrate a simple way of determining data types to support the getting of data from a collection at constant time.

To execute this example, `cd` into example/benchmark and run the command  `go test -bench=. -count=4`. This will ensure that each test case is run at least 4 times. You should get an output like this:

```
goos: <OS Type>
goarch: <Processor type>
pkg: <Go package under tests>
cpu: <CPU Type>
BenchmarkSlice/input_size_2-4           382399405                3.086 ns/op
BenchmarkSlice/input_size_2-4           383030371                3.085 ns/op
BenchmarkSlice/input_size_2-4           371197184                3.163 ns/op
BenchmarkSlice/input_size_2-4           385822216                3.092 ns/op
BenchmarkSlice/input_size_99-4          30049200                38.33 ns/op
BenchmarkSlice/input_size_99-4          30642824                38.44 ns/op
BenchmarkSlice/input_size_99-4          28422703                39.94 ns/op
BenchmarkSlice/input_size_99-4          30906968                38.24 ns/op
BenchmarkMap/input_size_2-4             185563402                6.492 ns/op
BenchmarkMap/input_size_2-4             185706861                6.464 ns/op
BenchmarkMap/input_size_2-4             179746323                6.462 ns/op
BenchmarkMap/input_size_2-4             185306089                6.471 ns/op
BenchmarkMap/input_size_99-4            176514410                6.777 ns/op
BenchmarkMap/input_size_99-4            176255048                6.782 ns/op
BenchmarkMap/input_size_99-4            177405297                6.762 ns/op
BenchmarkMap/input_size_99-4            165562826                6.774 ns/op
PASS
ok      github.com/paulwizviz/learn-go/example/benchmark        26.009s
```

## Refences

* [Benchmarking in Golang: Improving function performance](https://blog.logrocket.com/benchmarking-golang-improve-function-performance/)