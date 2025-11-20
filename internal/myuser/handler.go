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

		var name string

		userdb, err := handler.UserRepository.GetByNameUser(body.Name)
		if err != nil {
			name = body.Name
		} else if userdb.Name == "" {
			name = userdb.Name
		} else {
			name = body.Name
		}

		user := Users{
			Name: name,
			Code: code,
		}

		id := uuid.New()

		user.SessionID = id.String()

		var createdUser *Users
		if userdb.Name == "" {
			createdUser, err = handler.UserRepository.CreateUser(&user)
		} else {
			createdUser, err = handler.UserRepository.UpdateUser(&user)
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

		userdb, err := handler.UserRepository.GetByNameUser(body.Name)
		if err != nil {
			fmt.Println("Ошибка чтения из базы")
			return
		}

		if userdb.SessionID == body.SessionID && userdb.Code == body.Code {
			userdb.Token, err = jwt.NewJWT(handler.Config.MyUser.Secret).Create(userdb.Phone)
			res.Json(w, userdb, 201)
		} else {
			res.Json(w, nil, 201)
		}

	}
}

func verifiCode(n int) int {

	return rand.Intn(n * 99)
}
