package cancellation_test

import (
	"context"
	"fmt"
)

func Example_cancel() {

	ch := make(chan int)
	cx, cancel := context.WithCancel(context.Background())

	n := 0
	go func(ctx context.Context) {
		// Loop forever
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
			}
		}
	}(cx)

	for {
		i := <-ch
		if i == 10 {
			cancel() // send cancel to context
			fmt.Printf("Index: %v Done\n", i)
			return
		} else {
			fmt.Printf("Index: %v\n", i)
		}
	}

	// Output:
	// Index: 0
	// Index: 1
	// Index: 2
	// Index: 3
	// Index: 4
	// Index: 5
	// Index: 6
	// Index: 7
	// Index: 8
	// Index: 9
	// Index: 10 Done

}
