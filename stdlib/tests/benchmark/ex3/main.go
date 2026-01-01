package main

import (
	"log"
	"testing"

	"github.com/paulwizviz/learn-go/tests/benchmark/ex3/fib"
)

func BenchmarkFibRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib.FibRecur(30)
	}
}

func main() {
	bc := testing.Benchmark(BenchmarkFibRec)
	log.Println("Memory allocations ", bc.MemAllocs)
	log.Println("Memory bytes ", bc.MemBytes)
	log.Println("Number of runs ", bc.N)
	log.Println("Time taken ", bc.T)
}
