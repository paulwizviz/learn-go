package ex1

import "testing"

// benchFunc is a function to be benchmarked
func benchFunc() {
	for i := 0; i < 100; i++ {

	}
}

func BenchmarkOneOff(b *testing.B) {
	for count := 0; count < b.N; count++ {
		benchFunc()
	}
}
