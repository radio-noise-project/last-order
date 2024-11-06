package main

import (
	"log/slog"

	"github.com/radio-noise-project/last-order/internal/api"
)

func main() {
	slog.Info("Start Last-Order")
	api.Server()
}
