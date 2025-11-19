package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// Создаем отдельный логгер для middleware
var logger = logrus.New()

func init() {
	// Настраиваем логгер один раз
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logrus.SetFormatter(&logrus.JSONFormatter{})

		l := map[string]any{
			"Method": r.Method,
			"URL":    r.URL.Path,
		}

		logrus.WithFields(logrus.Fields(l)).Info("Request received")
		next.ServeHTTP(w, r)

	})
}
