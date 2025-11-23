package middleware

import (
	"context"
	"fmt"
	"go/api/configs"
	"go/api/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextPhoneKey  key = "ContextPhoneKey"
	ContextUserIDKey key = "ContextUserIDKey"
)

func writeUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authedHeader, "Bearer ") {
			writeUnauthed(w)
			return
		}
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.MyUser.Secret).Pasre(token)
		if !isValid {
			writeUnauthed(w)
			return
		}

		user, err := userRepo.GetByNameUser(data.Phone)
		if err != nil {
			writeUnauthed(w)
			return
		}

		fmt.Println(token)
		fmt.Println(isValid)
		fmt.Println(data)

		ctx := context.WithValue(r.Context(), ContextPhoneKey, data.Phone)
		ctx = context.WithValue(r.Context(), ContextUserIDKey, user.ID)

		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})

}
