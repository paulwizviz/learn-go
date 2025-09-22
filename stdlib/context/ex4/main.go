package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	log.Println("Process is blocked for 10 secs")
	<-ctx.Done()
	log.Println("Exiting...")
}
