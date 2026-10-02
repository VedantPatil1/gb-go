package main

import (
	"log/slog"
	"os"

	"github.com/VedantPatil1/gb-go/internal/server"
)


func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	server := server.NewServer(server.Config{Addr: ":8001"}, server.Dependencies{Logger: logger})

	logger.Info("starting server", "addr", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		logger.Error("server stopped", "error", err)
	}
}
