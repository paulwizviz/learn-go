package setup

import (
	"log"
	"os"
	"testing"
)

type dbService struct {
	state string
}

func (d *dbService) Setup() {
	log.Println("Setting up")
	d.state = "initial setup"
}

func (d *dbService) MutateState(state string) {
	d.state = state
}

func (d *dbService) Reset() {
	log.Println("Resetting")
	d.state = "initial setup"
}

func (d dbService) State() string {
	return d.state
}

var databaseService = dbService{}

func TestMain(m *testing.M) {
	log.Println("Setup test")
	exitVal := m.Run()
	log.Println("End test")
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
	defer func() {
		log.Printf("Database state before reset: %v", databaseService.State())
		databaseService.Reset()
		log.Printf("Database state after reset: %v", databaseService.State())
	}()
	log.Println("Executing first test")
	databaseService.MutateState("Mutating first test")
}

func TestSecond(t *testing.T) {
	defer func() {
		log.Printf("Database state before reset: %v", databaseService.State())
		databaseService.Reset()
		log.Printf("Database state after reset: %v", databaseService.State())
	}()
	log.Println("Executing second test")
	databaseService.MutateState("Mutating second test")
}
