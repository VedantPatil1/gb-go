// Package server provides the HTTP server for the application.
package server

import (
	"log/slog"
	"net/http"
)

type Config struct {
	Addr string
}

type Dependencies struct {
	Logger *slog.Logger
}

func NewServer(cfg Config, deps Dependencies) *http.Server {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		deps.Logger.Info("handling request", "method", r.Method, "path", r.URL.Path)

		_, _ = w.Write([]byte("Hello World"))
	})

	return &http.Server{
		Addr:    cfg.Addr,
		Handler: mux,
	}

}
