package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func cleanup() {
	fmt.Println("cleanup")
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		switch <-c {
		// kill -SIGHUP XXXX [XXXX - PID for your program]
		case syscall.SIGHUP:
			fmt.Println("Signal hang up triggered.")
			cleanup()
			os.Exit(0)
		// kill -SIGINT XXXX or Ctrl+c  [XXXX - PID for your program]
		case syscall.SIGINT:
			fmt.Println("Signal interrupt triggered.")
			cleanup()
			os.Exit(0)
		// kill -SIGTERM XXXX [XXXX - PID for your program]
		case syscall.SIGTERM:
			fmt.Println("Signal terminte triggered.")
			cleanup()
			os.Exit(0)
		// kill -SIGQUIT XXXX [XXXX - PID for your program]
		case syscall.SIGQUIT:
			fmt.Println("Signal quit triggered.")
			cleanup()
			os.Exit(0)
		default:
			fmt.Println("Unknown signal.")
			cleanup()
			os.Exit(1)
		}
	}()

	for {
		fmt.Println("sleeping...")
		time.Sleep(10 * time.Second) // or runtime.Gosched() or similar per @misterbee
	}
}
