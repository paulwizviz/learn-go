package ex2

import (
	"log"
	"testing"
)

func TestThird(t *testing.T) {
	defer func() {
		log.Printf("Database state before reset: %v", databaseService.State())
		databaseService.Reset()
		log.Printf("Database state after reset: %v", databaseService.State())
	}()
	log.Println("Third")
	databaseService.MutateState("Mutating second test")
}
