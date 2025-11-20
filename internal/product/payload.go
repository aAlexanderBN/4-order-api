package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"Description"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
}
