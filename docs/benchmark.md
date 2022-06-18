# Benchmark

This section discusses techniques to use benchmarking tool. For further description the use of Go benchmarking tools, refer to these references:

* [Benchmarking in Golang: Improving function performance](https://blog.logrocket.com/benchmarking-golang-improve-function-performance/)

## BigO notation

Big O notation is used in Computer Science to describe the performance or complexity of an algorithm [A beginner's guide to Big O Notation](https://rob-bell.net/2009/06/a-beginners-guide-to-big-o-notation).

## O(1)

O(1) describe an algorithm that will always execute in the same time (or space -- i.e. memory) regardless of the size of input data. 

A benchmark working demonstrate a constant time O(1) operation to determine if a data exists in a map data type is found [here](../example/benchmark/benchmark_test.go).

Running the example you will find stats verifying contant time operation.

```
BenchmarkO1/input_1-4         	176922589	         7.441 ns/op	       0 B/op	       0 allocs/op
BenchmarkO1/input_20-4        	183734162	         6.986 ns/op	       0 B/op	       0 allocs/op
BenchmarkO1/input_50-4        	181818711	         6.839 ns/op	       0 B/op	       0 allocs/op
BenchmarkO1/input_100-4       	182213510	         6.725 ns/op	       0 B/op	       0 allocs/op
BenchmarkO1/input_900-4       	174299604	         6.817 ns/op	       0 B/op	       0 allocs/op
BenchmarkO1/input_1050-4      	175941924	         6.997 ns/op	       0 B/op	       0 allocs/op
BenchmarkO1/input_5000-4      	166095769	         7.698 ns/op	       0 B/op	       0 allocs/op
BenchmarkO1/input_9999-4      	139513252	         8.662 ns/op	       0 B/op	       0 allocs/op
```

**NOTE:** results will vary with the platform this test is executed.

The key takeaway is the results shows input of increasing size but time to execute relatively similar at about 7 ns.

### O(N)

O(N) describes an algorithm whose performance will grow linearly and in direct proportion to the size of the input data set. 

A benckmark of a slice data set of 1,000,000 integer with operation to get data is found [here](../example/benchmark/bigo/bigo_test.go).

Running the example, produces a result demonstrating a linearly increasing time require to complete.

```
BenchmarkON/input_1-4         	771780032	         1.529 ns/op	   0 B/op	       0 allocs/op
BenchmarkON/input_20-4        	159324080	         7.776 ns/op	   0 B/op	       0 allocs/op
BenchmarkON/input_50-4        	74474521	        15.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkON/input_100-4       	28960838	        43.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkON/input_900-4       	 4424012	       273.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkON/input_1050-4      	 3823518	       353.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkON/input_5000-4      	  761822	      1454 ns/op	       0 B/op	       0 allocs/op
BenchmarkON/input_9999-4      	  407455	      2917 ns/op	       0 B/op	       0 allocs/op
```

**NOTE:** results will vary with the platform this test is executed.

## O(N^2)

O(N^2) represents an algorithm whose performance is directly proportional to the square of the size of the input data set. 

A benchmark of an algorithm to determine if two int slices are duplicates. The example is [here](../example/benchmark/bigo/bigo_test.go).

Running the test reveal this result:

```
BenchmarkN2/input_1-4         	180612632	         6.385 ns/op	   0 B/op	       0 allocs/op
BenchmarkN2/input_2-4         	18520770	        65.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkN2/input_3-4         	 3321892	       363.5 ns/op	       0 B/op	       0 allocs/op
```

## O(2^N)

O(2^N) denotes an algorithm whose growth doubles with each addition to the input data set. The growth curve of an O(2^N) function is exponential — starting off very shallow, then rising meteorically. An example of an O(2^N) function is a recursive operation is demo [here](../example/benchmark/bigo/bigo_test.go)

```
Benchmark2PowN/input_1-4         	333059952	         3.564 ns/op	   0 B/op	       0 allocs/op
Benchmark2PowN/input_50-4        	13613836	        88.38 ns/op	       0 B/op	       0 allocs/op
Benchmark2PowN/input_100-4       	 7393389	       165.0 ns/op	       0 B/op	       0 allocs/op
Benchmark2PowN/input_1000-4      	  796298	      1632 ns/op	       0 B/op	       0 allocs/op
```