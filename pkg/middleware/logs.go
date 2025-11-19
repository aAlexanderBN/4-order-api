package middleware

import (
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logrus.SetFormatter(&logrus.JSONFormatter{})
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)

	})
}
