package ex2

import (
	"log"
	"testing"
)

func TestFourth(t *testing.T) {
	defer func() {
		log.Printf("Database state before reset: %v", databaseService.State())
		databaseService.Reset()
		log.Printf("Database state after reset: %v", databaseService.State())
	}()
	log.Println("Fourth test")
	databaseService.MutateState("Mutating second test")
}
