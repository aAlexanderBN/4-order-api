package cart

func create(newname string) *Carts {

	return &Carts{
		Id:   1,
		Name: newname,
	}
}

func getById(id int) *Carts {

	return &Carts{
		Id:   1,
		Name: "найденное имя",
	}
}
