package app

import (
	"fmt"
	"net/http"

	"github.com/windstep/go-web-skeleton/app/router"
	"github.com/windstep/go-web-skeleton/config"
	"github.com/windstep/go-web-skeleton/logger"
)

var cfg *config.Config

func init() {
	cfg = config.GetInstance()
}

func Run() {
	logger.Logger.Infof("Server is running on port %d", cfg.Port)
	listen(cfg.Port)
}

func listen(port int) {
	mux := router.New()
	logger.Logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
