package main

import (
	"flag"
	"fmt"
	"log"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func withWorkers(numJobs int, numWorkers int) {
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	for i := 0; i < numWorkers; i++ {
		go func(id int, jobs <-chan int, results chan<- string) {
			for j := range jobs {
				results <- fmt.Sprintf("Worker: %v Jobs: [Cap: %v Len: %v] Input: %v Result: %v", id, cap(jobs), len(jobs), j, fib(j))
			}
		}(i, jobs, results)
	}

	for i := 0; i < numJobs; i++ {
		fmt.Println("Input: ", i)
		jobs <- i
	}
	close(jobs)

	for i := 0; i < numJobs; i++ {
		r := <-results
		fmt.Println("From", r)
	}
}

func main() {

	jobs := flag.Int("seq", 0, "Sequence of fibonace series to generate")
	workers := flag.Int("workers", 0, "Number of workers")

	flag.Parse()

	if *jobs == 0 {
		log.Fatal("The sequence should more than 1")
	}

	if *workers == 0 {
		log.Fatal("There should be 1 or more workers")
	}

	withWorkers(*jobs, *workers)
}
