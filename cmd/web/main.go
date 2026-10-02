package main

import "github.com/VedantPatil1/gb-go/internal/server"


func main() {
	server := server.NewServer(server.Config{Addr: ":8001"}, server.Dependencies{})
	server.ListenAndServe()
}
