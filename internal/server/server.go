package server

import "net/http"

type Config struct {
	Addr string
}
type Dependencies struct {
}

func NewServer(cfg Config, deps Dependencies) *http.Server {

	mux :=  http.NewServeMux()
	mux.HandleFunc("GET /", func (w http.ResponseWriter, r *http.Request)  {
		_,_ = w.Write([]byte("Hello World"))
	})

	return &http.Server{
		Addr: cfg.Addr,
		Handler: mux,
	}

}
