package main

import (
	"context"
	"log"
	"time"
)

// NOTE: This example is serve only to demonstrate concept
// You should not attempt this is mission critical
// deployment. There are potential for memory leak
func main() {
	ctx := context.Background()
	// The function cancel returned from WithTimeout
	// could cause memory leak without be used.
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)

	log.Println("Process is blocked for 10 secs")
	<-ctx.Done()
	log.Println("Exiting...")
}
