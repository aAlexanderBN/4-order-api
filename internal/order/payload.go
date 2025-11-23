package order

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID   int
	Products []int `json:"products"`
}
