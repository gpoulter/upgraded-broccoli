package main

import (
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	os.Getenv("PORT")
	server = http.Server{
		Addr: net.JoinHostPort("", port),
		Handler: nil,
	}
	http.ListenAndServe()
}