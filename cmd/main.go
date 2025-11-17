package main

import (
	"fmt"
	"go/api/configs"
	myUser "go/api/internal/myuser"
	"net/http"
)

func main() {

	conf := configs.LoadConfig("aalexanderbn@yandex.ru")

	router := http.NewServeMux()
	myUser.NewAuthHandler(router, myUser.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")

	server.ListenAndServe()

}
