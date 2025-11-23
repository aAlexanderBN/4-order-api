package order

import (
	"go/api/internal/product"

	"gorm.io/gorm"
)

type OrderProduct struct {
	OrderID   uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
	Product   product.Product
	Quantity  int `json:"quantity" validate:"required,min=1"`
}

type Order struct {
	gorm.Model
	UserID   uint           `json:"user_id" validate:"required"`
	Products []OrderProduct `gorm:"many2many:order_products; json:"products""`
}
