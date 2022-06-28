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
		log.Println("Wait for cancellation ...")
		<-ctx.Done()
		log.Fatal("Cancelled")
	}()

	for {
		time.Sleep(10 * time.Second)
		cancel()
	}
}
