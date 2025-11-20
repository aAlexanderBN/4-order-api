package product

import "gorm.io/gorm"

type ProductRepositories struct {
	Database *gorm.DB
}

func NewProductRepository(database *gorm.DB) *ProductRepositories {

	return &ProductRepositories{
		Database: database,
	}
}

func (repo *ProductRepositories) Create(pr *Product) (*Product, error) {
	result := repo.Database.Create(pr)

	if result.Error != nil {
		return nil, result.Error
	}
	return pr, nil
}

func (repo *ProductRepositories) Update(pr *Product) (*Product, error) {
	result := repo.Database.Save(pr)

	if result.Error != nil {
		return nil, result.Error
	}
	return pr, nil
}

func (repo *ProductRepositories) Delete(pr *Product) (*Product, error) {
	result := repo.Database.Delete(pr)

	if result.Error != nil {
		return nil, result.Error
	}
	return pr, nil
}

func (repo *ProductRepositories) GetById(id uint) (*Product, error) {

	var pr Product

	result := repo.Database.First(&pr, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &pr, nil
}
