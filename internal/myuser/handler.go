package myuser

import (
	"fmt"
	"go/api/configs"
	"go/api/pkg/jwt"
	"go/api/pkg/req"
	"go/api/pkg/res"
	"math/rand"
	"net/http"

	"github.com/google/uuid"
)

type UserHandlerDeps struct {
	UserRepository *UserRepositories
	*configs.Config
}

type UserHandler struct {
	UserRepository *UserRepositories
	*configs.Config
}

const lenVC = 4

func NewUserHandler(router *http.ServeMux, deps UserHandlerDeps) {

	handler := &UserHandler{
		UserRepository: deps.UserRepository,
		Config:         deps.Config,
	}

	router.HandleFunc("GET /autByPhone", handler.AuthByPhone())
	router.HandleFunc("GET /verify", handler.Verify())
}

func (handler *UserHandler) AuthByPhone() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[Users](&w, r)
		if err != nil {
			return
		}
		code := verifiCode(lenVC)

		userdb, err := handler.UserRepository.GetByNameUser(body.Phone)

		id := uuid.New()

		var createdUser *Users
		if userdb.Phone == "" {
			// Пользователь НЕ найден - создаем нового
			user := Users{
				Phone:     body.Phone,
				Code:      code,
				SessionID: id.String(),
			}
			createdUser, err = handler.UserRepository.CreateUser(&user)
		} else {
			// Пользователь найден - обновляем СУЩЕСТВУЮЩИЙ объект userdb
			// userdb уже содержит ID из БД!
			userdb.Code = code
			userdb.SessionID = id.String()
			createdUser, err = handler.UserRepository.UpdateUser(userdb)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		createdUser.Code = 0

		res.Json(w, createdUser, 201)

		fmt.Println("Send by phone code=", code)
	}
}

func (handler *UserHandler) Verify() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[Users](&w, r)
		if err != nil {
			return
		}

		userdb, err := handler.UserRepository.GetByNameUser(body.Phone)
		if err != nil {
			fmt.Println("Ошибка чтения из базы")
			return
		}

		if userdb.SessionID == body.SessionID && userdb.Code == body.Code {
			token, err := jwt.NewJWT(handler.Config.MyUser.Secret).Create(userdb.Phone)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			response := map[string]string{"token": token}

			res.Json(w, response, 201)
		} else {
			http.Error(w, "Invalid code or sessionId", http.StatusUnauthorized)
		}

	}
}

func verifiCode(n int) int {

	return rand.Intn(n * 100)
}
