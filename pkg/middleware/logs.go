package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logrus.SetFormatter(&logrus.JSONFormatter{})

		l := map[string]any{
			"Method": r.Method,
			"URL":    r.URL.Path,
		}

		logrus.WithFields(logrus.Fields(l))
		next.ServeHTTP(w, r)

	})
}
