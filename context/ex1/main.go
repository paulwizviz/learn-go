package main

import (
	"context"
	"log"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		log.Println("Waiting for cancellation ...")
		<-ctx.Done() // Blocks this Go routine until cancel() is called
		log.Fatalf("Cancelled")
	}()

	for {
		time.Sleep(10 * time.Second)
		cancel()
	}
}
