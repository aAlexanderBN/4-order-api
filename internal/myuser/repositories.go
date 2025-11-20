package myuser

import "gorm.io/gorm"

type UserRepositories struct {
	Database *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserRepositories {

	return &UserRepositories{
		Database: database,
	}
}

func (repo *UserRepositories) CreateUser(user *Users) (*Users, error) {
	result := repo.Database.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepositories) UpdateUser(user *Users) (*Users, error) {
	result := repo.Database.Save(user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepositories) GetByNameUser(name string) (*Users, error) {

	var user Users

	result := repo.Database.First(&user, name)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
