package order

import (
	"go/api/configs"
	"go/api/pkg/middleware"
	"go/api/pkg/req"
	"go/api/pkg/res"
	"net/http"
	"strconv"
)

type OrderHandlerDeps struct {
	OrderRepository *OrderRepositories
	Config          *configs.Config
}

type OrderHandler struct {
	OrderRepository *OrderRepositories
}

func NewOrderHandler(router *http.ServeMux, deps OrderHandlerDeps) {
	handler := &OrderHandler{
		OrderRepository: deps.OrderRepository}

	router.Handle("POST /order", middleware.IsAuthed(handler.Create(), deps.Config))
	router.HandleFunc("GET /order/{id}", handler.GetById())
	router.HandleFunc("GET /order/", handler.GetAll())
}

func (handler *OrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Прочитать боди
		body, err := req.HandleBody[Order](&w, r)
		if err != nil {
			return
		}

		product := Order{
			UserID:   body.UserID,
			Products: body.Products,
		}

		createdLProduct, err := handler.OrderRepository.Create(&product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdLProduct, 201)
	}
}

func (handler *OrderHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 0)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		createdLProduct, err := handler.OrderRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdLProduct, 201)
	}
}

func (handler *OrderHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var UserId uint
		UserId = 4

		createdLProduct, err := handler.OrderRepository.GetByAll(UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdLProduct, 201)
	}
}
