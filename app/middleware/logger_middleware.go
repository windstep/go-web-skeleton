package middleware

import (
	"github.com/windstep/go-web-skeleton/logger"
	"net/http"
)

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Logger.Infof("[%s] %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}
