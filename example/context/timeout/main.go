package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)

	log.Println("Process is blocked for 10 secs")
	<-ctx.Done()
	log.Println("Exiting...")
}
