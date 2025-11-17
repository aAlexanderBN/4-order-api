package product

import "go/api/internal/cart"

func get(id int) *Product {

	return &Product{
		Id:         1,
		Name:       "test name",
		Desciption: "test name",
	}
}

func delete(id int) bool {

	return true
}

func create(*cart.Carts) bool {
	return true
}

func update(*cart.Carts) bool {

	return true
}
