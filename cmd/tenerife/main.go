package main

import (
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"upgraded-broccoli/internal/application"
	"upgraded-broccoli/internal/diagnostics"
)

func main() {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)

	logger.Info("Starting the application")

	port := os.Getenv("PORT")
	if port == "" {
		logger.Fatal("Port is not provided")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", application.HomeHandler(logger))
	r.HandleFunc("/healthz", diagnostics.HealthzHandler(logger))

	server := http.Server{
		Addr: net.JoinHostPort("", port),
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Error("Server stopped with an error: %v", err)
	}
}
