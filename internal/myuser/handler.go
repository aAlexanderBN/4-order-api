package myUser

import (
	"fmt"
	"go/api/configs"
	"math/rand"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

const lenVC = 4

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config}
	router.HandleFunc("GET /autByPhone", handler.AuthByPhone())
	router.HandleFunc("GET /verify/{code}", handler.Verify())
}

func (handler *AuthHandler) AuthByPhone() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		code := verifiCode(lenVC)
		//send code by phone
		//record to DB
		fmt.Println("Dend  by phone code=", code)
	}
}

func (handler *AuthHandler) Verify() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		code := r.PathValue("code")
		_ = code
		// чтение из БД данные о коде в паке pkg
		user_code := "1234"
		if user_code == code {
			fmt.Println("Verify true")
		} else {
			fmt.Println("Verify false")
		}

	}
}

func verifiCode(n int) int {

	return rand.Intn(n * 99)
}
