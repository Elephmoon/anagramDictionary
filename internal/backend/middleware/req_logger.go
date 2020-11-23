package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func LogHTTPRequest(logger logrus.FieldLogger) func(handler http.Handler) http.Handler {
	log := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.WithFields(logrus.Fields{
				"uri":        r.URL.Path,
				"method":     r.Method,
				"user-agent": r.UserAgent(),
			}).Info()
			next.ServeHTTP(w, r)
		})
	}
	return log
}
