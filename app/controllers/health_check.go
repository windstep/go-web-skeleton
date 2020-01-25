package controllers

import (
	"github.com/windstep/go-web-skeleton/logger"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Status: OK"))
	if err != nil {
		logger.Logger.Error(err)
	}
}
