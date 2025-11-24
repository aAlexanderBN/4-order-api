package product

import (
	"go/api/configs"
	"go/api/internal/myuser"
	"go/api/pkg/middleware"
	"go/api/pkg/req"
	"go/api/pkg/res"
	"net/http"
	"strconv"

	"github.com/lib/pq"
)

type ProductHandlerDeps struct {
	ProductRepository *ProductRepositories
	Config            *configs.Config
	UserRepository    *myuser.UserRepositories
}

type ProductHandler struct {
	ProductRepository *ProductRepositories
}

func NewProductHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := &ProductHandler{
		ProductRepository: deps.ProductRepository}

	router.Handle("POST /product", middleware.IsAuthed(handler.Create(), deps.Config, deps.UserRepository))
	router.HandleFunc("PATCH /product", handler.Update())
	router.HandleFunc("DELETE /product/{id}", handler.Delete())
	router.HandleFunc("GET /product/{id}", handler.GetById())
}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Прочитать боди
		body, err := req.HandleBody[Product](&w, r)
		if err != nil {
			return
		}

		product := Product{
			Name:        body.Name,
			Description: body.Description,
			Images:      pq.StringArray{},
		}

		createdLProduct, err := handler.ProductRepository.Create(&product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdLProduct, 201)
	}
}

func (handler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Прочитать боди
		body, err := req.HandleBody[Product](&w, r)
		if err != nil {
			return
		}

		// product := Product{
		// 	Name:        body.Name,
		// 	Description: body.Description,
		// 	//Images:      pq.StringArray{},
		// }

		createdLProduct, err := handler.ProductRepository.Update(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdLProduct, 201)
	}
}

func (handler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Прочитать боди
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var pr *Product
		pr, err = handler.ProductRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		_, err = handler.ProductRepository.Delete(pr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, "Удалено", 200)
	}
}

func (handler *ProductHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Прочитать боди
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var createdLProduct *Product

		createdLProduct, err = handler.ProductRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		res.Json(w, createdLProduct, 202)

	}
}
