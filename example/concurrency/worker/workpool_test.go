package worker_test

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// In this case, we'll only receive on jobs channel
// only ever to send on results channel
// if arrows are not specified the channel is
// considered to be bidireactional
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func Example_workerPool() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Multiple workers working concurrently
	// Each worker pull off the queue
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 10; j++ {
		fmt.Printf("Result: %v\n", <-results)
	}

}
