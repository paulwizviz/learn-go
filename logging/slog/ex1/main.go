package main

import (
	"log/slog"
	"os"
)

type Item struct {
	Name string
	Size string
}

func main() {

	item := struct {
		Name string
		Size string
	}{
		Name: "Hello",
		Size: "1000",
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Item",
		"name", item.Name,
		"size", item.Size,
	)

}
