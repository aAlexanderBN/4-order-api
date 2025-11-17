package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id         int
	Name       string
	Desciption string
}
