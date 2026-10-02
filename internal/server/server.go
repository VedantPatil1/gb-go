// Package server provides the HTTP server for the application.
package server

import (
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/VedantPatil1/gb-go/internal/assets"
	"github.com/VedantPatil1/gb-go/internal/server/components"
)

type Config struct {
	Addr string
}

type Dependencies struct {
	Logger *slog.Logger
}

func NewServer(cfg Config, deps Dependencies) *http.Server {

	mux := http.NewServeMux()

	staticFS, err := fs.Sub(assets.Files, "static")
	if err != nil {
		panic(err)
	}

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		deps.Logger.Info("handling request", "method", r.Method, "path", r.URL.Path)

		if err := components.HelloWorld().Render(r.Context(), w); err != nil {
			deps.Logger.Error("failed to render page", "error", err)

			http.Error(w, "internal server error", http.StatusInternalServerError)
		}

	})

	return &http.Server{
		Addr:    cfg.Addr,
		Handler: mux,
	}

}
