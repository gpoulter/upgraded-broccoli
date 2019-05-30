package main

import (
	"log"
	"net"
	"net/http"
	"os"

	_ "net/http/pprof"

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
	log.Println(server.ListenAndServe())
}