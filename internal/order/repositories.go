package order

import "gorm.io/gorm"

type OrderRepositories struct {
	Database *gorm.DB
}

func NewOrderRepository(database *gorm.DB) *OrderRepositories {

	return &OrderRepositories{
		Database: database,
	}
}

func (repo *OrderRepositories) Create(Ord *Order) (*Order, error) {
	result := repo.Database.Create(Ord)

	if result.Error != nil {
		return nil, result.Error
	}
	return Ord, nil
}

func (repo *OrderRepositories) GetById(id uint) (*Order, error) {

	var Ord Order

	result := repo.Database.Preload("Products.Product").First(&Ord, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &Ord, nil
}

func (repo *OrderRepositories) GetAll(id uint) ([]Order, error) {

	var arrord []Order

	repo.Database.Preload("Products.Product").Table("orders").
		Where("user_id = ?", id).
		Scan(&arrord)

	return arrord, nil
}
