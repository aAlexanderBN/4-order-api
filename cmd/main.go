package main

import (
	"fmt"
	"go/api/configs"
	"go/api/internal/myuser"
	myUser "go/api/internal/myuser"
	"go/api/internal/product"
	"go/api/pkg/db"
	"go/api/pkg/middleware"
	"net/http"
)

func main() {

	conf := configs.LoadConfig("aalexanderbn@yandex.ru")

	db1 := db.NewDb(conf)

	router := http.NewServeMux()

	productRepositories := product.NewProductRepository(db1.DB)

	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepositories,
	})

	userRepositories := myuser.NewUserRepository(db1.DB)
	myUser.NewUserHandler(router, myUser.UserHandlerDeps{
		UserRepository: userRepositories,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logging(router),
	}

	fmt.Println("Server is listening on port 8081")

	er := server.ListenAndServe()
	if er != nil {
		fmt.Println("ERR Server isnot working")
	}

}
