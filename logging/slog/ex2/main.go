package main

import (
	"log/slog"
	"os"
)

func DoSomething() {
	slog.Info("Do something",
		slog.Int64("value", 6000),
	)
}

func main() {
	jsonHdl := slog.NewJSONHandler(os.Stdout, nil)
	slog.SetDefault(slog.New(jsonHdl))
	DoSomething()
}
