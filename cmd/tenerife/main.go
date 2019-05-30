package main

import (
	"net"
	"net/http"
	"os"

	"upgraded-broccoli/internal/application"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	r.HandleFunc("/", application.HomeHandler)
	server := http.Server{
		Addr: net.JoinHostPort("", port),
		Handler: r,
	}
	server.ListenAndServe()
}