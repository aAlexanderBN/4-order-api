package main

import (
	"fmt"
	"go/api/configs"
	"go/api/internal/myuser"
	myUser "go/api/internal/myuser"
	"go/api/internal/order"
	"go/api/internal/product"
	"go/api/pkg/db"
	"go/api/pkg/middleware"
	"net/http"
)

func App() http.Handler {

	conf := configs.LoadConfig("aalexanderbn@yandex.ru")

	db1 := db.NewDb(conf)

	router := http.NewServeMux()

	userRepositories := myuser.NewUserRepository(db1.DB)
	myUser.NewUserHandler(router, myUser.UserHandlerDeps{
		UserRepository: userRepositories,
		Config:         conf,
	})

	productRepositories := product.NewProductRepository(db1.DB)

	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepositories,
		UserRepository:    userRepositories,
		Config:            conf,
	})

	orderRepositories := order.NewOrderRepository(db1.DB)
	order.NewOrderHandler(router, order.OrderHandlerDeps{
		OrderRepository: orderRepositories,
		Config:          conf,
		UserRepository:  userRepositories,
	})

	stack := middleware.Logging(router)
	return stack
}

func main() {

	// conf := configs.LoadConfig("aalexanderbn@yandex.ru")

	// db1 := db.NewDb(conf)

	// router := http.NewServeMux()

	// userRepositories := myuser.NewUserRepository(db1.DB)
	// myUser.NewUserHandler(router, myUser.UserHandlerDeps{
	// 	UserRepository: userRepositories,
	// 	Config:         conf,
	// })

	// productRepositories := product.NewProductRepository(db1.DB)

	// product.NewProductHandler(router, product.ProductHandlerDeps{
	// 	ProductRepository: productRepositories,
	// 	UserRepository:    userRepositories,
	// 	Config:            conf,
	// })

	// orderRepositories := order.NewOrderRepository(db1.DB)
	// order.NewOrderHandler(router, order.OrderHandlerDeps{
	// 	OrderRepository: orderRepositories,
	// 	Config:          conf,
	// 	UserRepository:  userRepositories,
	// })

	app := app()

	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	fmt.Println("Server is listening on port 8081")

	er := server.ListenAndServe()
	if er != nil {
		fmt.Println("ERR Server isnot working")
	}

}
