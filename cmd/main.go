package main

import (
	"fmt"
	"go/api/configs"
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

	myUser.NewAuthHandler(router, myUser.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logging(router),
	}

	fmt.Println("Server is listening on port 8081")

	server.ListenAndServe()

}
