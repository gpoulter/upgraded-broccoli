package diagnostics

import (
	"net/http"
	"github.com/sirupsen/logrus"
) 

func HealthzHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Caught a healthz request.")
		w.WriteHeader(http.StatusOK)
	}
}