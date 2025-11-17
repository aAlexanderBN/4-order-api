package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Desciption string `json:"desciption"`
}
