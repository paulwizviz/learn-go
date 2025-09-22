package ex3

import (
	"runtime"
	"sync"
	"testing"
)

func unfairWorker(yield bool) {
	for i := range 100000 {
		// Simulate some work
		_ = i * i
		if yield && i%1000 == 0 {
			runtime.Gosched() // Yield periodically
		}
	}
}

func BenchmarkYield(b *testing.B) {
	b.Run("no yield", func(b *testing.B) {
		var wg sync.WaitGroup
		for range b.N {
			wg.Add(2)
			go func() {
				defer wg.Done()
				unfairWorker(false)
			}()
			go func() {
				defer wg.Done()
				unfairWorker(false)
			}()
			wg.Wait()
		}
	})
	b.Run("yield processor", func(b *testing.B) {
		var wg sync.WaitGroup
		for range b.N {
			wg.Add(2)
			go func() {
				defer wg.Done()
				unfairWorker(true)
			}()
			go func() {
				defer wg.Done()
				unfairWorker(true)
			}()
			wg.Wait()
		}
	})
}
