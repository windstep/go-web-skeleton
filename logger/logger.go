package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

func init() {
	Logger = log.New()
	Logger.SetFormatter(&log.TextFormatter{})
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(log.InfoLevel)
}
