package main

import (
	"log/slog"

	"github.com/radio-noise-project/last-order/internal/api/server"
)

func main() {
	slog.Info("Start Last-Order")
	server.Start()
}
