package main

import (
	"github.com/sirupsen/logrus"
	"github.com/windstep/go-web-skeleton/app"
	"github.com/windstep/go-web-skeleton/config"
	"github.com/windstep/go-web-skeleton/logger"
)

var l *logrus.Logger
var conf *config.Config

func init() {
	conf = config.GetInstance()
	l = logger.Logger
	l.SetLevel(conf.LogLevel)
}

func main() {
	app.Run()
}
