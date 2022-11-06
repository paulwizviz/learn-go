package main

import (
	"context"
	"flag"
	"fmt"
	"time"
)

func main() {

	var tm int
	flag.IntVar(&tm, "time", 1, "time to sleep")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	go func(cx context.Context) {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("After 2 seconds")
		case <-cx.Done():
			fmt.Println("Cancelled")
		}
	}(ctx)

	time.Sleep(time.Duration(tm) * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
