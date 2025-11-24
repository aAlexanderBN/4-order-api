package order

import (
	"go/api/internal/product"

	"gorm.io/gorm"
)

type OrderProduct struct {
	OrderID   uint            `gorm:"primaryKey" json:"order_id"`
	ProductID uint            `gorm:"primaryKey" json:"product_id"`
	Quantity  int             `json:"quantity" validate:"required,min=1"`
	Product   product.Product `gorm:"foreignKey:ProductID"`
	//Order     Order           `gorm:"foreignKey:OrderID"`
}

type Order struct {
	gorm.Model
	UserID   uint           `json:"user_id" validate:"required"`
	Products []OrderProduct `gorm:"foreignKey:OrderID" json:"products""`
}
