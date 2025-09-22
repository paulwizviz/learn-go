package ex4

import (
	"runtime"
	"sync"
	"testing"
)

func myFunc() {
	for i := range 10 {
		i += i
	}
}

func BenchmarkGOPROMAX(b *testing.B) {
	var wg sync.WaitGroup
	numCPUs := runtime.NumCPU()
	b.Run("match max", func(b *testing.B) {
		for range b.N {
			runtime.GOMAXPROCS(numCPUs)
			for range numCPUs {
				wg.Add(1)
				go func() {
					defer wg.Done()
					myFunc()
				}()
			}
			wg.Wait()
		}
	})
	b.Run("more max", func(b *testing.B) {
		for range b.N {
			runtime.GOMAXPROCS(numCPUs)
			for range numCPUs + 10 {
				wg.Add(1)
				go func() {
					defer wg.Done()
					myFunc()
				}()
			}
			wg.Wait()
		}
	})
}
