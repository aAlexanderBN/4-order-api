package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id         int            `json:"id"`
	Name       string         `json:"name"`
	Desciption string         `json:"desciption"`
	Images     pq.StringArray `json:"images"`
}
