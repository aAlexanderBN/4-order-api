package order

import (
	"go/api/internal/product"

	"gorm.io/gorm"
)

// type tableProduct struct {
// 	Product product.Product `validate:"required"`
// 	cout    int             `validate:"required"`
// }

type Order struct {
	gorm.Model
	UserID   uint              `json:"user_id" validate:"required"`
	Products []product.Product `gorm:"many2many:order_products; json:"products""`
}
